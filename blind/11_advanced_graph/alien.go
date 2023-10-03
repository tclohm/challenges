package main

import "fmt"

// return a string of the unique letters in the new alien language sorted in lexicographically
// increasing order
// smaller word comes first: ex, ape, app, apple,
func alien(words []string) string {
	adjancency := make(map[string]struct{})

	for _, word := range words {
		for _, r := range word {
			s := string(r)
			adjancency[s] = struct{}{}
		}
	}

	for i := 0 ; i < len(words) - 1 ; i++ {
		w1, w2 := words[i], word[i + 1]
		smaller := min(len(w1), len(w2))
		if w1[:smaller] == w2[:smaller] && len(w1) > len(w2) {
			return ""
		}

		for index := 0 ; index < smaller ; index++ {
			if w1[index] != w2[index] {
				adjancency[string(w1[index])]
				break
			}
		}
	}

	fmt.Println(adjancency)

	return "ok"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// sorted order
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	alien(words)
}