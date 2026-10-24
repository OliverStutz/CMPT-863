package domain

import (
	sc "study/api/pkg/studyconfig"
)

// Allocate picks the group and denial slide for one consenting session and
// increments the counters (5.5). It is called inside the consent event's
// transaction, which also reads and writes counters/allocation, so concurrent
// consents serialise on that document.
//
//  1. Group: the group with the lower count; on a tie the first in the fixed
//     order (checker, neutral). Counts only tie after both groups got the
//     same number, so this is round robin as sign-ups arrive.
//  2. Denial slide: within the group, the flagged slide with the lowest count;
//     on a tie, the next in slide order after the group's last assigned
//     denial slide. With one flagged slide it is always that slide.
//  3. Increment both counters.
func Allocate(c *Counters, flagged []string) (sc.Group, string) {
	if c.Groups == nil {
		c.Groups = map[string]int64{}
	}
	if c.Denial == nil {
		c.Denial = map[string]map[string]int64{}
	}
	if c.LastDenial == nil {
		c.LastDenial = map[string]string{}
	}
	group := sc.Groups[0]
	for _, g := range sc.Groups[1:] {
		if c.Groups[string(g)] < c.Groups[string(group)] {
			group = g
		}
	}
	g := string(group)
	if c.Denial[g] == nil {
		c.Denial[g] = map[string]int64{}
	}
	counts := c.Denial[g]

	denial := ""
	if len(flagged) > 0 {
		minCount := counts[flagged[0]]
		for _, id := range flagged[1:] {
			if counts[id] < minCount {
				minCount = counts[id]
			}
		}
		start := 0
		for i, id := range flagged {
			if id == c.LastDenial[g] {
				start = i + 1
				break
			}
		}
		for k := 0; k < len(flagged); k++ {
			id := flagged[(start+k)%len(flagged)]
			if counts[id] == minCount {
				denial = id
				break
			}
		}
		counts[denial]++
		c.LastDenial[g] = denial
	}
	c.Groups[g]++
	return group, denial
}
