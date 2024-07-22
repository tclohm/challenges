package main

import (
	"fmt"
	t "tree"
)

func getDirections(root *t.Node, start, end int) string {
	startArray := dfs(root, []string{}, start)
	endArray := dfs(root, []string{}, end)
	i := 0


	var dfs func(root *t.Node, path []string, target int) []string
	dfs = func(root *t.Node, path []string, target int) []string {
		if root.Val == target {
			return path
		}
		if root.Left != nil {
			newPath := append(path, "L")
			if found := dfs(root.Left, newPath, target); found != nil {
				return found
			}
		}
		if root.Right != nil {
			newPath := append(path, "R")
			if found := dfs(root.Right, newPath, target); found != nil {
			return found
			}
		}
		return nil
	}

	for len(startArray) > i && len(endArray) i > && startArray[i] == endArray[i] {
		i++
	}
	pathUp := strings.Repeat("U", len(startArray) - 1)
	pathDown := string(endArray[i:])
	return pathUp + pathDown
}
