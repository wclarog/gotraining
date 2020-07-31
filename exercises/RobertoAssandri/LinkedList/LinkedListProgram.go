package main

import "fmt"

const emptyListMsg = "list has no elements\n"

type Node struct {
	next *Node
	value int
}

type List struct {
	first *Node
	last *Node
}

func InsertFirst(list *List, number int){
	node := Node{
		value: number,
	}

	if list.first == nil{
		list.last = &node
	} else {
		node.next = list.first
	}

	list.first = &node
}

func PrintList(list List){
	if list.first != nil{
		PrintListRec(*list.first)
	} else {
		fmt.Print(emptyListMsg)
	}
}

func PrintListRec(node Node){
	if node.next == nil{
		fmt.Printf("%d\n", node.value)
	}else if node.next != nil{
		fmt.Printf("%d - ", node.value)
		PrintListRec(*node.next)
	}
}

func ReturnFirst(list List) (int, error){
	if list.first != nil{
		return list.first.value, nil
	}

	return 0, fmt.Errorf(emptyListMsg)
}

func PrintElement(value int, err error){
	if err != nil{
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Printf("The element is %d\n", value)
	}
}

func DeleteFirst(list *List){
	if list.first != nil{
		list.first = list.first.next
	} else{
		fmt.Print(emptyListMsg)
	}
}

func Contains(list List, value int) (bool, error){
	if list.first == nil{
		return false, fmt.Errorf(emptyListMsg)
	} else {
		return ContainsRec(*list.first, value), nil
	}
}

func ContainsRec(node Node, value int) bool{
	if node.value == value{
		return true
	} else if node.next == nil{
		return false
	} else {
		return ContainsRec(*node.next, value)
	}
}

func CopyList(list List) List{
	if list.first == nil{
		return List{}
	} else {
		copyList := List{}
		copyList.first = CopyListRec(list.first, &copyList)
		return copyList
	}
}

func CopyListRec(original *Node, copyList *List) *Node {
	copyNode := Node {}
	copyNode.value = original.value
	if original.next != nil{
		copyNode.next = CopyListRec(original.next, copyList)
	} else {
		copyList.last = &copyNode
	}
	return &copyNode
}

func FreeList(list *List){
	if list.first != nil{
		node := list.first
		defer CleanNode(node)
		list.first = list.first.next
		FreeList(list)
	}
}

func CleanNode(node *Node){
	node.next = nil
}

func main() {
	list := List{}
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("************ Insert First ************")
	InsertFirst(&list, 10)
	InsertFirst(&list, 5)
	InsertFirst(&list, 3)
	InsertFirst(&list, 1)
	fmt.Println("Insert 10, 5, 3 and 1 in that order")
	fmt.Println("Print list")
	PrintList(list)
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")

	fmt.Println("")
	fmt.Println("")

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("************ Return First ************")
	list2 := List{}
	fmt.Println("Return and print the first element of an empty list")
	first, err := ReturnFirst(list2)
	PrintElement(first, err)

	fmt.Println("Add 200 to be first element of the list and print it")
	InsertFirst(&list2, 200)
	first, err = ReturnFirst(list2)
	PrintElement(first, err)

	fmt.Println("Add 5 to be first element of the list and print it")
	InsertFirst(&list2, 5)
	first, err = ReturnFirst(list2)
	PrintElement(first, err)

	fmt.Println("Print the list with 5 and 200")
	PrintList(list2)

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")

	fmt.Println("")
	fmt.Println("")

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("************ Delete First ************")

	list3 := List{}
	fmt.Println("Try deleting the first element in an empty list")
	DeleteFirst(&list3)
	fmt.Println("Add 100 and 200 to the list")
	InsertFirst(&list3, 100)
	InsertFirst(&list3, 200)
	fmt.Println("Print the list")
	PrintList(list3)
	fmt.Println("Delete the first element (200)")
	DeleteFirst(&list3)
	fmt.Println("Print the list, it should only display 100")
	PrintList(list3)

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")

	fmt.Println("")
	fmt.Println("")

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("******* List Contains Element ********")

	containsElement, contErr := Contains(list, 8)
	if contErr == nil{
		if containsElement{
			fmt.Printf("The list contains %d\n", 8)
		} else {
			fmt.Printf("The list does not contains %d\n", 8)
		}
	}

	containsElement, contErr = Contains(list, 10)
	if contErr == nil{
		if containsElement{
			fmt.Printf("The list contains %d\n", 10)
		} else {
			fmt.Printf("The list does not contains %d\n", 10)
		}
	}

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")

	fmt.Println("")
	fmt.Println("")

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("************* Copy List **************")
	fmt.Println("Create a copy of the first list (1 - 3 - 5 - 10)")
	list4 := CopyList(list)
	fmt.Println("Print the first list")
	PrintList(list)
	fmt.Println("Print the copy list, they should be identical")
	PrintList(list4)
	fmt.Println("Add 1000 to the copy list")
	InsertFirst(&list4, 1000)
	fmt.Println("Print the first list")
	PrintList(list)
	fmt.Println("Print the copy list, should be the same but starting with 1000")
	PrintList(list4)
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")

	fmt.Println("")
	fmt.Println("")

	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("************ Clean List **************")
	fmt.Println("Print a list with elements (1000 - 1 - 3 - 5 - 10)")
	PrintList(list4)
	fmt.Println("Free the list")
	FreeList(&list4)
	fmt.Println("Print the list, it should be empty")
	PrintList(list4)
	fmt.Println("**************************************")
	fmt.Println("**************************************")
	fmt.Println("**************************************")

	// Esto lo hice con punteros, que me juzgue la historia
}
