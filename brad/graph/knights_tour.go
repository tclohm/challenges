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

	for row := 0 ; row < board_size ; row++ {
		for col := 0 ; col < board_size ; col++ {
			to_row, to_col := legal_moves_from(row, col, board_size)

			node1 := Set{
				row: row,
				col: col,
			}

			node2 := Set{
				row: to_row,
				col: to_col,
			}
			add_edge(graph, node1, node2)
		}
	}
	return graph
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
	return 0,0
}

// MARK: -- DFS, search tree by one branch of the tree as deeply as possible
func first_true(sequence []Set) (int, Set) {
	for _, item := range sequence {
		if item != nil {
			return _, item
		}
	}
	return 0, _
}

func find_solution_for(board_size int, heuristic func(g Graph) []Set) {
	g := build_graph(board_size)
	total_squares := board_size * board_size

	func traverse(path []Set, current_node Set) {
		if len(path) + 1 == total_squares {
			new_path := path[:]
			new_path = append(new_path, current_node)
			return new_path
		}

		have_not_been_visited := notVisited(graph[current_node], path)

		if len(have_not_been_visited) == 0 {
			// no unvisited neighbor, dead end
			return false
		}

		next_nodes = sorted(have_not_been_visited, heuristic(graph))
		new_path := path[:]
		new_path = append(new_path, current_node)

		// MARK: -- ok, i'm moving forward 
		return first_true(traverse(new_path))
	}
}

func main() {
}