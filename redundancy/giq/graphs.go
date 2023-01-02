package main

import (
	"fmt"
	_"strconv"
)

type Node struct {
	Value int
	Neighbors []*Node
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
}