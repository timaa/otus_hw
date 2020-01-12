package linkedlist

import "errors"

// List structure
type List struct {
	listSize int
	first    *Item
	last     *Item
}

// Len function returns number of list elements.
func (l *List) Len() int {
	return l.listSize
}

//First function return first element in list.
func (l *List) First() *Item {
	return l.first
}

//Last function return last element in list.
func (l *List) Last() *Item {
	return l.last
}

//PushFront function insert element in front.
func (l *List) PushFront(v interface{}) {
	item := &Item{value: v, list: l}
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

//PushBack function insert element to end.
func (l *List) PushBack(v interface{}) {
	item := &Item{value: v, list: l}
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

// Remove function  remove element from list.
func (l *List) Remove(i *Item) (*Item, error) {
	if l != i.list {
		err := errors.New("Item not from this list")
		return i, err
	}
	if i.prev != nil {
		i.prev.next = i.next
	} else {
		l.first = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	} else {
		l.last = i.prev
	}

	i.list = nil
	i.next = nil
	i.prev = nil
	l.listSize--
	return i, nil
}

//Item structure
type Item struct {
	list  *List
	value interface{}
	prev  *Item
	next  *Item
}

// Value function return value of element.
func (i *Item) Value() interface{} {
	return i.value
}

//Next function return next element.
func (i *Item) Next() *Item {
	return i.next
}

//Prev function return previous element.
func (i *Item) Prev() *Item {
	return i.prev
}
