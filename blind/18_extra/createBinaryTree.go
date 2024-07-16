package main 

import "fmt"
// arrays[i] = [parent, child, isLeft]
// isLeft == 1, the child is left
// isLeft == 0, the child is right
func create(arrays [][]int) []int {
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
