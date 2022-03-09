package main

import "fmt"

func iterative_sum(numbers []int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sum_of(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + sum_of(numbers[1:])
}

func to_string(n, base int) string {
	CHAR_FOR_INT := "0123456789abcdef"

	if n < base {
		return string(CHAR_FOR_INT[n])
	}

	return to_string(n / base, base) + string(CHAR_FOR_INT[n % base])
}

func main() {
	arr := []int{1,3,5,7,9}
	fmt.Println(iterative_sum(arr))

	fmt.Println(sum_of(arr))

	fmt.Println(to_string(1453, 16))
}