package main

import "fmt"

// not a sliding window
func largest(nums string) string {
	var result uint8
	for i := 0 ; i < len(nums) - 2 ; i++ {
		if nums[i] == nums[i + 1] && nums[i + 1] == nums[i + 2] {
			if nums[i] > result {
				result = nums[i]
			}
		}
	}
	
	if result != 0 {
		s := string(result)
		return fmt.Sprintf("%v%v%v", s, s, s)
	}

	return ""
}

func main() {
	fmt.Println(largest("6777133339"))
	fmt.Println(largest("2300019"))
	fmt.Println(largest("423522338"))
}