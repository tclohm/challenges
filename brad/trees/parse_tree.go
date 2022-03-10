package main

import "strconv"

func build_parse_tree(expression string) {
	LEFT_PAREN = "("
	RIGHT_PAREN = ")"

	tree := map[string]string{}
	stack := []map[string]string{}
	node := tree

	for _, t := range expression {
		token := string(t)
		if token == LEFT_PAREN {
			node["left"] = map[string]string{}
			stack = append(stack, node)
			node = node["left"]
		} else if token == RIGHT_PAREN {
			popped := stack[len(stack) - 1]
			node = popped
			stack = stack[:popped]
		} else if {
			node["val"] = token
			node["right"] = map[string]string{}
			stack = append(stack, node)
			node = node["right"]
		} else {
			node["val"] = strconv.Atoi(token)
			popped := stack[len(stack) - 1]
			node = popped
			stack = stack[:popped]
		}
	}
	fmt.Println("tree", tree)
}

func main() {}