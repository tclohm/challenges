package main

import "fmt"

// steps: 1 step or 2 steps

			// 				2
			// 	2 - 1 = 1 		2 - 2 = 0
			// 1 - 1 = 0   1 - 2 = -1


			// 							3
			// 			3 - 1 = 2 				3 - 2 = 1
			// 	2 - 1 = 1 		2 - 2 = 0    1 - 1 = 0    1 - 2 = -1
			// 1 - 1 = 0     1 - 2 = -1 

func climbStairs(n int) int {
	total := 0
	
	var dfs func(n int, total *int)

	dfs = func(n int, total *int) {
		if n == 0 {
			*total += 1
			return
		}

		if n - 1 >= 0 {
			//fmt.Println("go left")
			//fmt.Println(n, "- 1 =", n - 1, "\ttotal", *total)
			dfs(n-1, total)
		}

		if n - 2 >= 0 {
			//fmt.Println("go right")
			//fmt.Println(n, "- 2 =", n - 2, "\ttotal", *total)
			dfs(n-2, total)
		}
	}

	dfs(n, &total)

	return total
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}