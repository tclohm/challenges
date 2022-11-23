package main

import "fmt"

type Set map[int]bool

func lcs(nums []int) int {
	hm := make(Set)

	for _, i := range nums {
		hm[i] = true
	}

	longest := 0

	for _, n := range nums {
		_, exist := hm[n-1]; if !exist {
			length := 0
			// while
			for {
				_, exist := hm[n+length]; if exist {
					length += 1
					continue
				} else {
					if longest < length {
						longest = length
					}
					break
				}
			}
		}
	}

	return longest	 
}

func main() {
	// 99? nope
	// 3? yes
	//	length++ 2? yes
	//	length++ 1? yes
	//	length++ 0? nope
	// 199? nope
	// 0? nope
	// 2? yes
	//	length++ 1? yes
	//	length++ 0? nope
	// 1? yes
	//	length++ 0? nope
	n := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(lcs(n))
}