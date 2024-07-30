package main

import ("fmt")

func prefix(nums []int) []int {
	prefix := make([]int, len(nums) + 1, len(nums) + 1)
	
	for i, n := range nums {
		prefix[i + 1] = prefix[i] + n
	}	
	return prefix
}

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {

	n := NumArray{
		prefix: prefix(nums),
	}

	return n

}

func (this *NumArray) SumRange(left, right int) int {
	return this.prefix[right + 1] + this.prefix[left]
}

func main() {
	obj := Constructor([]int{-2,0,3,-5,2,-1})
	fmt.Println(obj.SumRange(0,2))
	fmt.Println(obj.SumRange(2,5))
	fmt.Println(obj.SumRange(0,5))
}
