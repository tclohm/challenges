package main

import "fmt"

func top_k_frequency(numbers []int, k int) []int {
	var freq = make([][]int, 6, 6)
	var count = map[int]int{}

	for _, v := range numbers {
		if _, ok := count[v]; !ok {
			count[v] = 1
		} else {
			count[v] += 1
		}
	}

	for k, v := range count {
		freq[v] = append(freq[v], k)
	}

	res := []int{}
	
	for i := len(freq) - 1 ; i > 0 ; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	nums_2 := []int{1}

	fmt.Println(top_k_frequency(nums, 2))
	fmt.Println(top_k_frequency(nums_2, 1))
}