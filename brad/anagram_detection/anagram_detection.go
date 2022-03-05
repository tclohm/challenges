package main

import (
	"strings"
	"fmt"
	"sort"
	"reflect"
)

// MARK: -- O(n^2)
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

// MARK: -- O(n log n)
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

// MARK: -- O(n)
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

// MARK: -- O(n)
func anagram_with_counter(s1, s2 string) bool {
	var dict1 = make(map[string]int)
	var dict2 = make(map[string]int)

	for i := 0 ; i < len(s1) ; i++ {
		char := string(s1[i])
		if _, ok := dict1[char]; ok {
			dict1[char]++
		} else {
			dict1[char] = 1
		}
		
	}

	for i := 0 ; i < len(s2) ; i++ {
		char := string(s2[i])
		if _, ok := dict2[char]; ok {
			dict2[char]++
		} else {
			dict2[char] = 1
		}
	}
	
	return reflect.DeepEqual(dict1, dict2)
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
	fmt.Println("---")
	fmt.Println(anagram_with_counter("apple", "pleap"), "( apple, pleap )")
	fmt.Println(anagram_with_counter("apple", "applf"), "( apple, applf )")
}