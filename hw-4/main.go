package main

import (
	"fmt"
	"github.com/timaa/otus_hw/hw-4/linkedlist"
)

func main() {

	//l := linkedlist.List{}
	//l.PushFront(1)
	//l.Remove(l.First())
	//fmt.Println(l.First())


	list := linkedlist.List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushFront(4)
	elem:=list.Last()
	fmt.Println(list.ToString())

	list.Remove(elem.Prev())
	fmt.Println(list.ToString())

}
