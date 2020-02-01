package workerpools

import (
	"errors"
	"sync/atomic"
	"testing"
)

var threadSafeTaskCounter int64

func TestRunGoodKeys(t *testing.T) {
	threadSafeTaskCounter = 0
	funcs := []func() error{f1, f2, f3, f4}
	Run(funcs, 4, 4)

	finishCounter := atomic.LoadInt64(&threadSafeTaskCounter)
	if finishCounter != int64(len(funcs)) {
		t.Errorf("Expected %d task done  got %d", len(funcs), finishCounter)
	}
}

func TestRunBadKeys(t *testing.T) {
	threadSafeTaskCounter = 0
	funcs := []func() error{f1Err, f2Err, f1, f1, f1, f2Err, f3Err, f1, f2, f3Err, f4Err, f3, f4}
	n, m := 2, 2
	Run(funcs, n, m)

	if threadSafeTaskCounter > int64(m+n) {
		t.Errorf("Expected less than M + N= %d task done  got %d", m+n, threadSafeTaskCounter)
	}
}

func f1Err() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return errors.New("f1 err")
}

func f2Err() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return errors.New("f2 err")
}

func f3Err() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return errors.New("f3 err")
}

func f4Err() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return errors.New("f4 err")
}

func f1() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return nil
}

func f2() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return nil
}

func f3() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return nil
}

func f4() error {
	atomic.AddInt64(&threadSafeTaskCounter, 1)
	return nil
}
