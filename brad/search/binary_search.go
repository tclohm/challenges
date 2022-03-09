package main

import "fmt"

func binary_search(items []int, item int) bool {
	if len(items) == 0 {
		return false
	}

	midpoint := len(items) / 2

	if items[midpoint] == item {
		return true
	}

	if item < items[midpoint] {
		return binary_search(items[:midpoint], item)
	}

	return binary_search(items[midpoint + 1:], item)
}

func main() {
	list := []int{0, 1, 2, 8, 13, 17, 19, 32, 42}
	fmt.Println(binary_search(list, 3))
	fmt.Println(binary_search(list, 1))
	fmt.Println(binary_search(list, 42))
	fmt.Println(binary_search(list, 50))

}