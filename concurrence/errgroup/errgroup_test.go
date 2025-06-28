package errgroup

import (
	"context"
	"crypto/md5"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestFirstErr(t *testing.T) {
	var eg errgroup.Group

	eg.Go(func() error {
		time.Sleep(5 * time.Second)
		t.Log("exec #1")
		return nil
	})

	eg.Go(func() error {
		time.Sleep(10 * time.Second)
		t.Log("exec #2")
		return errors.New("failed to exec #2")
	})

	eg.Go(func() error {
		time.Sleep(15 * time.Second)
		t.Log("exec #3")
		return nil
	})

	if err := eg.Wait(); err != nil {
		t.Log(err)
	} else {
		t.Log("success")
	}
}

func TestAllErr(t *testing.T) {
	var eg errgroup.Group
	var errs = make([]error, 3)

	eg.Go(func() error {
		time.Sleep(5 * time.Second)
		t.Log("exec #1")
		errs[0] = nil
		return nil
	})

	eg.Go(func() error {
		time.Sleep(10 * time.Second)
		t.Log("exec #2")
		errs[1] = errors.New("failed to exec #2")
		return errs[1]
	})

	eg.Go(func() error {
		time.Sleep(15 * time.Second)
		t.Log("exec #3")
		errs[2] = nil
		return nil
	})

	if err := eg.Wait(); err != nil {
		t.Log(errs)
	} else {
		t.Log("success")
	}
}

func TestPipelines(t *testing.T) {
	m, err := MD5All(context.Background(), ".")
	if err != nil {
		log.Fatal(err)
	}

	for k, sum := range m {
		t.Logf("%s:\t%x\n", k, sum)
	}
}

type result struct {
	path string
	sum  [md5.Size]byte
}

func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)

		return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})

	c := make(chan result)
	const numDigesters = 20

	for i := 0; i < numDigesters; i++ {
		g.Go(func() error {
			for path := range paths {
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				select {
				case c <- result{path: path, sum: md5.Sum(data)}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			log.Fatal(err)
		}
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		m[r.path] = r.sum
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return m, nil
}
