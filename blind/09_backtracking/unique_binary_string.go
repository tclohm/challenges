package main

import (
	"fmt"
)

func findDifferentBinaryString(nums []string) string {
	var set = make(map[string]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	var length = len(nums[0])

	combination, visited := make([]byte, length), make([]bool, length)

	var result string

	var backtrack func (index int) bool
	backtrack = func (index int) bool {
		if index == length {
			if _, ok := set[string(combination)]; !ok {
				result = string(combination)
				return true
			}
			return false
		}
		
		if visited[index] == false {
			visited[index] = true

			combination[index] = '0'

			if backtrack(index + 1) {
				return true
			}

			combination[index] = '1'

			if backtrack(index + 1) {
				return true
			}

			visited[index] = false
		}

		return false
	}

	backtrack(0)
	return result
}

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01","10"}))
	fmt.Println(findDifferentBinaryString([]string{"00","01"}))
	fmt.Println(findDifferentBinaryString([]string{"111","011","001"}))
}