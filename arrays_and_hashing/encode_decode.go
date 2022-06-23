package main

import (
	"fmt"
	"strconv"
)

func encode(ls []string) string {
	var result string = ""
	for _, word := range ls {
		var length string = strconv.Itoa(len(word))
		result += length + "#" + word
	}
	return result
}

func decode(s string) []string {
	var result = []string{}
	var i int = 0
	for i < len(s) {
		var j int = i

		for string(s[j]) != "#" {
			j += 1
		}

		length, err := strconv.Atoi(string(s[i:j]))
		if err != nil {
			fmt.Println("error")
			return result
		}
		result = append(result, s[j + 1 : j + 1 + length])
		i = j + 1 + length
	}
	return result
}

func main() {
	i := []string{"lint","code","love","you"}
	i2 := []string{"we", "say", ":", "yes"}

	encodedI := encode(i)
	encodedII := encode(i2)

	fmt.Println(encodedI)
	fmt.Println(encodedII)

	fmt.Println(decode(encodedI))
	fmt.Println(decode(encodedII))

}