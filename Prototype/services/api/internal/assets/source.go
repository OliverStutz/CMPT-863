// Package assets reads content from a local directory (dev, CONTENT_DIR) or a
// GCS bucket (prod, CONTENT_BUCKET) behind one interface, and serves the
// public subset of it (section 12.4).
package assets

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// ErrNotExist is returned for a missing object.
var ErrNotExist = fs.ErrNotExist

// Object is one content file in memory.
type Object struct {
	Data    []byte
	ModTime time.Time
}

// Source reads content files by slash-separated relative path.
type Source interface {
	Read(ctx context.Context, p string) (*Object, error)
	// List returns the relative paths of files directly under dir.
	List(ctx context.Context, dir string) ([]string, error)
	Name() string
}

// CleanPath validates a relative content path: no absolute paths, no ".."
// segments, no backslashes or NUL bytes. It returns "" when invalid.
func CleanPath(p string) string {
	if p == "" || strings.ContainsAny(p, "\\\x00") || strings.HasPrefix(p, "/") {
		return ""
	}
	for _, seg := range strings.Split(p, "/") {
		if seg == ".." || seg == "." || seg == "" {
			return ""
		}
	}
	return path.Clean(p)
}

// Local reads from a directory.
type Local struct{ Root string }

func (l Local) Name() string { return "dir " + l.Root }

func (l Local) Read(_ context.Context, p string) (*Object, error) {
	cp := CleanPath(p)
	if cp == "" {
		return nil, ErrNotExist
	}
	full := filepath.Join(l.Root, filepath.FromSlash(cp))
	st, err := os.Stat(full)
	if err != nil || st.IsDir() {
		return nil, ErrNotExist
	}
	b, err := os.ReadFile(full)
	if err != nil {
		return nil, err
	}
	return &Object{Data: b, ModTime: st.ModTime()}, nil
}

func (l Local) List(_ context.Context, dir string) ([]string, error) {
	d := l.Root
	if dir != "" {
		cp := CleanPath(dir)
		if cp == "" {
			return nil, ErrNotExist
		}
		d = filepath.Join(l.Root, filepath.FromSlash(cp))
	}
	ents, err := os.ReadDir(d)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	var out []string
	for _, e := range ents {
		if !e.IsDir() {
			out = append(out, path.Join(dir, e.Name()))
		}
	}
	sort.Strings(out)
	return out, nil
}

// GCS reads from a bucket and keeps every object it read in memory (12.4: a
// content version is 15 to 30 MB, so this fits in 512 MiB).
type GCS struct {
	Bucket string
	client *storage.Client
	mu     sync.Mutex
	cache  map[string]*Object
	size   int64
	// MaxCacheBytes bounds the cache; beyond it objects are read each time.
	MaxCacheBytes int64
}

func NewGCS(ctx context.Context, bucket string) (*GCS, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &GCS{Bucket: bucket, client: c, cache: map[string]*Object{}, MaxCacheBytes: 200 << 20}, nil
}

func (g *GCS) Name() string { return "gs://" + g.Bucket }

func (g *GCS) Read(ctx context.Context, p string) (*Object, error) {
	cp := CleanPath(p)
	if cp == "" {
		return nil, ErrNotExist
	}
	g.mu.Lock()
	if o, ok := g.cache[cp]; ok {
		g.mu.Unlock()
		return o, nil
	}
	g.mu.Unlock()
	obj := g.client.Bucket(g.Bucket).Object(cp)
	r, err := obj.NewReader(ctx)
	if errors.Is(err, storage.ErrObjectNotExist) {
		return nil, ErrNotExist
	}
	if err != nil {
		return nil, err
	}
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	o := &Object{Data: b, ModTime: r.Attrs.LastModified}
	g.mu.Lock()
	if g.size+int64(len(b)) <= g.MaxCacheBytes {
		g.cache[cp] = o
		g.size += int64(len(b))
	}
	g.mu.Unlock()
	return o, nil
}

func (g *GCS) List(ctx context.Context, dir string) ([]string, error) {
	prefix := ""
	if dir != "" {
		prefix = strings.TrimSuffix(dir, "/") + "/"
	}
	it := g.client.Bucket(g.Bucket).Objects(ctx, &storage.Query{Prefix: prefix, Delimiter: "/"})
	var out []string
	for {
		a, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		if a.Name != "" {
			out = append(out, a.Name)
		}
	}
	sort.Strings(out)
	return out, nil
}
