package main

import (
	"fmt"
)

func groupAnagrams(strings []string) [][]string {
	// var groups = map[string]int{}

	for _, s := range strings {
		var count = make([]byte, 26, 26)
		for _, c := range s {
			n := c - 'a'
			count[n] += 1
		}
		fmt.Println(count)
	}
	//	fmt.Println(groups)
	return [][]string{}
}

func main() {
	fmt.Println(groupAnagrams([]string{"act","pots","tops","cat","stop","hat"}))
	fmt.Println(groupAnagrams([]string{"x"}))
	fmt.Println(groupAnagrams([]string{""}))

}
