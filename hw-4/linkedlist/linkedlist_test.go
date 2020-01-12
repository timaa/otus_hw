package linkedlist

import (
	"testing"
)

func TestList_PushBack(t *testing.T) {
	tests := []int{0, 1, 2}

	list := List{}
	for i := 0; i < len(tests); i++ {
		num := tests[i]
		list.PushBack(num)
		value, _ := list.Last().Value().(int)

		if num != value {
			t.Errorf("Got %d, expected %d ", value, num)
		}

	}
}

func TestList_PushFront(t *testing.T) {
	tests := []int{0, 1, 2}

	list := List{}
	for i := 0; i < len(tests); i++ {
		num := tests[i]
		list.PushFront(num)
		value, _ := list.First().Value().(int)

		if num != value {
			t.Errorf("Got %d, expected %d ", value, num)
		}
	}
}

func TestList_Remove(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushFront(4)
	elem := list.Last()
	listSize := list.Len()

	if listSize != 4 {
		t.Errorf("Got %d, expected %d ", list.Len(), 4)
	}

	list.Remove(elem.Prev())

	listSize = list.Len()

	if listSize != 3 {
		t.Errorf("Got %d, expected %d ", list.Len(), 3)
	}

	list = List{}
	list.PushFront(0)
	_, err := list.Remove(list.First())

	if err != nil && list.First() != nil {
		t.Errorf("Not remove element from begin. Got %v, expected %v ", list.First(), nil)

	}

	list = List{}
	list.PushBack(0)
	_, err = list.Remove(list.Last())

	if err != nil && list.First() != nil {
		t.Errorf("Not remove element from end.  Got %v, expected %v ", list.Last(), nil)
	}

	list = List{}
	list.PushFront(0)
	list2 := List{}
	list2.PushBack(3)

	item := list2.First()
	_, err = list.Remove(item)

	if err == nil {
		t.Errorf("Cannot remove element from other list.  Got %v, expected error ", err)
	}

}
