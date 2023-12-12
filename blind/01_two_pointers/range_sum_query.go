package main

import "fmt"

type NumberArray struct {
	prefix []int
}

func Constructor(nums []int) NumberArray {
	array := NumberArray{
		prefix: make([]int, len(nums)),
	}
	cur := 0
	for i := 0 ; i < len(nums) ; i++ {
		cur += nums[i]
		array.prefix[i] = cur
	}

	return array
}

func (self *NumberArray) sumQuery(left, right int) int {
	rightSum := self.prefix[right]
	var leftSum int
	if left > 0 {
		leftSum = self.prefix[left - 1]
	} else {
		leftSum = 0
	}
	return rightSum - leftSum
}

func main() {
	array := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(array.sumQuery(0, 3))
}