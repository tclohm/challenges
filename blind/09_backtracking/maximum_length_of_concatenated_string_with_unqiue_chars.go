package main

import "fmt"

func maxLength(chars []string) int {
	var backtrack func (index int)
	backtrack = func (index int) {
		if index == len(chars) {
			return len(set)
		}
		var result = 0
		if overlap(set, chars[index]) {
			for _, letter := range chars[index] {
				set[chars[index]] = 
			}
		}
	}
}


func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
}