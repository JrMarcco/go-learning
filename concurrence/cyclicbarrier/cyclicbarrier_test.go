package cyclicbarrier

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
	"log"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	b     cyclicbarrier.CyclicBarrier
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	ctx := context.Background()
	if err := h2o.semaH.Acquire(ctx, 1); err != nil {
		log.Fatal(err)
	}

	releaseHydrogen()

	if err := h2o.b.Await(ctx); err != nil {
		log.Fatal(err)
	}
	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	ctx := context.Background()
	if err := h2o.semaO.Acquire(ctx, 1); err != nil {
		log.Fatal(err)
	}

	releaseOxygen()

	if err := h2o.b.Await(ctx); err != nil {
		log.Fatal(err)
	}

	h2o.semaO.Release(1)
}

func TestCyclicBarrier(t *testing.T) {
	h2o := New()

	N := 100
	ch := make(chan string, 3*N)

	releaseHydrogen := func() {
		ch <- "H"
	}

	releaseOxygen := func() {
		ch <- "O"
	}

	var wg sync.WaitGroup
	wg.Add(3 * N)

	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}

	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	wg.Wait()

	if len(ch) != 3*N {
		log.Fatal("wrong total")
	}

	s := make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)

		water := s[0] + s[1] + s[2]

		if water != "HHO" {
			log.Fatal("not water")
		}
		t.Log(water)
	}
}
