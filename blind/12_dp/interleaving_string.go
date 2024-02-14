package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) { return false }

	var isInterleaved func(p1, p2, p3 int) bool
	isInterleaved = func(p1, p2, p3 int) bool {
		if p3 == len(s3) {
			if p1 == len(s1) && p2 == len(s2) { return true }
			return false
		}

		result := false

		if p1 < len(s1) && s1[p1] == s3[p3] {
			result = result || isInterleaved(p1 + 1, p2, p3 + 1)
		}
		if p2 < len(s2) && s2[p2] == s3[p3] {
			result = result || isInterleaved(p1, p2 + 1, p3 + 1)
		}

		return result
	}
	return isInterleaved(0, 0, 0)
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
}