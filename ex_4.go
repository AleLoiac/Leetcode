package main

import (
	"container/list"
	"fmt"
)

func main() {

	l1 := list.New()
	l2 := list.New()

	l1.PushBack(1)
	l1.PushBack(3)
	l1.PushBack(11)

	l2.PushBack(1)
	l2.PushBack(7)

	l := list.New()

	l = mergeTwoLists(l1, l2)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}

}

func mergeTwoLists(list1 *list.List, list2 *list.List) *list.List {

	result := list.New()

	x := list1.Front()
	y := list2.Front()

	for x != nil && y != nil {

		val1 := x.Value.(int)
		val2 := y.Value.(int)

		if val1 <= val2 {
			result.PushBack(val1)
			x = x.Next()
		} else {
			result.PushBack(val2)
			y = y.Next()
		}
	}

	for x != nil {
		result.PushBack(x.Value.(int))
		x = x.Next()
	}

	for y != nil {
		result.PushBack(y.Value.(int))
		y = y.Next()
	}

	return result

}
