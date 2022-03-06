package main

import (
	"fmt"
	"strconv"
)

// 233 base 10 === 11101001 base 2
// interpreted
// 2 * 10**2 + 3 * 10**1 + 3 * 10**0
//
// also intrepreted as
//
// 1 * 2**7 + 1 * 2**6 + 1 * 2**5 + 0 * 2**4 
// + 1 * 2**3 + 0 * 2**2 + 0 * 2**1 + 1 * 2**0

// integer greater than 0
// simple itration divide the decimal number by 2
// and keep track of remainder
// first division by gives value whether it's odd or even
// even value, remainder will be 0
// odd remainder will be 1

func convert_to_binary(decimal_number int) string {
	remainder_stack := []int{}

	for decimal_number > 0 {
		remainder := decimal_number % 2
		remainder_stack = append(remainder_stack, remainder)
		decimal_number = decimal_number / 2
	}

	binary_digits := ""
	
	for len(remainder_stack) > 0 {
		top := len(remainder_stack) - 1
		binary_digits += strconv.Itoa(remainder_stack[top])
		remainder_stack = remainder_stack[:top]
	}

	return binary_digits
}

func convert_to_base(decimal_number, base int) string {
	DIGITS := "0123456789abcdef"

	remainder_stack := []int{}

	for decimal_number > 0 {
		remainder := decimal_number % base
		remainder_stack = append(remainder_stack, remainder)
		decimal_number = decimal_number / base
	}

	new_digits := ""

	for len(remainder_stack) > 0 {
		top := len(remainder_stack) - 1
		index := remainder_stack[top]
		new_digits += string(DIGITS[index])
		remainder_stack = remainder_stack[:top]
	}

	return new_digits
}


func main() {
	fmt.Println("233 to binary is", convert_to_binary(233))
	fmt.Println("42 to binary is", convert_to_binary(42))
	fmt.Println("25 base 2 is", convert_to_base(25, 2))
	fmt.Println("25 base 16 is", convert_to_base(25, 16))
}