package main

import (
	"fmt"
	"sort"
	"math"
)

func minPushes(word string) int {
	counts := make([]rune, 26, 26)
	for _, r := range word {
		counts[r - 'a'] += 1
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	res := 0
	for i, n := range counts {
		if n > 0 {
			res += int(n) * int(math.Floor(float64(1 + i / 8)))
		}
	}

	return res 
}

func main() {
	fmt.Println(minPushes("abcde"))
	fmt.Println(minPushes("xyzxyzxyzxyz"))
	fmt.Println(minPushes("aabbccddeeffgghhiiiiii"))
}
