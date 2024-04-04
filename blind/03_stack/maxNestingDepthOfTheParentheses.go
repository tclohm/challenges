package main

import "fmt"

func maxDepth(s string) int {
    return 1
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
}