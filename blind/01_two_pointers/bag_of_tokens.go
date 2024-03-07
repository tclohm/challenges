package main

import "fmt"
// return the MAX possible socre you can achieve after playing any number of tokens

// power
// score = 0

// tokens

// goal is to max the total score by playing these tokens

// two decisions on a token
//		face up   -- current power is at least token[i], score++
//		face down -- current score is at least 1, play token, gaining the token power, and score--

func bagOfTokensScore(tokens []int, power int) int {
	score := 0

	for i := 0 ; i < len(tokens) ; i++ {
		if power >= tokens[i] {
			power -= tokens[i]
			score++
		} else if score >= 1 && power <= tokens[i] {
			score--
			power += tokens[i]
		}
	}

	return score
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))
	fmt.Println(bagOfTokensScore([]int{200, 100}, 150))
	fmt.Println(bagOfTokensScore([]int{100,200,300,400}, 200))
}