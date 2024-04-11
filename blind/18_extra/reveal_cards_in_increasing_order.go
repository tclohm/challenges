package main

import (
	"fmt"
	"sort"
)

type Queue struct {
	container []int
}

func Instantiate(items []int) Queue {
	var q = Queue{}
	for i, _ := range items {
		q.container = append(q.container, i)
	}
	return q
}

func (q *Queue) Pop() int {
	el := q.container[0]
	q.container = q.container[1:]
	return el
}

func (q *Queue) Push(item int) {
	q.container = append(q.container, item)
}

func (q *Queue) isEmpty() bool {
	return len(q.container) == 0
}

func revealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	result := make([]int, len(deck), len(deck))

	q := Instantiate(deck)

	for _, card := range deck {
		index := q.Pop()
		result[index] = card
		if !q.isEmpty() {
			q.Push(q.Pop())
		}
	}
	return result
}

func main() {
	fmt.Println(revealedIncreasing([]int{17,13,11,2,3,5,7}))
	fmt.Println(revealedIncreasing([]int{1,1000}))
}