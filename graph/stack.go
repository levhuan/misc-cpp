package stack

import (
	"fmt"
)

var ()

const (
	MAX_STACK   = 100
	GROW_FACTOR = 2
)

type Element interface{}

/*
 * Stack object
 */
type Stack struct {
	element []*Element
	next    *Stack
}

func New() *Stack {
	return &Stack{make([]*Element, MAX_STACK, GROW_FACTOR*MAX_STACK), 0}
}

func (s *Stack) push(elem *Element) *Element {
	if s.stackSize < cap(s.element) {
		s.stackSize += 1
		s.element[s.stackSize] = elem
		return elem
	}
	return nil
}

func (s *Stack) pop() *Element {
	fmt.Println("s.stackSize: ", s.stackSize)
	defer func() {
		s.stackSize = s.stackSize - 1
		fmt.Println("Deferred s.stackSize: ", s.stackSize)
	}()
	return s.element[s.stackSize]
}
