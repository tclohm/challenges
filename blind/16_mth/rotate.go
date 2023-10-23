package main

import "fmt"

func rotate(matrix [][]int) {
	for array := range matrix {
		fmt.Println(array)
	}
}

func main() {
	rotate([][]int{{1,2,3},{4,5,6},{7,8,9}})
	//rotate([][]int{{5,1,9,11},{2,4,8,10},{3,3,6,7},{15,14,12,16}})
}