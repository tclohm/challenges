package main

import "fmt"

type stack struct {
	container [][]int
}

func (s *stack) length() int {
	return len(s.container)
}

func (s *stack) pop() (int, int) {
	length := len(s.container) - 1
	index, height := s.container[length][0], s.container[length][1]
	s.container = s.container[0:length]
	return index, height
}

func (s *stack) push(a []int) {
	s.container = append(s.container, a)
}

func (s *stack) lastHeight() int {
	return s.container[len(s.container) - 1][1]
}

func largestArea(heights []int) int {
	maxArea := 0
	stack := stack{}

	for i, h := range heights {
		
		start := i
		
		for stack.length() > 0 && stack.lastHeight() > h {
			index, height := stack.pop()
			maxArea = max(maxArea, height * (i - index))
			start = index
		}

		stack.push([]int{start, h})
	}

		for _, array := range stack.container {
			i, h := array[0], array[1]
			maxArea = max(maxArea, h * (len(heights) - i))
		}

	return maxArea
}

func main() {
	fmt.Println(largestArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestArea([]int{2,4}))
}