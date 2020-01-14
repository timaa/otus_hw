package main

import (
	"errors"
	"fmt"
	"github.com/timaa/otus_hw/hw-5/worker_pools"
	"time"
)

func main() {
	funcs:= []func() error {f1, f2, f3, f4}
	worker_pools.Run(funcs, 4, 4)
	time.Sleep(1*time.Second)
}

func f1() error{
	fmt.Println("f1")
	return errors.New("f1 err")
}

func f2() error{
	fmt.Println("f2")
	return errors.New("f2 err")
}

func f3() error{
	fmt.Println("f3")
	return errors.New("f3 err")
}

func f4() error{
	fmt.Println("f4")
	return errors.New("f4 err")
}