package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	//"reflect"
)

type Graph map[string][]string

func pop(q *[][]string) ([]string, [][]string) {
	// MARK: -- Parens around pointer allows for indexing
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped, *q
}

func notIn(s string, arr []string) bool {
	for _, str := range arr {
		if s == str {
			return false
		}
	}
	return true
}

func build_neighbors(node []string, visited []string) []string {
	unvisited := []string{}
	for _, neighbor := range node {
		var in bool = false
		for _, visited_neighbor := range visited {
			if neighbor == visited_neighbor {
				in = true
			}
		}
		if in == false {
			unvisited = append(unvisited, neighbor)
			fmt.Println(unvisited)
		}
		in = false
	}
	return unvisited
}

func build_graph_for_words() Graph {
	buckets := map[string][]string{}
	graph := map[string][]string{}

	file, err := os.Open("vocabulary.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		word := scanner.Text()
		for index, _ := range word {
			// MARK: -- value representation %v
			bucket := fmt.Sprintf("%v_%v", word[:index], word[index + 1:])

			if _, ok := buckets[bucket]; !ok {
				buckets[bucket] = []string{}
			}

			buckets[bucket] = append(buckets[bucket], word)
		}
	}

	for _, mutual_neighbors := range buckets {

		for _, word1 := range mutual_neighbors {

			for _, word2 := range mutual_neighbors {

				if word1 != word2 {

					if _, ok := graph[word1]; !ok {
						graph[word1] = []string{}
					}

					if _, ok := graph[word2]; !ok {
						graph[word2] = []string{}
					}

					if not := notIn(word1, graph[word2]); not {
						graph[word2] = append(graph[word2], word1)
					}

					if not := notIn(word2, graph[word1]); not {
						graph[word1] = append(graph[word1], word2)
					}
				}
			}
		}
	}

	file.Close()

	return graph
	
}

func traverse(graph Graph, starting_place, end_place string) []string {
	visited := []string{}

	var queue = [][]string{{starting_place}}

	for len(queue) > 0 {

		path := queue[0]
		queue = queue[1:]
		
		end := len(path) - 1
		node := path[end]

		if node == end_place {
			return path
		}

		// check neighbors in graph except for the visited
		unvisited_neighbors := build_neighbors(graph[node], visited)
		
		for _, neighbor := range unvisited_neighbors {
			if neighbor == end_place {
				path = append(path, neighbor)
				return path
			}
			// add neighbors into visited
			visited = append(visited, neighbor)
			// append the path + neighbor to queue
			path = append(path, neighbor)
			queue = append(queue, path)
		}
	}
	n := []string{"shit"}
	return n
}

func main() {
	g := build_graph_for_words()
	fmt.Println(traverse(g, "FOOL", "SAGE"))

}