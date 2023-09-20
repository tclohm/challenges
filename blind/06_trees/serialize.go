package main

import (
	tree "trees/common"
	"fmt"
	"strconv"
	"strings"
)

// Serializes a tree to a single string.
func serialize(root *tree.Node) string {
	var serialized = []string{}
    var dfs func(root *tree.Node)

    dfs = func(root *tree.Node) {
    	if root == nil { 
    		serialized = append(serialized, "null") 
    	} else {
    		serialized = append(serialized, strconv.Itoa(root.Value))
    		dfs(root.Left)
    		dfs(root.Right)
    	}
    }
    dfs(root)
    return strings.Join(serialized, ",")
}

// Deserializes your encoded data to tree.
func deserialize(data string) *tree.Node {    
    var build_tree func() *tree.Node
    var position = 0
    var values = strings.Split(data, ",")

    build_tree = func() *tree.Node {
    		if values[position] == "null" {
    			position += 1
    			return nil
    		}
    		val, _ := strconv.Atoi(values[position])
    		node := &tree.Node{Value: val}
    		position += 1
    		node.Left = build_tree()
    		node.Right = build_tree()
    		return node
    }
    return build_tree()
}

func main() {
	r := tree.Build_binary_tree([]int{1,2,3,4,5,6,7,8,9,10,11,12})
	tree.Print_binary_tree(r, 0, "Root: ")
	s := serialize(r)
	tree.Line_break()
	fmt.Println("serialized", s)
	d := deserialize(s)
	tree.Print_binary_tree(d, 0, "Root: ")
	tree.Line_break()
}