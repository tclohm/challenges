package main

import "fmt"

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



func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}