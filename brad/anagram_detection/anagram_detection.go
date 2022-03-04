package main

import (
	"strings"
	"fmt"
	"sort"
)

func anagram_checking_off(string1, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}

	to_check_off := strings.Split(string2, "")

	for _, char := range string1 {
		s := string(char)
		for index, other_char := range to_check_off {
			if s == other_char {
				to_check_off[index] = ""
				break
			}
		}
	}

	for i := 0 ; i < len(to_check_off) ; i++ {
		if to_check_off[i] != "" {
			return false
		}
	}
	return true
}

func anagram_sort_and_compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// O(n)
	arr1 := strings.Split(s1, "")
	arr2 := strings.Split(s2, "")

	// O(n log n)
	sort.Strings(arr1)
	sort.Strings(arr2)

	// O(n)
	for i := 0 ; i < len(arr1) ; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func anagram_count_compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	alpha1 := make([]int, 26)
	alpha2 := make([]int, 26)

	for _, char := range s1 {
		pos := char - 'a'
		alpha1[pos] += 1
	}

	for _, char := range s2 {
		pos := char - 'a'
		alpha2[pos] += 1
	}

	for i := 0 ; i < len(alpha1) ; i++ {
		if alpha1[i] != alpha2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(anagram_checking_off("abcd", "dcba"))
	fmt.Println(anagram_checking_off("abcd", "abcc"))
	fmt.Println("---")
	fmt.Println(anagram_sort_and_compare("abcd", "dcba"), "( abcd, dcba )")
	fmt.Println(anagram_sort_and_compare("abcd", "abcc"), "( abcd, abcc )")
	fmt.Println("---")
	fmt.Println("")
	fmt.Println(anagram_count_compare("apple", "pleap"), "( apple, pleap )")
	fmt.Println(anagram_count_compare("apple", "applf"), "( apple, applf )")
}