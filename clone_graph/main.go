package main

import "fmt"

type Node struct {
	Val int
	Neighbors []*Node
}

func dfs(node *Node, hashmap map[*Node]*Node) *Node {
	n, exist := hashmap[node]
	if exist {
		return n
	}

	var duplicate = &Node{Val: node.Val}
	hashmap[node] = duplicate

	for index := 0 ; index < len(node.Neighbors) ; index++ {
		neighbor := node.Neighbors[index]
		
		duplicate.Neighbors = append(duplicate.Neighbors, dfs(neighbor, hashmap))
	}
	return duplicate
}

func clone_graph(node *Node) *Node {
	if node == nil {
		return &Node{}
	}
	oldToNew := make(map[*Node]*Node)

	return dfs(node, oldToNew)
}

func main() {
	adjacency_list := [][]int{{2,4}, {1, 3}, {2, 4}, {1, 3}} 
	fmt.Println(adjacency_list)
}