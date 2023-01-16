package main

import (
	"fmt"
	"os"
	"bufio"
)

type stack struct {
	elements []rune
}

func (s *stack) push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *stack) pop()(r rune) {
	length := len(s.elements) - 1
	if length < 0 { return }
	r = s.elements[length]
	s.elements = s.elements[:length]
	return
}

func (s *stack) addToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}


func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	stacks := make([]stack, 9)

	sc.Scan()

	for (sc.Text() != " 1   2   3   4   5   6   7   8   9 ") {
		for i, r := range sc.Text() {
			if r != ' ' && r != '[' && r != ']' {
				// add to stack
				// fmt.Println(i, i/4, string(r))
				stacks[i / 4].addToBottom(r)
			}
		}
		sc.Scan()
	}

	sc.Scan()

	for sc.Scan() {
		var amount, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &amount, &from, &to)
		
		for move := 0 ; move < amount ; move++ {
			stacks[to - 1].push(stacks[from - 1].pop())
		}
	}

	for _, s := range stacks {
		fmt.Print(string(s.pop()))
	}
}