package main

import (
	"fmt"
	"os"
	"bufio"
)

// A Rock, X
// B Paper, Y
// C Scissors, Z


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
		"C Y": 2, 
		"A Z": 3, 
		"A X": 4,
		"B Y": 5,
		"C Z": 6,
		"C X": 7,
		"A Y": 8, 
		"B Z": 9,
	}

	for sc.Scan() {
		score += scores[sc.Text()]
	}

	fmt.Println(score)
}