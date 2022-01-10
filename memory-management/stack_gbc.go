package main

import (
	"fmt"
	"sync"
)

type ReferenceCounter struct {
	num *uint32
	pool *sync.Pool
	removed *uint32
}

func newReferenceCounter() *ReferenceCounter {
	return &ReferenceCounter{
		num: new(uint32),
		pool: &sync.Pool{},
		removed: new(uint32),
	}
}

type Stack struct {
	references []*ReferenceCounter
	Count int
}

func (stack *Stack) New() {
	stack.references = make([]*ReferenceCounter, 0)
}

func (stack *Stack) Push(reference *ReferenceCounter) {
	stack.references = append(stack.references[:stack.Count], reference)
	stack.Count = stack.Count + 1
}

func (stack *Stack) Pop() *ReferenceCounter {
	if stack.Count == 0 {
		return nil
	}

	var length int = len(stack.references)
	var reference *ReferenceCounter = stack.references[length - 1]
	if length > 1 {
		stack.references = stack.references[:length - 1]
	} else {
		stack.references = stack.references[0:]
	}

	stack.Count = len(stack.references)
	return reference
}

func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var r1 *ReferenceCounter = newReferenceCounter()
	var r2 *ReferenceCounter = newReferenceCounter()
	var r3 *ReferenceCounter = newReferenceCounter()
	var r4 *ReferenceCounter = newReferenceCounter()

	stack.Push(r1)
	stack.Push(r2)
	stack.Push(r3)
	stack.Push(r4)

	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}