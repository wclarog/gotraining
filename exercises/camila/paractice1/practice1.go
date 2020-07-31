package paractice1

import (
	"errors"
)

type List []int

func (l *List) Insert(x int) {
	*l = append([]int{x}, *l...)
}

func (l List) First() (int, error) {
	if l != nil && len(l) > 0 {
		v := l[0]
		return v, nil
	}
	return 0, errors.New("error: Empty list")
}

func (l *List) Delete() {
	if *l != nil && len(*l) > 1 {
		*l = (*l)[1:]
	}
}

func (l List) Belongs(x int) bool {
	for _, e := range l {
		if e == x {
			return true
		}
	}
	return false
}

func (l List) Copy() List {
	var dest List = make([]int, len(l))
	copy(dest, l)
	return dest
}

func (l *List) FreeMemory() {
	*l = nil
}