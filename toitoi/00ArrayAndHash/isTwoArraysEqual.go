package main

import "fmt"

func isEqual(target, array []int) bool {
	ht := map[int]int{}
	for i := range target {
		n1, n2 := target[i], array[i]
		ht[n1] += 1
		ht[n2] -= 1
	}
	for _, value := range ht {
		if value != 0 { return false }
	}
	return true
}

func main() {
	fmt.Println(isEqual([]int{1,2,3,4}, []int{2,4,1,3}))
	fmt.Println(isEqual([]int{7}, []int{7}))
	fmt.Println(isEqual([]int{3,7,9}, []int{3,7,11}))

}
