package main

import (
	"fmt"
	"sort"
)

func findContentChildren(children, cookies []int) int {
	// at most one cookie for one child
	// each child, i, has a greed factor. min size of a cookie that the child will be content with
	// each cookie, j, has a size of s[j]
	// if s[j] > g[i], the child, i, will be content
	// maximize the number of your content children

	fed := 0

	hashtable := make(map[int]int)

	for _, cookie := range cookies {
		hashtable[cookie]++
	}

	for _, greed_level := range children {
		if hashtable[greed_level] > 0 {
			hashtable[greed_level]--
			fed++
		}
	}
	return fed
}

func findContentChildrenSorted(children, cookies[]int) int {
	sort.Ints(children)
	sort.Ints(cookies)

	var child_index, cookie_index = 0, 0

	for child_index < len(children) {
		for cookie_index < len(cookies) && children[child_index] > cookies[cookie_index] {
			cookie_index++
		}
		if cookie_index == len(cookies) {
			break
		}
		child_index++
		cookie_index++
	}
	return child_index
}

func main() {
	fmt.Println(findContentChildrenSorted([]int{1,2,3}, []int{1,1}))
	fmt.Println(findContentChildrenSorted([]int{1,2}, []int{1,2,3}))
	fmt.Println(findContentChildrenSorted([]int{1,2,3}, []int{1,2,3}))
}