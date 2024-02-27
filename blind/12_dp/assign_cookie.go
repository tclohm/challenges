package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g, s []int) int {
	// at most one cookie for one child
	// each child, i, has a greed factor. min size of a cookie that the child will be content with
	// each cookie, j, has a size of s[j]
	// if s[j] > g[i], the child, i, will be content
	// maximize the number of your content children

	fed := 0

	hashtable := make(map[int]int)

	for _, cookie := range s {
		hashtable[cookie]++
	}

	for _, greed_level := range g {
		if hashtable[greed_level] > 0 {
			hashtable[greed_level]--
			fed++
		}
	}
	return fed
}

func findContentChildrenSorted(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var i, j = 0, 0

	for i < len(g) {
		for j < len(s) && g[i] > s[j] {
			j++
		}
		if j == len(s) {
			break
		}
		i++
		j++
	}
	return i
}

func main() {
	fmt.Println(findContentChildrenSorted([]int{1,2,3}, []int{1,1}))
	fmt.Println(findContentChildrenSorted([]int{1,2}, []int{1,2,3}))
	fmt.Println(findContentChildrenSorted([]int{1,2,3}, []int{1,2,3}))
}