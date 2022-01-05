package main

import (
	"fmt"
)

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) Add(element int) {
	if !set.Contains(element) {
		set.integerMap[element] = true
	}
}

func (set *Set) Delete(element int) {
	delete(set.integerMap, element)
}

func (set *Set) Contains(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

func (set *Set) Intersect(other *Set) *Set {
	var intersect = &Set{}
	intersect.New()
	
	for value, _ := range set.integerMap {
		if other.Contains(value) {
			intersect.Add(value)
		} 
	}
	return intersect
}

func (set *Set) Union(other *Set) *Set {
	var union = &Set{}
	union.New()

	for value, _ := range other.integerMap {
		union.Add(value)
	}

	for value, _ := range set.integerMap {
		union.Add(value)
	}

	return union
}

func main() {
	var set *Set
	set = &Set{}

	set.New()

	set.Add(42)
	set.Add(21)
	fmt.Println(set)

	fmt.Println(set.Contains(21))
	fmt.Println(set.Contains(1))

	var other *Set
	other = &Set{}

	other.New()
	
	other.Add(2)
	other.Add(21)
	other.Add(5)

	fmt.Println(set.Intersect(other))

	fmt.Println(set.Union(other))

}