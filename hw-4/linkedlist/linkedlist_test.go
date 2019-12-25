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

	list.Remove(*elem.Prev())

	listSize = list.Len()

	if listSize != 3 {
		t.Errorf("Got %d, expected %d ", list.Len(), 3)
	}
}
