package main

import "fmt"

func commonChars(words []string) []string {
	common := map[string]int{}

	for _, r := range words[0] {
		letter := string(r)
		common[letter]++
	}

	words = words[1:]
	for _, word := range words {
		tmp := map[string]int{}

		for _, r := range word {
			letter := string(r)
			if common[letter] > 0 {
				tmp[letter]++
				common[letter]--
			}
		}

		common = tmp
	}
	
	result := []string{}

	for k, v := range common {
		for i := 0 ; i < v ; i++ {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	//fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}
