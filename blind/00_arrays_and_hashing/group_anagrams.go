package main

import (
	"fmt"
	"sort"
	"strings"
)

func mySolution(strs []string) [][]string {
	ht := make(map[string][]string)
	results := make([][]string, 0)

	for _, word := range strs {
		sorted := sortString(word)
		ht[sorted] = append(ht[sorted], word)
	}

	for _, v := range ht {
		results = append(results, v)
	}

	return results
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	ht := make(map[[26]int][]string)
	for _, word := range strs {
		var count [26]int
		for _, r := range word {
			count[r - 'a']++
		}
		ht[count] = append(ht[count], word)
	}
	result := make([][]string, len(ht))
	index := 0
	for _, v := range ht {
		result[index] = v
		index++
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
}