package main

import "fmt"

func minOp(logs []string) int {
	result := 0
	for _, log := range logs {
		if log == "./" {
			continue
		} else if log == "../" {
			result = max(0, result - 1)
		} else {
			result += 1
		}
	}

	return result
}

func main() {
	fmt.Println(minOp([]string{"d1/","d2/","../","d21/","./"}))
	fmt.Println(minOp([]string{"d1/","d2/","./","d3/","../","d31/"}))
	fmt.Println(minOp([]string{"d1/","../","../","../"}))

}
