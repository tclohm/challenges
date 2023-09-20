package main

import (
	tree "trees/common"
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.

// empty string and store the tree
// perform a depth-first traversal (pre-order traversal) of the binary tree
// each node append it's value to the string, followed by a delimiter
// if a node is empty, append with special marker, "null"
// continue process for all nodes in the tree
// result string will represent serialized binary tree
func (self *Codec) serialize(root *tree.Node) string {
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
// split the string using the delimiter to obtain a list of node values
// init an index to 0, keep track of the current position in the list of values and 
// index an arguments and return next node in tree
// recurse, if the current value at the index is the special marker, "null", increment the index and return nil
// otherwise create a new node with the current values and increment the index
// recurse on left and right child and return node
func (self *Codec) deserialize(data string) *tree.Node {    
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
	c := Constructor()
	r := tree.Build_binary_tree([]int{1,2,3,4,5,6,7,8,9,10,11,12})
	tree.Print_binary_tree(r, 0, "Root: ")
	s := c.serialize(r)
	tree.Line_break()
	fmt.Println("serialized", s)
	d := c.deserialize(s)
	tree.Print_binary_tree(d, 0, "Root: ")
	tree.Line_break()
}