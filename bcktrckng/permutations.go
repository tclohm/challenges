package main

import "fmt"

// all possible permutations --- 3 choices * 2 choices * 1 choice = 6 permutations
// input: [1, 2, 3]
// output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

//	Decision tree
							 // 	1,     2,      3
							 //  2  3    1  3    1  2
						 	 // 3    2  3    1  2    1

// MARK: -- more of a math problem
func permute(nums []int) [][]int{
	n := len(nums)
	result := make([][]int, 0)
	current := make([]int, 0, n)
	visited := make(map[int]int)

	var backtrack func(index int)
	backtrack = func(index int) {
		// add current array to the result if the current length equals length of nums
		if len(current) == n {
			result = append(result, append([]int{}, current...))
		}

		for i := 0 ; i < n ; i++ {
			if visited[i] == 0 {
				visited[i]++
				current = append(current, nums[i])
				backtrack(i + 1)
				current = current[:len(current) - 1]
				visited[i]--
			}
		}
	}
	backtrack(0)
	return result
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{0,1}))
	fmt.Println(permute([]int{1}))
}