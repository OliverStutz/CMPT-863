package analysis

import (
	"math"
	"sort"
)

// Descriptive and inferential statistics for the dashboard (plan 06, section
// 2). No external dependency: the t distribution comes from the regularised
// incomplete beta function (continued fraction, Numerical Recipes 6.4).
// Functions return nil pointers where a value is undefined.

func fp(v float64) *float64 { return &v }

func mean(xs []float64) *float64 {
	if len(xs) == 0 {
		return nil
	}
	s := 0.0
	for _, x := range xs {
		s += x
	}
	return fp(s / float64(len(xs)))
}

// variance is the sample variance (n - 1); nil for n < 2.
func variance(xs []float64) *float64 {
	m := mean(xs)
	if m == nil || len(xs) < 2 {
		return nil
	}
	s := 0.0
	for _, x := range xs {
		s += (x - *m) * (x - *m)
	}
	return fp(s / float64(len(xs)-1))
}

func sd(xs []float64) *float64 {
	v := variance(xs)
	if v == nil {
		return nil
	}
	return fp(math.Sqrt(*v))
}

// percentile uses linear interpolation between order statistics (R type 7,
// the spreadsheet default); q in [0, 1].
func percentile(xs []float64, q float64) *float64 {
	if len(xs) == 0 {
		return nil
	}
	s := append([]float64(nil), xs...)
	sort.Float64s(s)
	h := q * float64(len(s)-1)
	lo := int(math.Floor(h))
	if lo >= len(s)-1 {
		return fp(s[len(s)-1])
	}
	return fp(s[lo] + (h-float64(lo))*(s[lo+1]-s[lo]))
}

func median(xs []float64) *float64 { return percentile(xs, 0.5) }

// lbeta is log B(a, b).
func lbeta(a, b float64) float64 {
	la, _ := math.Lgamma(a)
	lb, _ := math.Lgamma(b)
	lab, _ := math.Lgamma(a + b)
	return la + lb - lab
}

// betacf evaluates the continued fraction of the incomplete beta function
// by the modified Lentz method.
func betacf(a, b, x float64) float64 {
	const (
		maxIter = 100000
		eps     = 1e-15
		tiny    = 1e-300
	)
	qab, qap, qam := a+b, a+1, a-1
	c, d := 1.0, 1-qab*x/qap
	if math.Abs(d) < tiny {
		d = tiny
	}
	d = 1 / d
	h := d
	for m := 1; m <= maxIter; m++ {
		fm := float64(m)
		m2 := 2 * fm
		aa := fm * (b - fm) * x / ((qam + m2) * (a + m2))
		d = 1 + aa*d
		if math.Abs(d) < tiny {
			d = tiny
		}
		c = 1 + aa/c
		if math.Abs(c) < tiny {
			c = tiny
		}
		d = 1 / d
		h *= d * c
		aa = -(a + fm) * (qab + fm) * x / ((a + m2) * (qap + m2))
		d = 1 + aa*d
		if math.Abs(d) < tiny {
			d = tiny
		}
		c = 1 + aa/c
		if math.Abs(c) < tiny {
			c = tiny
		}
		d = 1 / d
		del := d * c
		h *= del
		if math.Abs(del-1) < eps {
			break
		}
	}
	return h
}

// betaInc is the regularised incomplete beta function I_x(a, b).
func betaInc(a, b, x float64) float64 {
	switch {
	case x <= 0:
		return 0
	case x >= 1:
		return 1
	}
	front := math.Exp(a*math.Log(x) + b*math.Log1p(-x) - lbeta(a, b))
	if x < (a+1)/(a+b+2) {
		return front * betacf(a, b, x) / a
	}
	return 1 - front*betacf(b, a, 1-x)/b
}

// tTwoSided is the two-sided p value of a t statistic with df degrees of
// freedom: P(|T| >= |t|) = I_{df/(df+t^2)}(df/2, 1/2).
func tTwoSided(t, df float64) float64 {
	if math.IsInf(t, 0) {
		return 0
	}
	return betaInc(df/2, 0.5, df/(df+t*t))
}

// tCDF is P(T <= t).
func tCDF(t, df float64) float64 {
	half := tTwoSided(t, df) / 2
	if t >= 0 {
		return 1 - half
	}
	return half
}

// tQuantile returns the t with P(T <= t) = p (0 < p < 1), by bisection on the
// CDF.
func tQuantile(p, df float64) float64 {
	if p == 0.5 {
		return 0
	}
	if p < 0.5 {
		return -tQuantile(1-p, df)
	}
	lo, hi := 0.0, 1.0
	for tCDF(hi, df) < p && hi < 1e12 {
		hi *= 2
	}
	for i := 0; i < 200; i++ {
		mid := (lo + hi) / 2
		if tCDF(mid, df) < p {
			lo = mid
		} else {
			hi = mid
		}
		if hi-lo < 1e-12 {
			break
		}
	}
	return (lo + hi) / 2
}

// normTwoSided is the two-sided p value of a standard normal z.
func normTwoSided(z float64) float64 { return math.Erfc(math.Abs(z) / math.Sqrt2) }

// ci95 is the t-based 95% interval m +- t(.975, df) * se; nil when df is not
// positive.
func ci95(m, se, df float64) *[2]float64 {
	if df <= 0 || math.IsNaN(se) {
		return nil
	}
	h := tQuantile(0.975, df) * se
	return &[2]float64{m - h, m + h}
}

// oneSampleResult holds a one-sample or paired t test against 0 (df = n - 1).
type oneSampleResult struct {
	n              int
	mean, t, df, p *float64
	ci             *[2]float64
}

// oneSampleT tests the mean of xs against 0.
func oneSampleT(xs []float64) oneSampleResult {
	r := oneSampleResult{n: len(xs), mean: mean(xs)}
	s := sd(xs)
	if s == nil {
		return r
	}
	n := float64(len(xs))
	se := *s / math.Sqrt(n)
	r.df = fp(n - 1)
	r.ci = ci95(*r.mean, se, n-1)
	if se > 0 {
		r.t = fp(*r.mean / se)
		r.p = fp(tTwoSided(*r.t, n-1))
	}
	return r
}

// pairedT is the paired t test of a - b (equal length, same participants).
func pairedT(a, b []float64) oneSampleResult {
	d := make([]float64, len(a))
	for i := range a {
		d[i] = a[i] - b[i]
	}
	return oneSampleT(d)
}

type welchResult struct {
	diff, t, df, p, g *float64
	ci                *[2]float64
}

// welchT compares the means of x and y (x minus y) without assuming equal
// variances, with Welch-Satterthwaite degrees of freedom, and gives Hedges g
// (Cohen's d with the pooled SD, times J = 1 - 3 / (4 (n1 + n2) - 9)).
func welchT(x, y []float64) welchResult {
	var r welchResult
	mx, my := mean(x), mean(y)
	if mx == nil || my == nil {
		return r
	}
	r.diff = fp(*mx - *my)
	vx, vy := variance(x), variance(y)
	if vx == nil || vy == nil {
		return r
	}
	n1, n2 := float64(len(x)), float64(len(y))
	a, b := *vx/n1, *vy/n2
	se := math.Sqrt(a + b)
	if se > 0 {
		df := (a + b) * (a + b) / (a*a/(n1-1) + b*b/(n2-1))
		r.df = fp(df)
		r.t = fp(*r.diff / se)
		r.p = fp(tTwoSided(*r.t, df))
		r.ci = ci95(*r.diff, se, df)
	} else {
		// Zero variance in both groups: t and p are undefined, the interval
		// collapses to the difference (as ci95 does for one sample).
		r.ci = &[2]float64{*r.diff, *r.diff}
	}
	sp := math.Sqrt(((n1-1)**vx + (n2-1)**vy) / (n1 + n2 - 2))
	if sp > 0 {
		j := 1 - 3/(4*(n1+n2)-9)
		r.g = fp(*r.diff / sp * j)
	}
	return r
}

type propResult struct {
	p1, p2, diff, z, p *float64
}

// twoProportionZ is the pooled two-proportion z test of x1/n1 against x2/n2.
func twoProportionZ(x1, n1, x2, n2 int) propResult {
	var r propResult
	if n1 > 0 {
		r.p1 = fp(float64(x1) / float64(n1))
	}
	if n2 > 0 {
		r.p2 = fp(float64(x2) / float64(n2))
	}
	if r.p1 == nil || r.p2 == nil {
		return r
	}
	r.diff = fp(*r.p1 - *r.p2)
	pp := float64(x1+x2) / float64(n1+n2)
	se := math.Sqrt(pp * (1 - pp) * (1/float64(n1) + 1/float64(n2)))
	if se > 0 {
		r.z = fp(*r.diff / se)
		r.p = fp(normTwoSided(*r.z))
	}
	return r
}

// ranks gives average ranks (1-based) with ties sharing the mean rank.
func ranks(xs []float64) []float64 {
	idx := make([]int, len(xs))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(a, b int) bool { return xs[idx[a]] < xs[idx[b]] })
	out := make([]float64, len(xs))
	for i := 0; i < len(idx); {
		j := i
		for j+1 < len(idx) && xs[idx[j+1]] == xs[idx[i]] {
			j++
		}
		r := float64(i+j)/2 + 1
		for k := i; k <= j; k++ {
			out[idx[k]] = r
		}
		i = j + 1
	}
	return out
}

// pearson is the correlation of x and y; nil when undefined.
func pearson(x, y []float64) *float64 {
	if len(x) < 2 || len(x) != len(y) {
		return nil
	}
	mx, my := *mean(x), *mean(y)
	var sxy, sxx, syy float64
	for i := range x {
		sxy += (x[i] - mx) * (y[i] - my)
		sxx += (x[i] - mx) * (x[i] - mx)
		syy += (y[i] - my) * (y[i] - my)
	}
	if sxx == 0 || syy == 0 {
		return nil
	}
	return fp(sxy / math.Sqrt(sxx*syy))
}

// spearman is the rank correlation (Pearson on average ranks, so ties are
// handled); nil for fewer than 3 pairs or a constant variable.
// DECISION: 3 pairs at least, since two points always correlate +-1.
func spearman(x, y []float64) *float64 {
	if len(x) < 3 {
		return nil
	}
	return pearson(ranks(x), ranks(y))
}
