package main

import (
	"fmt"
	"sort"
)
// return the MAX possible socre you can achieve after playing any number of tokens

// power
// score = 0

// tokens

// goal is to max the total score by playing these tokens

// two decisions on a token
//		face up   -- current power is at least token[i], score++
//		face down -- current score is at least 1, play token, gaining the token power, and score--

func bagOfTokensScore(tokens []int, power int) int {

	sort.Ints(tokens)

	result, score, left, right := 0, 0, 0, len(tokens) - 1

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			result = max(result, score)
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else { break }
	}

	return result
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))
	fmt.Println(bagOfTokensScore([]int{200, 100}, 150))
	fmt.Println(bagOfTokensScore([]int{100,200,300,400}, 200))
}