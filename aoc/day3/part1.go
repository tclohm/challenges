package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// rucksack two large compartments
// elf packing failed to follow rule
// for exactly one item type per rucksack

// find the errors
// capitalized and uncap are the types

// vJrwpWtwJgWrhcsFMMfFFhFp
// first compartment has p, second compartment has p too

// lowercase item a-z, priority 1-26
// uppercase item A-Z, priority 27-52

// QUESTION:
// find the item type that appears in both compartments of each rucksack
// what is the sum of the priorities of those item types?

func duplicate(str1, str2 string) string {
	ht := map[string]bool{}
	for i := 0 ; i < len(str1) ; i++ {
		word := string(str1[i])
		ht[word] = true
	}

	for i := 0 ; i < len(str2) ; i++ {
		word := string(str2[i])
		if _, ok := ht[word]; ok {
			return word
		}
	}

	return "nope"
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

	sum := 0

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		length := len(scanner.Text())
		compartmentOne, compartmentTwo := scanner.Text()[0: length / 2], scanner.Text()[length / 2:length]
		
		dup := duplicate(compartmentOne, compartmentTwo)

		sum += dict[dup]
	}

	fmt.Println(sum)
}