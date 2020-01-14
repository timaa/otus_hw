package worker_pools

import (
	"fmt"
	"sync"
	"time"
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

func (t threadSafeCounter) Value() int {
	return t.counter
}

func Run(tasks []func() error, N int, M int) error {
	funcs := make(chan func() error, 1)
	counter := &threadSafeCounter{maxErrorCount: M}

	for i := 0; i < N; i++ {
		go worker(i, funcs, counter)
	}

	for j := 0; j < M; j++ {
		funcs <- tasks[j]
	}

	time.Sleep(1 * time.Second)

	return nil
}

func worker(id int, funcs chan func() error, counter *threadSafeCounter) {

	fmt.Printf("counter=%d \n", counter.Value())
	for f := range funcs {
		fmt.Printf("Worker with id %d \n", id)
		err := f()
		if err != nil {
			counter.Inc()
			if counter.counter >= counter.maxErrorCount {
				close(funcs)
			}
		}
	}
	fmt.Printf("closed %d goroutine \n", id)
}
