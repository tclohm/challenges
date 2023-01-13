package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func duplicates(str1, str2, str3 string) string {

	for i := 0 ; i < len(str1) ; i++ {
		for j := 0 ; j < len(str2) ; j++ {
			for k := 0 ; k < len(str3) ; k++ {
				if str1[i] == str2[j] && str2[j] == str3[k] {
					return string(str1[i])
				}
			}
		}
	}
	
	return ""
}

func createDict(str string) (set map[string]bool) {
	set = make(map[string]bool)
	for _, char := range str {
		set[string(char)] = true
	}
	return
}

func buildMap() map[string]int {
	var abc = "abcdefghijklmnopqrstuvwxyz"
	var upABC = strings.ToUpper(abc)
	var concat = abc + upABC

	var result = map[string]int{}

	for i, r := range concat {
		letter := string(r)
		result[letter] = i+1
	}

	return result
}

func main() {

	dict := buildMap()

	file, err := os.Open("file2.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total = 0

	for scanner.Scan() {
		first := createDict(scanner.Text())
		two := createDict(scanner.Text())
		three := createDict(scanner.Text())

		for letter := range first {
			l := string(letter)
			if two[l] && three[l] {
				total += dict[l]
			}
			break
		}
	}

	fmt.Println(total)
}