package main

import "fmt"

type Stack struct {
	Box []int
}

func (s *Stack) Push(item int) {
	s.Box = append(s.Box, item)
}

func (s *Stack) Pop() int {
	length := len(s.Box)
	if length == 0 { return -1 }
	item := s.Box[length - 1]
	s.Box = s.Box[:length - 1]
	return item
}

func (s *Stack) Length() int {
	return len(s.Box)
}

// topological sort, inDegree

// DAG - Direct Acyclical Graph
type doublearray [][]int

func canFinish(n int, prerequisites doublearray) bool {
	var inDegree = make([]int, n)
	var al = make([][]int, n)
	for i := 0 ; i < len(prerequisites) ; i++ {
		var pre, next = prerequisites[i][0], prerequisites[i][1]
		inDegree[pre] += 1
		al[next] = append(al[next], pre)
	}

	var s = Stack{}
	for i := 0 ; i < len(inDegree) ; i++ {
		if inDegree[i] == 0 {
			s.Push(i)
		}
	}

	var count = 0
	for s.Length() > 0 {
		var current = s.Pop()
		count++
		var adjacent = al[current]

		for i := 0 ; i < len(adjacent) ; i++ {
			var next = adjacent[i]
			inDegree[next] -= 1
			if inDegree[next] == 0 {
				s.Push(next)
			}
		}
	}
	return count == n
}

func main() {
	prereqs := [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}
	fmt.Println(canFinish(6, prereqs))
}