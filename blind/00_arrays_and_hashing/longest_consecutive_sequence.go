package main

import "fmt"

func mySolution(nums []int) int {
    largest := 0

    for _, n := range nums {
    	if n > largest {
    		largest = n
    	}
    }

    consecutive := make([]int, largest+1)

    for _, n := range nums {
    	consecutive[n] = 1
    }

    longest := 0
    for i := 0 ; i < len(consecutive) ; i++ {
    	if consecutive[i] == 1 {
    		tmp := 0
    		for consecutive[i] != 0 {
               
                tmp++
    			if tmp > longest {
    				longest = tmp
    			}
                if i == len(consecutive) - 1 { break }
    			i++
    		}
    	}
    }

    return longest
}

func longestConsecutive(nums []int) int {
    set := map[int]bool{}

    for _, num := range nums {
        set[num] = true
    }

    res := 0

    for _, num := range nums {
        if set[num-1] { continue }
        sequence := 1
        tmp := num + 1

        for set[tmp] {
            sequence++
            tmp++
        }

        if sequence > res {
            res = sequence
        }
    }

    return res
}

func main() {
	fmt.Println(longestConsecutive([]int{100,4,200,1,3,2}))
	fmt.Println(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}