package main

import "fmt"

func gcdOfStrings(s, t string) string {
	if s + t != t + s {
		return ""
	}
	x := gcd(len(s), len(t))
	return s[:x]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}