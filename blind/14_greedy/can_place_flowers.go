package main

import "fmt"

func canPlace(flowerbed []int, n int) bool {
	// no thing can be adjacent
	for i := 0 ; i < len(flowerbed) ; i++ {
		if flowerbed[i] == 1 { continue }
		if i - 1 >= 0 || i + 1 < len(flowerbed) {
			if flowerbed[i + 1] == 1 || flowerbed[i - 1] == 1 { continue }
		}
		flowerbed[i] = 1
		n -= 1
	}
	return n == 0
}

func main() {
	fmt.Println(canPlace([]int{1,0,0,0,1}, 1))
	fmt.Println(canPlace([]int{1,0,0,0,1}, 2))
}