package main

import "fmt"

type IntegerList struct{
	list []int
}


func (l *IntegerList) insertItemAtFirst(n int) {
	b := make([]int, 1, len(l.list) + 1)
	b[0] = n
	l.list = append(b, l.list...)
}

func (l *IntegerList) removeFirstItem() {
	if len(l.list) > 0 {
		l.list = l.list[1:]
	}
}

func main() {
	var b IntegerList
	b.list = []int {4, 6, 8, 10}
	b.insertItemAtFirst(2)
	fmt.Println(b.list)
	b.removeFirstItem()
	fmt.Println(b.list)
}
