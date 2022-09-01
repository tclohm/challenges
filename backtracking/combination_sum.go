package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	if len(candidates) == 0 {
		return result
	}

	sort.Ints(candidates)

	tmp := []int{}
	backtrack(candidates, tmp, target, &result, 0)
	return result
}

func backtrack(candidates []int, tmp []int, target int, result *[][]int, indexPosition int) {
	if indexPosition == len(candidates) || target < 0 {
		return
	}

	if target == 0 {
		r := append([]int{}, tmp...)
		*result = append(*result, r)
		return
	}

	for i := indexPosition ; i < len(candidates) ; i++ {
		t := append(tmp, candidates[i])
		backtrack(candidates, t, target - candidates[i], result, i)
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}

	fmt.Println("candidates for 7 ", combinationSum(candidates, 7))
	fmt.Println("candidates for 9 ", combinationSum(candidates, 9))
	fmt.Println("candidates for 6 ", combinationSum(candidates, 6))
	fmt.Println("candidates for 13", combinationSum(candidates, 13))

}