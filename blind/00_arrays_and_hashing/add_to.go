package main

import "fmt"

func reverse(array []int) []int {
	r := make([]int, len(array))
	for i := 0 ; i < len(array) ; i++ {
		r[i] = array[len(array) - i - 1]
	}

	return r
}

func add(array []int, k int) []int {
	n := reverse(array)
	idx := 0
	for k > 0 {
		digit := k % 10
		
		if idx < len(n) {
			n[idx] += digit
		} else {
			n = append(n, digit)
		}

		carry := n[idx] / 10
		n[idx] = n[idx] % 10

		k = k / 10
		k += carry
		idx += 1
	}

	return reverse(n)
}

func main() {
	fmt.Println(add([]int{1,2,3}, 1444))
}