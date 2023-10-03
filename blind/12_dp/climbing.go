package main

import "fmt"

func climbStairs(n int) int {
    var climb func(int) int

    climb = func(n int) int {
    	if n == 1 {
    		return 1
    	} else if n == 2 {
    		return 2
    	}

    	return climb(n - 1) + climb(n - 2)
    }

    return climb(n)
}

var mem map[int]int = map[int]int{}
func dpclimb(n int) int {
	if n == 1 { 
		return 1 
	} else if n == 2 {
		return 2
	} else if val, ok := mem[n]; ok {
		return val
	}

	res := dpclimb(n - 1) + dpclimb(n - 2)
	mem[n] = res
	return res
}

func dpiterative(n int) int {
	res := 0
	next := 0
	prev := 0

	for i := 1 ; i <= n ; i++ {
		if i == 1 {
			res = 1
		} else if i == 2 {
			res = 2
		} else {
			res = next + prev
		}

		prev = next
		next = res
	}

	return res
}

func optimized(n int) int {
	tail, head := 0, 1
	for i := 1 ; i <= n ; i++ {
		tail, head = head, tail + head
	}

	return head
}

func main() {
	fmt.Println(optimized(2))
	fmt.Println(optimized(3))
}