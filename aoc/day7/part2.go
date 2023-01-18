package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

// Find a directory you can delete that will free up enough space

// Find the smallest directory that, if deleted, 
// would free up enough space on the filesystem 
// to run the update. What is the total size of that directory?
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

	// TOTAL_DISK_SPACE := 70000000
	// SPACE_NEEDED := 30000000

	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	var curr *node

	var dirs = []*node{}

	for sc.Scan() {
		line := strings.Fields(sc.Text())

		if len(line) > 2 {
			if line[2] == ".." {
				curr = curr.parent
			} else if line[2] == "/" {
				curr = &node{
					name: "/",
					size: 0,
					isFile: false,
					children: make(map[string]*node),
					parent: nil,				
				}
				dirs = append(dirs, curr)
			} else {
				curr = curr.children[line[2]]
			}
		} else if line[0] == "dir" {
			curr.children[line[1]] = &node{
				name: line[1],
				size: 0,
				isFile: false,
				children: make(map[string]*node),
				parent: curr,
			}
			dirs = append(dirs, curr.children[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			curr.children[line[1]] = &node{
				name: line[1],
				size: size,
				isFile: true,
				children: nil,
				parent: curr,
			}
		}
	}

	var free = 30000000 - (70000000 - calculate(*dirs[0]))
	var second int = calculate(*dirs[0])

	for _, dir := range dirs {
		size := calculate(*dir)
		if size > free && size - free < second - free {
			second = size
		}
	}

	fmt.Println(second)
}