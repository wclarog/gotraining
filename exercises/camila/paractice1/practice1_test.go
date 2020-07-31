package paractice1

import "testing"

func TestInsert(t *testing.T) {
	var list List
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	if len(list) != 3 {
		t.Error("Length should be 3")
	}
}

func TestFirst(t *testing.T) {
	var list List
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	first, err := list.First()
	if first != 3 || err != nil{
		t.Error("The fist number should be 3")
	}
}

func TestFirstNoPanic(t *testing.T) {
	var list List
	_, err := list.First()
	if err == nil {
		t.Error("Should be returned an error")
	}
}

func TestDelete(t *testing.T) {
	var list List
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Delete()
	if len(list) != 2 {
		t.Error("Length should be 2")
	}

	first, err := list.First()
	if err != nil || first != 2 {
		t.Error("Error deleting first value")
	}
}

func TestBelongs(t *testing.T) {
	var list List
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	if !list.Belongs(3) {
		t.Error("3 should belong to the list")
	}

	if list.Belongs(5) {
		t.Error("5 should not be in the list")
	}
}

func TestCopy(t *testing.T) {
	var list List
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	copy := list.Copy()
	if &copy == &list {
		t.Error("Should be different lists")
	}
}

func TestFreeMemory(t *testing.T) {
	var list List
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.FreeMemory()
	if list != nil {
		t.Error("Should be nil")
	}
}