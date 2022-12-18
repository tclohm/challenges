package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	container []interface{}
}

func (s *Stack) print(msg string) {
	fmt.Println(msg, "\nstack:", s.container)
}

func (s *Stack) Pop() interface{} {
	end := len(s.container) - 1
	item := s.container[end]
	s.container = s.container[0:end]
	s.print("\tpop")
	return item
}

func (s *Stack) Push(item interface{}) {
	s.container = append(s.container, item)
	s.print("\tpush")
} 

func (s *Stack) isEmpty() bool {
	return len(s.container) == 0
}

func (s *Stack) length() int {
	return len(s.container)
}

func isValid(s string) bool {
	pairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	st := Stack{}

	for _, r := range s {

		if opening, ok := pairs[r]; ok {
			fmt.Println(string(opening), "==", string(r))
			if opening != st.Pop() { return false }
		} else {
			st.Push(r)
		}
		
	}

	return st.isEmpty()
}

func mRemoveToMakeValid(str string) string {
	var result = strings.Split(str, "")
	var stack = Stack{}

	for i := 0 ; i < len(result) ; i++ {
		if result[i] == "(" {
			stack.Push(i)
		} else if result[i] == ")" && !stack.isEmpty() {
			stack.Pop()
		} else if result[i] == ")" {
			result[i] = ""
		}
	}

	for !stack.isEmpty() {
		var index = stack.Pop()
		result[index.(int)] = ""
	}

	return fmt.Sprint("result: ", strings.Join(result, ""))
}

func main() {
	not_valid := "{([])]"
	valid := "{([])}"
	fmt.Println("\tNOT VALID")
	fmt.Println(isValid(not_valid))
	fmt.Println("\n\tVALID")
	fmt.Println(isValid(valid))

	// remove the least amount of brackets so the string is valid
	// string is considered valid if it is empty or if the brackets are closed
	s1 := "a)bc(d)" // abc(d)
	s2 := "(ab(c)d" // ab(c)d
					// (abc)d
	s3 := "))(("

	fmt.Println("--- a)bc(d) ---")
	fmt.Println(mRemoveToMakeValid(s1))
	fmt.Println("--- (ab(c)d ---")
	fmt.Println(mRemoveToMakeValid(s2))
	fmt.Println("--- ))(( ---")
	fmt.Println(mRemoveToMakeValid(s3))
}