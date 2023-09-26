package main

import "fmt"

func combination(candidates []int, target int) [][]int {

	var result = make([][]int, 0)
	var current = make([]int, 0)
	var backtrack func (int, int, []int)

	backtrack = func (index int, sum int, current []int) {
		// check if sum === target
		// if true, append current to result
		if sum == target {
			result = append(result, append([]int{}, current...))
			return
		}
		// if sum is greater than target, return
		if sum > target {
			return
		}
		// the meat of the algo, move through numbers and call func
		// after func called, decrease currents length
		// index is the starting point
		for i := index ; i < len(candidates) ; i++ {
			current = append(current, candidates[i])
			backtrack(i, sum+candidates[i], current)
			// pop the last element off our current array
			current = current[:len(current) - 1]
		}
	}
	backtrack(0,0,current)
	return result
}

func main() {
	fmt.Println(combination([]int{2,3,6,7}, 7))
	fmt.Println(combination([]int{2,3,5}, 8))
	fmt.Println(combination([]int{2}, 1))
}