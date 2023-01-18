package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

// $ are the commands we execute
// cd - change directory
// cd x - move in one level
// cd .. - move out one level
// cd / switches current dir to the outermost dir
// ls - list
// 123 abc - 123 size, abc file
// dir xyz - dir name xyz

// find directories that are good candidates for deletion
// determine total size of each directory
// total size of a dir is the sum of the sizes of the files it contains, directly or indirectly

// find all of the directories with a total size of at most 100,000
// calculate the sum of their total sizes

type node struct {
	name string
	size int
	isFile bool
	children map[string]*node
	parent *node
}

func calculate(root node) (size int) {
	if root.isFile {
		return root.size
	}

	for _, d := range root.children {
		size += calculate(*d)
	}

	return
}


func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	var currentDir *node

	dirs := []*node{}

	for sc.Scan() {
		line := strings.Fields(sc.Text())
		if len(line) > 2 {
			if line[2] == ".." {
				currentDir = currentDir.parent
			} else if line[2] == "/" {
				currentDir = &node{"/", 0, false, make(map[string]*node), nil}
			} else {
				currentDir = currentDir.children[line[2]]
			}
		} else if line[0] == "dir" {
			currentDir.children[line[1]] = &node{line[1], 0, false, make(map[string]*node), currentDir}
			dirs = append(dirs, currentDir.children[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			currentDir.children[line[1]] = &node{line[1], size, true, nil, currentDir}
		}
	}

	var total int

	for _, dir := range dirs {
		size := calculate(*dir)
		if size <= 100000 {
			total += size
		}
	}

	fmt.Println(total)

}