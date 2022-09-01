package main

import "fmt"

func rps(players int) [][]string {
	var moves = []string{"rock", "paper", "scissors"}
	var games [][]string
	var results []string
	// MARK : -- if we remove our index, which is 0
	// 			 we get the permutations
	// 			 index allows us to constrain the permutations
	dfs(players, &results, &games, moves, 0)
	return games
}

func dfs(players int, results *[]string, games *[][]string, moves []string, index int) {
	if players < 1 {
		*games = append(*games, *results)
		return
	}

	for i := index ; i < len(moves) ; i++ {
		// MARK: -- new array for each decision
		tmp := []string{}
		tmp = append(tmp, *results...)
		tmp = append(tmp, moves[i])
		dfs(players - 1, &tmp, games, moves, i)
	}

	return
}

func main() {
	fmt.Println(rps(2))
}