package main

import "fmt"

type Set struct {
	container map[int]bool
}

func NewSet() *Set {
	s := &Set{
		container: make(map[int]bool),
	}
	return s
}

func (s *Set) Exists(key int) bool {
	_, exists := s.container[key]
	return exists
}

func Max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func (s *Set) Add(key int) {
	s.container[key] = true
}

func lcs(array []int) int {
	var set *Set = NewSet()
	var longest int = 0

	for _, v := range array {
		set.Add(v)
	}
	
	for _, n := range array {
		exists := set.Exists(n - 1); if !exists {
			var length int = 0
			for {
				exists := set.Exists(n + length)
				if exists {
					length += 1
					continue
				} else {
					longest = Max(length, longest)
					break
				}
			}
		}
	}

	return longest
}

func main() {
	var arr = []int{100, 4, 200, 1, 3, 2}
	fmt.Println(lcs(arr))
}