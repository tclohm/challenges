package main

import (
	"fmt"
	"sort"
)

func firstPermute(nums []int) [][]int {
	var perms [][]int

	var backtrack func([]int, []int)
	backtrack = func(nums []int, path []int) {
		if len(nums) == 0 {
			perms = append(perms, append([]int(nil), path...))
			return
		}

		for i := 0 ; i < len(nums) ; i++ {
			newPerms := append([]int(nil), nums[:i]...)
			newPerms = append(newPerms, nums[i + 1:]...)
			newPath := append([]int(nil), path...)
			newPath = append(newPath, nums[i])
			backtrack(newPerms, newPath)
		}
	}

	backtrack(nums, []int{})
	return perms
}

func in(n int, array []int) bool {
	for _, num := range array {
		if n == num { return true }
	}
	return false
}

func permute(nums []int) [][]int {
	
	var result = [][]int{}
	
	var backtrack func(currentPermutation []int)
	backtrack = func(currentPermutation []int) {
		if len(currentPermutation) == len(nums) {
			result = append(result, append([]int(nil), currentPermutation...))
			return
		}

		for _, num := range nums {
			if !in(num, currentPermutation) {
				currentPermutation = append(currentPermutation, num)
				backtrack(currentPermutation)
				currentPermutation = currentPermutation[:len(currentPermutation) - 1]
			}
		}
	}

	backtrack([]int{})
	return result
}

func main() {
	fmt.Println(generateThroughLexicographicOrder([]int{1,2,3}))
	fmt.Println(generateThroughLexicographicOrder([]int{8}))
	fmt.Println(generateThroughLexicographicOrder([]int{0,1}))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n - 1)

}
// A,B,C,D -> alphabetical order
// the intuition 
// if I can generate the "next" permutation in alpha (lexico) order,
// we can start in sorted first and  then generate the next one
func generateThroughLexicographicOrder(nums []int) [][]int {
	var nextPermutation func (nums []int) []int
	nextPermutation = func (nums []int) []int {
		var i = len(nums) - 2
		for i >= 0 && nums[i] >= nums[i + 1] {
			i -= 1
		}

		if i >= 0 {
			var j = len(nums) - 1
			for nums[j] <= nums[i] {
				j -= 1
			}
			nums[i], nums[j] = nums[j], nums[i]
		}

		var left, right = i + 1, len(nums) - 1

		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left += 1
			right -= 1
		}

		return nums
	}
	
	sort.Ints(nums)
	var n = len(nums)
	var result = [][]int{nums}
	for x := 1 ; x < factorial(n) ; x++ {
		var end = len(result) - 1
		var last = result[end]
		var nextPerm = nextPermutation(last)
		result = append(result, nextPerm)
	}

	return result
}
