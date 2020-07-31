package main

import "errors"

type intListNode struct {
	value int
	next *intListNode
}

type intList struct {
	nodes *intListNode
}

func (p *intList) insert(value int) {
	node := &intListNode{
		value: value,
		next: p.nodes,
	}
	p.nodes = node
}

func (p *intList) getFirst() (int, error) {
	if p.nodes == nil {
		return 0, errors.New("getFirst failed due to empty list")
	}

	return p.nodes.value, nil
}

func (p *intList) deleteFirst() error {
	if p == nil {
		return errors.New("deleteFirst failed due to empty list")
	}

	p.nodes = p.nodes.next
	return nil
}

func (p *intList) isMember(value int) bool {
	for current := p.nodes; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}

	return false
}

func (p *intList) copyList() *intList {
	new := &intList{
		nodes: nil,
	}
	currentNew := &(new.nodes)

	for current := p.nodes; current != nil; current = current.next {
		nodeNew := &intListNode{
			value:   current.value,
			next: nil,
		}
        *currentNew = nodeNew
		currentNew = &nodeNew.next
	}

	return new
}

func (p *intList) freeList() {
	p.nodes = nil
}
