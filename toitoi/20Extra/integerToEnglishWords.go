package main

import "fmt"

func integerToEnglish(num int) string {
	if num == 0 {
		return "Zero"
	}

	ones := map[int]string{
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

	tens := map[int]string{
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}

	var getString func(n int)
	getString = func(n int) {
		return
	}
	
	var postfix = []string{"", "Thousand", "Million", "Billion"}
	var i = 0
	var arr = []string{}
	for num > 0 {
		var digits = num % 1000
		var s = getString(digits) + postfix[i]
		if s != "" {
			arr = append(arr, s)
		}
		num = math.Floor(float64(num/1000))
		i += 1
	}
	var res = ""
	for i := len(arr) ; i >= 0 ; i-- {
		res = res + arr[i]
	}

	return res
}

func main() {
	fmt.Println(integerToEnglish(123))
	fmt.Println(integerToEnglish(12345))
	fmt.Println(integerToEnglish(1234567))
}
