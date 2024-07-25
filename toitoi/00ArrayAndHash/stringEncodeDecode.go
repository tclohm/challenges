package main

import "fmt"

// join
func encode(arr []string) string {
	encoded := ""
	for _, word := range arr {
		encoded += fmt.Sprintf("%s;", word)
	}
	return encoded
}

// split
func decode(s string) []string {
	decoded := []string{}
	i := 0
	j := 0
	for i < len(s) - 1 {

		for s[j] != ';' {
			j++
		}

		decoded = append(decoded, s[i:j])

		i = j + 1
		j = i + 1
	}
	return decoded
}

func main() {
	input := []string{"I", "Love", "You"}
	input2 := []string{"we", "said", ":", "yes"}
	input3 := []string{"Someday", "I", "CAN'T", "wait", "to", "be", "KING!"}
	fmt.Println(encode(input))
	fmt.Println(encode(input2))
	fmt.Println(encode(input3))
	fmt.Println(decode(encode(input)))
	fmt.Println(decode(encode(input2)))
	fmt.Println(decode(encode(input3)))
}

