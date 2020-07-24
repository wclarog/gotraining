package main

import "fmt"

func main() {
	var (
		ls1, ls2 *intSimpleList
		vs1, vs2 int
	)

	ls1 = nil

	ls1 = insert(ls1, 1)
	ls1 = insert(ls1, 2)
	ls1 = insert(ls1, 3)
	ls1 = insert(ls1, 4)

	vs1, _ = getFirst(ls1)
	fmt.Printf("%d\n", vs1)

	ls2 = copyList(ls1)
	ls1 = freeList(ls1)

	vs2, _ = getFirst(ls2)
	fmt.Printf("%d\n", vs2)
	fmt.Printf("%v\n", isMember(ls2, 3))
	fmt.Printf("%v\n", isMember(ls2, 23))
}