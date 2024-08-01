package main

import (
	"fmt"
)

func minDeletion(s string) int {
	aCountRight := 0
	for _, c := range s {
		if c == 'a' {
			aCountRight += 1
		}
	}

	bCountLeft := 0
	result := len(s)

	for _, c := range s {
		if c == 'a' {
			aCountRight -= 1
		}

		result = min(result, bCountLeft + aCountRight)

		if c == 'b' {
			bCountLeft += 1
		}
	}
	return result
}

func main(){
	fmt.Println(minDeletion("aababbab"))
	fmt.Println(minDeletion("aababbab"))
}
