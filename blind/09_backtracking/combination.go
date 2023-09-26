package main

import "fmt"

func combination(candidates []int, target int) [][]int {
	var result = make([][]int, 0)
	var current = make([]int, 0)
	var backtrack func (int, int, []int)
	backtrack = func (index int, sum int, current []int) {
		if sum == target {
			result = append(result, append([]int{}, current...))
			return
		}
		if sum > target {
			return
		}
		for i := index ; i < len(candidates) ; i++ {
			current = append(current, candidates[i])
			backtrack(i, sum+candidates[i], current)
			// move the needle forward
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