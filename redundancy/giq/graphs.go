package main

import (
	"fmt"
	_"strconv"
)

type Queue struct {
	Container []int
}

func (q *Queue) Push(node int) {
	q.Container = append(q.Container, node)
}

func (q *Queue) Pop() int {
	node := q.Container[0]
	q.Container = q.Container[1:]
	return node
}

func (q *Queue) Length() int {
	return len(q.Container)
}

type Stack struct {
	Container []int
}

func (s *Stack) Push(node int) {
	s.Container = append(s.Container, node)
}

func (s *Stack) Pop() int {
	end := s.Length() - 1
	node := s.Container[end]
	s.Container = s.Container[0:end]
	return node
}

func (s *Stack) Length() int {
	return len(s.Container)
}

type Node struct {
	Value int
	Neighbors []*Node
}

type graph [][]int

func BFS(graph graph) []int {
	var q = Queue{}
	q.Push(0)
	values := []int{}
	seen := map[int]bool{}

	for q.Length() > 0 {
		var node = q.Pop()
		values = append(values, node)
		seen[node] = true

		var connections = graph[node]
		for i := 0 ; i < len(connections) ; i++ {
			var connection = connections[i]
			if _, ok := seen[connection]; !ok {
				q.Push(connection)
			}
		}
	}

	return values
}

func DFS(graph graph) []int {
	var s = Stack{}
	s.Push(0)
	values := []int{}
	seen := map[int]bool{}

	for s.Length() > 0 {
		var node = s.Pop()
		values = append(values, node)
		seen[node] = true

		var connections = graph[node]
		for i := 0 ; i < len(connections) ; i++ {
			var connection = connections[i]
			if _, ok := seen[connection]; !ok {
				s.Push(connection)
			}
		}
	}
	return values
}

func main() {
	var node0 = Node{Value: 0}
	var node1 = Node{Value: 1}
	var node2 = Node{Value: 2}
	var node3 = Node{Value: 3}
	var node4 = Node{Value: 4}
	var node5 = Node{Value: 5}

	node0.Neighbors = []*Node{&node3,}
	node1.Neighbors = []*Node{&node3,}
	node2.Neighbors = []*Node{&node3,}
	node3.Neighbors = []*Node{&node0,&node1,&node2,&node4}
	node4.Neighbors = []*Node{&node3,&node5,}
	node5.Neighbors = []*Node{&node4,}

	// graph representation

	// adjacency list
	adjList := [][]*Node{
		{&node3,},
		{&node3,},
		{&node3,},
		{&node0,&node1,&node2,&node4},
		{&node3,&node5,},
		{&node4,},
	}


	for _, l := range adjList {
		for i, n := range l {
			fmt.Println(i, n)
			break
		}
	}

	anotherAdjList := map[string][]string{
		"A": []string{"D",},
		"B": []string{"D",},
		"C": []string{"D",},
		"D": []string{"A", "B", "C", "E",},
		"E": []string{"D", "F",},
		"F": []string{"E",},
	}

	fmt.Println(anotherAdjList)
	// adjacency matrix

	adjMatrix := [][]int{
		{0, 0, 0, 1, 0, 0,},
		{0, 0, 0, 1, 0, 0,},
		{0, 0, 0, 1, 0, 0,},
		{1, 1, 1, 0, 1, 0,},
		{0, 0, 0, 1, 0, 1,},
		{0, 0, 0, 0, 1, 0,},
	}

	fmt.Println(adjMatrix)

	adjl := [][]int{
		{3, 1},
		{0},
		{3, 8},
		{0, 4, 5, 2},
		{3, 6},
		{3},
		{4, 7},
		{6},
		{2},
	}



	fmt.Println(BFS(adjl))
	fmt.Println(DFS(adjl))
}