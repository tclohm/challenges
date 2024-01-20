package main

import "fmt"

func play(game []string) [][]string {
	var perm_list [][]string
	var backtrack func(game []string, current_game []string)
	
	backtrack = func(game []string, current_game []string) {
		if len(current_game) == len(game) {
			perm_list = append(perm_list, append([]string{}, current_game...))
			return
		}	

		for i := 0 ; i < len(game) ; i++ {
			current_game = append(current_game, game[i])
			backtrack(game, current_game)
			current_game = current_game[:len(current_game) - 1]
		}

	}

	backtrack(game, []string{})
	return perm_list
}

func main() {
	fmt.Println(play([]string{"rock", "papper", "scissors"}))
}