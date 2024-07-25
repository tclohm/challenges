package main

import (
	"fmt"
)

func groupAnagrams(strings []string) [][]string {
	var groups = map[string][]string{}
	var result = [][]string{}
	for _, s := range strings {
		var count = make([]byte, 26, 26)
		for _, c := range s {
			n := c - 'a'
			count[n] += 1
		}
		alpha := bytesToString(count)
		groups[alpha] = append(groups[alpha], s)
	}
	
	for _, v := range groups {
		result = append(result, v)
	}

	return result
}

func bytesToString(array []byte) string {
	var result = ""
	for _, n := range array {
		result += fmt.Sprintf("%d", n)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"act","pots","tops","cat","stop","hat"}))
	fmt.Println(groupAnagrams([]string{"x"}))
	fmt.Println(groupAnagrams([]string{""}))

}
