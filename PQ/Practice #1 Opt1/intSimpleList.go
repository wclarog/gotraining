package main

import "errors"

type intSimpleList struct {
	value int
	next *intSimpleList
}

func insert(p *intSimpleList, value int) *intSimpleList {
	newNode := &intSimpleList{
		value: value,
		next: p,
	}

	return newNode
}

func getFirst(p *intSimpleList) (int, error) {
	if p == nil {
		return 0, errors.New("getFirst failed due to empty list")
	}

	return p.value, nil
}

func deleteFirst(p *intSimpleList) (*intSimpleList, error) {
	if p == nil {
		return nil, errors.New("deleteFirst failed due to empty list")
	}

	return p.next, nil
}

func isMember(p *intSimpleList, value int) bool {
	if p == nil {
		return false
	} else if p.value == value {
		return true
	} else {
		return isMember(p.next, value)
	}
}

func copyList(p *intSimpleList) *intSimpleList {
	if p == nil {
		return nil
	} else {
		return insert(copyList(p.next), p.value)
	}
}

func freeList(p *intSimpleList) *intSimpleList {
	for p != nil {
		p, _ = deleteFirst(p)
	}

	return p
}
