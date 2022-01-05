package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	value int
}

func (element *Element) String() string {
	return strconv.Itoa(element.value)
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0, 0)
}

type Stack struct {
	elements []*Element
	length int
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.length], element)
	stack.length = stack.length + 1
}

func (stack *Stack) Pop() *Element {
	if stack.length == 0 {
		return nil
	}

	var length = len(stack.elements)
	var element *Element = stack.elements[length - 1]
	if length > 1 {
		stack.elements = stack.elements[:length - 1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.length = len(stack.elements)
	return element
}

func main() {
	var stack *Stack = &Stack{}
	stack.New()

	var element1 *Element = &Element{34}
	var element2 *Element = &Element{44}
	var element3 *Element = &Element{52}
	var element4 *Element = &Element{91}

	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)

	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())	
}