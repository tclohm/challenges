package main 

import "fmt"
// arrays[i] = [parent, child, isLeft]
// isLeft == 1, the child is left
// isLeft == 0, the child is right
func create(arrays [][]int) []int {
	var adjList = map[int][]int{}

	for _, node := range arrays {
		var (
			value	= node[0]
			child = node[1]
			isLeft = node[2]
		)
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
	
	var root = findRoot(adjList)
	fmt.Println(root)
	return []int{0}
}

func findRoot(adj map[int][]int) int {
	var children = map[int]bool{}
	for _, childs := range adj {
		for _, child := range childs {
			children[child] = true
		}		
	}

	for key, _ := range adj {
		if _, ok := children[key]; !ok {
			return key
		}
	}
	return -1
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
