package main

import "fmt"

type Solution struct {
	canindates []int
	target 	   int
	result 	   [][]int
}

func (self *Solution) Run(canindates []int, target int) [][]int {
	if len(canindates) == 0 {
		return self.result
	}

	self.canindates = canindates
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

	for i := start ; i < len(self.canindates) ; i++ {
		self.sum(append(current, self.canindates[i]), summation+self.canindates[i], i)
	}
}

func combinationSum(canindates []int, target int) [][]int {
	solution := Solution{}
	return solution.Run(canindates, target)
}



func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}