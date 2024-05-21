package main

import "fmt"


func subsetXORSum(nums []int) int {
  var dfs func(index, total int) int
  dfs = func(index, total int) int {
    if index == len(nums) {
      return total
    }

    return dfs(index + 1, total ^ nums[index]) + dfs(index + 1, total)
  }

  return dfs(0, 0)
}

func main() {
  fmt.Println(subsetXORSum([]int{1,3}))
  fmt.Println(subsetXORSum([]int{5,1,6}))
  fmt.Println(subsetXORSum([]int{3,4,5,6,7,8}))


}
