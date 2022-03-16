package main

import "fmt"

type Set struct {
	row int
	col int
}

type Graph map[Set][]Set


func add_edge(g Graph, node1, node2 Set) {
	if _, ok := g[node1]; !ok {
		g[node1] = []Set{}
	}

	if _, ok := g[node2]; !ok {
		g[node2] = []Set{}
	}
	g[node1] = append(g[node1], node2)
	g[node2] = append(g[node2], node1)
}

func build_graph(board_size int) Graph {
	graph := Graph{}

	for row := 0 ; row < len(board_size) ; row++ {
		for col := 0 ; col < len(board_size) ; col++ {

		}
	}
}

func legal_moves_from(row int, col int, board_size int) (int, int)  {
	MOVE_OFFSETS := [][]int{
				{-1, -2}, {1, -2},
		{-2, -1},				{2, -1},
		{-2, 1},				{2, 1},
				{-1, 2}, {1, 2},
	}

	for _, offset := range MOVE_OFFSETS {
		row_offset, col_offset := offset[0], offset[1]
		move_row, move_col := row + row_offset, col + col_offset
		
		if 0 <= move_col && move_col < board_size  && 0 <= move_row && move_col < board_size {
			return move_row, move_col
		}
	}

	retu

func main() {
	g := Graph{}
	fmt.Println(g)
	n1 := Set{1, 2}
	n2 := Set{3, 4}
	add_edge(g, n1, n2)
	fmt.Println(g)
}