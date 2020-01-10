package linkedlist

// List structure
type List struct {
	listSize int
	first    *Item
	last     *Item
}

// Len function return number of elements
func (l *List) Len() int {
	return l.listSize
}

//First function return first element in list
func (l *List) First() *Item {
	return l.first
}

//Last function return last element in list
func (l *List) Last() *Item {
	return l.last
}

//PushFront function insert elem in front
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

//PushBack function insert elem to start
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

// Remove function  remove element
func (l *List) Remove(i Item) {
	if i.prev != nil {
		i.prev.next = i.next

	}

	if i.next != nil {
		i.next.prev = i.prev
	}

	i.list = nil
	i.next = nil
	i.prev = nil
	l.listSize--
}

//Item structure
type Item struct {
	list  *List
	value interface{}
	prev  *Item
	next  *Item
}

// Value return value of element
func (i *Item) Value() interface{} {
	return i.value
}

//Next function: return next elem
func (i *Item) Next() *Item {
	return i.next
}

//Prev function: return previous elem
func (i *Item) Prev() *Item {
	return i.prev
}
