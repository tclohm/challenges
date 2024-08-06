package main

import "fmt"

func minHeightShelves(books [][]int, shelfWidth int) int {
	var cache = map[int]int{}

	var helper func(index int) int
	helper = func(index int) int {
		if index == len(books) {
			return 0
		}
		if value, exist := cache[index]; exist {
			return value
		}
		currentWidth := shelfWidth
		maxHeight := 0
		cache[index] = 1000
		for i := index ; i < len(books) ; i++ {
			width, height := books[i][0], books[i][1]
			if currentWidth < width { break }
			currentWidth -= width
			maxHeight = max(maxHeight, height)
			cache[index] = min(cache[index], helper(i + 1) + maxHeight)
		}
		return cache[index]
	}

	return helper(0)
}

func main() {
	fmt.Println(minHeightShelves(
		[][]int{
			{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2},
		}, 4))
	fmt.Println(minHeightShelves([][]int{{1,3},{2,4},{3,2}}, 6))	
}


func bottomup(books [][]int, shelfWidth int) int {
	var dp = make([]int, len(books) + 1)

	for i := len(books) - 1 ; i >= 0 ; i-- {
		currentWidth := shelfWidth
		maxHeight := 0
		dp[i] = 1000
		for j := i ; j < len(books) ; j++ {
			width, height := books[i][0], books[i][1]
			if currentWidth < width { break }
			currentWidth -= width
			maxHeight = max(maxHeight, height)
			dp[i] = min(dp[i], dp[j + 1] + maxHeight)
		}
	}
	return dp[0]
}
