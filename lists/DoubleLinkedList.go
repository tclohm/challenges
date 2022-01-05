package main

import (
	"fmt"
	_"log"
	"strconv"
)

type Node struct {
	value int
	next *Node
	prev *Node
}

