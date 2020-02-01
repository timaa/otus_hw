package main

import (
	"errors"
	"fmt"
	"github.com/timaa/otus_hw/hw-5/workerpools"
)

func main() {
	funcs:= []func() error {f1, f2, f3, f4}
	workerpools.Run(funcs, 4, 4)
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
	return errors.New("f3 err")
}

func f4() error{
	fmt.Println("f4")
	return errors.New("f4 err")
}