package main 

import "fmt"


type Node struct {
	value int
	left int
	right int
}
// arrays[i] = [parent, child, isLeft]
// isLeft == 1, the child is left
// isLeft == 0, the child is right
func create(arrays [][]int) []int {
	var adjList = map[int][]int{}
	var children = map[int]bool{}

	for _, node := range arrays {
		var (
			value	= node[0]
			child = node[1]
			isLeft = node[2]
		)
		children[child] = true
		// if it doesnt exist, create
		if _, ok := adjList[value]; !ok {
			adjList[value] = make([]int, 2, 2)
		}
		if isLeft == 1 {
			adjList[value][0] = child
		} else {
			adjList[value][1] = child
		}
	}

	for _, node := range arrays {
		var (
			value	= node[0]
		)
		if _, ok := children[value]; !ok {
			fmt.Println("root:", value)
			break
		}
	}
	
	return []int{0}
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
