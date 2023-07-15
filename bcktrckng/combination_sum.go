package main

import (
	"fmt"
	"sort"
)

type Solution struct {
	candidates []int
	target 	   int
	result 	   [][]int
}

func (self *Solution) Run(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return self.result
	}

	self.candidates = candidates
	self.target = target
	self.sum([]int{}, 0, 0)
	return self.result
}

func (self *Solution) sum(current []int, summation, start int) {
	if summation >= self.target {
		if summation == self.target {
			cpy := make([]int, len(current))
			copy(cpy, current)
			self.result = append(self.result, cpy)
		}
		return
	}

	for i := start ; i < len(self.candidates) ; i++ {
		self.sum(append(current, self.candidates[i]), summation+self.candidates[i], i)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	solution := Solution{}
	return solution.Run(candidates, target)
}

func combSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0)
	sort.Ints(candidates)

	var backtrack func(index int, currentSum int, current []int)
	backtrack = func(index int, currentSum int, current []int) {
		if currentSum == target {
			result = append(result, append([]int{}, current...))
			return
		}

		if currentSum > target {
			return
		}

		for i := index ; i < len(candidates) ; i++ {
			if i > index && candidates[i] == candidates[i - 1] {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i + 1, currentSum+candidates[i], current)
			current = current[:len(current) - 1]
		}
	}
	backtrack(0, 0, current)
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combSum([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combSum([]int{2, 5, 2, 1, 2}, 5))
}