package worker_pools

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

func TestRunGoodKeys(t *testing.T) {
	funcs:= []func() error {f1, f2, f3, f4}
	Run(funcs, 4, 4)

	numGoroutine := 4
	if runtime.NumGoroutine() -2 != numGoroutine {
		t.Errorf("Expected %d goroutine got %d", numGoroutine,  runtime.NumGoroutine() - 2)
	}
}

func TestRunBadKeys(t *testing.T) {
	funcs:= []func() error {f1_err, f2_err, f3_err, f4_err}
	Run(funcs, 4, 4)

	numGoroutineAfterMError := 0
	if runtime.NumGoroutine() - 2 != numGoroutineAfterMError {
		t.Errorf("Expected %d goroutine got %d", numGoroutineAfterMError,  runtime.NumGoroutine() - 2)
	}
}

func f1_err() error{
	fmt.Println("f1")
	return errors.New("f1 err")
}

func f2_err() error{
	fmt.Println("f2")
	return errors.New("f2 err")
}

func f3_err() error{
	fmt.Println("f3")
	return errors.New("f3 err")
}

func f4_err() error{
	fmt.Println("f4")
	return errors.New("f4 err")
}



func f1() error{
	fmt.Println("f1")
	return nil
}

func f2() error{
	fmt.Println("f2")
	return nil
}

func f3() error{
	fmt.Println("f3")
	return nil
}

func f4() error{
	fmt.Println("f4")
	return nil
}