package main

import (
	"fmt"
	"os"
	"bufio"
)

// A Rock, X
// B Paper, Y
// C Scissors, Z

// X --- lose
// Y --- draw
// Z --- win


// highest score
// score, sum(scores for each round)

// score
// 1 Rock
// 2 Paper
// 3 Scissors
// 0 if lost, 3 if draw, 6 if won

// calculate the score you shoulg get if you were to follow the strategy

// what would be your score if you follow the guide

func main() {


	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	score := 0

	scores := map[string]int{
		"B X": 1, 
		"C X": 2,
		"A X": 3, 
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"C Z": 7,
		"A Z": 8,
		"B Z": 9,
	}

	for sc.Scan() {
		score += scores[sc.Text()]
	}

	fmt.Println(score)
}