package main 

import "fmt"


type Node struct {
	value int
	left *Node 
	right *Node
}
// arrays[i] = [parent, child, isLeft]
// isLeft == 1, the child is left
// isLeft == 0, the child is right
func create(arrays [][]int) Node {
	var nodes = map[int]*Node{}
	var children = map[int]bool{}

	for _, node := range arrays {
		var (
			value	= node[0]
			child = node[1]
			isLeft = node[2]
		)
		children[child] = true
		// if it doesnt exist, create
		if _, ok := nodes[value]; !ok {
			nodes[value] = &Node{value: value} 
		}
		n := nodes[value]
		ch := &Node{value: child}
		nodes[child] = ch
		if isLeft == 1 {
			n.left = ch		
		} else {
			n.right = ch	
		}
	}

	for _, node := range arrays {
		var (
			value	= node[0]
		)
		if _, ok := children[value]; !ok {
			return *nodes[value]
		}
	}
	
	return Node{value: -1}
}

func main() {
	fmt.Println(create([][]int{
		{20,15,1},
		{20,17,0},
		{50,20,1},
		{50,80,0},
		{80,19,1},
		}))

	fmt.Println(create([][]int{
		{1,2,1},
		{2,3,0},
		{3,4,1},
		}))
}
