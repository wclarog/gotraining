package main

import "fmt"

func main() {
	var (
		l1, l2 *intList
		v1, v2 int
	)

	l1 = &intList{
		nodes: nil,
	}

	l1.insert(1)
	l1.insert(2)
	l1.insert(3)
	l1.insert(4)

	v1, _ = l1.getFirst()
	fmt.Printf("%d\n", v1)

	l2 = l1.copyList()
	l1.freeList()
	v2, _ = l2.getFirst()
	fmt.Printf("%d\n", v2)
	fmt.Printf("%v\n", l2.isMember(3))
	fmt.Printf("%v\n", l2.isMember(23))
}