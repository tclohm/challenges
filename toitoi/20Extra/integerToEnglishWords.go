package main

import (
	"fmt"
	"math"
	"strings"
)

func integerToEnglish(num int) string {
	if num == 0 {
		return "Zero"
	}

	onesMap := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}

	tensMap := map[int]string{
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}

	var getString func(n int) string
	getString = func(n int) string {
		var res = []string{}
		var hundreds = int(math.Floor(float64(n / 100)))
		if hundreds != 0 {
			res = append(res, onesMap[hundreds], " Hundreds")
		}
		var lastTwo = n % 100
		if lastTwo >= 20 {
			var tens, ones = int(math.Floor(float64(lastTwo / 10))), lastTwo % 10
			res = append(res, tensMap[tens * 10])
			if ones != 0 {
				res = append(res, onesMap[ones])
			}
		} else if lastTwo != 0 {
			res = append(res, onesMap[lastTwo])
		}
		return strings.Join(res, " ")
	}
	
	var postfix = []string{"", "Thousand", "Million", "Billion"}
	var i = 0
	var arr = []string{}
	for num > 0 && i < len(postfix) {
		var digits = num % 1000
		var s = getString(digits) + postfix[i]
		if s != "" {
			arr = append(arr, s)
		}
		num = int(math.Floor(float64(num/1000)))
		i += 1
	}
	var res = ""
	for i := len(arr) - 1 ; i >= 0 ; i-- {
		res = res + arr[i]
	}

	return res
}

func main() {
	fmt.Println(integerToEnglish(123))
	fmt.Println(integerToEnglish(12345))
	fmt.Println(integerToEnglish(1234567))
}
