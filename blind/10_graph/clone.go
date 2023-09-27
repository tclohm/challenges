package main

import "fmt"

type Node struct {
	Value int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil { return node }
	visited := map[*Node]*Node{}
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if _, ok := visited[node]; ok {
		return visited[node]
	}

	new := &Node{Value: node.Value}
	visited[node] = new

	for _, neighbor := range node.Neighbors {
		new.Neighbors = append(new.Neighbors, clone(neighbor, visited))
	}
	return new
}

func main() {}