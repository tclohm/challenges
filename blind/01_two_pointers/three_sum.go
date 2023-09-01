package main

import (
    "fmt"
    "sort"
)

func threeSum(nums []int) [][]int {
    var res [][]int

    if len(nums) < 3 { return res }
    // sort the numbers
    sort.Ints(nums)
    // make sure we dont go over
    for i := 0 ; i < len(nums) - 2 ; i++ {

        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        left, right := i + 1, len(nums) - 1

        for left < right {
            target := nums[i] + nums[left] + nums[right]

            if target == 0 {
                res = append(res, []int{nums[i], nums[left], nums[right]})

                left, right = left+1, right-1
                // if left less than right and left equal to what's left of it
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                // if left less than right and right is equal to what's right of it
                for left < right && nums[right] == nums[right+1] {
                    right++
                }

            } else if target > 0 { 
                right-- 
            } else { 
                left++ 
            }
        }
    }
    return res
}

func main() {
    fmt.Println(threeSum([]int{-1,0,1,2,-1,-4})) // [[-1, -1, 2], [-1, 0, 1]] 
    fmt.Println(threeSum([]int{0,1,1})) // []
    fmt.Println(threeSum([]int{0,0,0})) // [0, 0, 0]
}