package main

import (
	"fmt"
)

func findDifferentBinaryString(nums []string) string {
	var result string = ""
	for i, _ := range nums {
		if string(nums[i][i]) == "0" {
			result += "1" 
		} else {
			result += "0"
		}
	}
	return result
}

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01","10"}))
	fmt.Println(findDifferentBinaryString([]string{"00","01"}))
	fmt.Println(findDifferentBinaryString([]string{"111","011","001"}))
}