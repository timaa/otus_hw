package linkedlist

import (
	"fmt"
	"strings"
)

type List struct {
	listSize int
	first    *Item
	last     *Item
}

func (l *List) Len() int {
	return l.listSize
}

func (l *List) First() *Item {
	return l.first
}

func (l *List) Last() *Item {
	return l.last
}

func (l *List) PushFront(v interface{}) {
	item := &Item{value: v}
	if l.listSize > 0 {
		l.first.prev = item
		item.next = l.first
		l.first = item
	} else {
		l.last = item
		l.first = item
	}
	l.listSize++
}

func (l *List) PushBack(v interface{}) {
	item := &Item{value: v}
	if l.listSize > 0 {
		l.last.next = item
		item.prev = l.last
		l.last = item
	} else {
		l.last = item
		l.first = item
	}
	l.listSize++
}

func (l *List) Remove(i Item) {
	if i.prev != nil {
		i.prev.next = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	}

	l.listSize--
}

func (l List) String() string {
	curElem:=l.first
	str:=strings.Builder{}

	for ;curElem != nil; {
		str.WriteString(fmt.Sprintf("%v \n", curElem.Value()))
		curElem = curElem.Next()
	}
	return str.String()
}

type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Prev() *Item {
	return i.prev
}
