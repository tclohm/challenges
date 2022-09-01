package main

import "fmt"

func rps(players int) [][]string {
	var moves = []string{"rock", "paper", "scissors"}
	var games [][]string
	var results []string
	return dfs(players, &results, &games, moves)
}

func dfs(players int, results *[]string, games *[][]string, moves []string) [][]string {
	if players < 1 {
		*games = append(*games, *results)
		return *games
	}

	for _, move := range moves {
		tmp := []string{}
		tmp = append(tmp, *results...)
		tmp = append(tmp, move)
		dfs(players - 1, &tmp, games, moves)
	}

	return *games
}

func main() {
	fmt.Println(rps(1))
}