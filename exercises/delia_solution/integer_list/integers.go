package main

import "fmt"

type IntegerList struct {
	list []int
}

func (l *IntegerList) insertItemAtFirst(n int) {
	b := make([]int, 1, len(l.list)+1)
	b[0] = n
	l.list = append(b, l.list...)
}

func (l *IntegerList) removeFirstItem() {
	if len(l.list) > 0 {
		l.list = l.list[1:]
	}
}

func (l *IntegerList) checkItemExists(i int) bool {
	for _, item := range l.list {
		if item == i {
			return true
		}
	}
	return false
}

func (l *IntegerList) copyList() []int {
	var cp = make([]int, len(l.list))
	copy(cp, l.list)
	return cp
}

func (l *IntegerList) releaseList() {
	l.list = nil
}

func main() {
	var b IntegerList
	b.list = []int{4, 6, 8, 10}
	b.insertItemAtFirst(2)
	fmt.Println("After insert item at first: ", b.list)
	b.removeFirstItem()
	fmt.Println("After removing first item", b.list)
	fmt.Println("Check if item 4 exists: ", b.checkItemExists(4))
	fmt.Println("Check if item 20 exists: ", b.checkItemExists(20))
	fmt.Println("Items copied", b.copyList())
}
