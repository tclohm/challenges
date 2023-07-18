package main

import "fmt"

type Node struct {
	val int
	neigbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil { return node }
	visited := map[*Node]*Node{}
	return clone(node, visited)
}

func close(node *Node, visited map[*Node]*Node) *Node {
	if _, ok := visited[node]; ok {
		return visited[node]
	}
	newNode := &Node{Val: node.val}
	visited[node] = newNode

	for _, neighbor := range node.neigbors {
		newNode.neigbors = append(newNode.neigbors, clone(neighbor, visited))
	}
	return newNode
}

func main() {
	adjacentList := [][]int{{2,4},{1,3},{2,4},{1,3}}
	buildgraph(adjacentList)
}