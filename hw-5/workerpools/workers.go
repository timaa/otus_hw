package workerpools

import (
	"sync"
)

type threadSafeCounter struct {
	mu      sync.Mutex
	counter int

	maxErrorCount int
}

func (t *threadSafeCounter) Inc() {
	t.mu.Lock()
	t.counter++
	t.mu.Unlock()
}

func (t *threadSafeCounter) Value() int {
	return t.counter
}
// Run ...
func Run(tasks []func() error, N int, M int) error {
	funcs := make(chan func() error, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(N)
	counter := &threadSafeCounter{maxErrorCount: M}

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			worker(funcs, counter)
		}()
	}

	for _, t := range tasks {
		funcs <- t
	}

	close(funcs)
	wg.Wait()

	return nil
}

func worker(funcs chan func() error, counter *threadSafeCounter) {
	for f := range funcs {
		err := f()
		if err != nil {
			counter.Inc()
		}
		if counter.counter >= counter.maxErrorCount {
			return
		}
	}
}
