package main

import "fmt"
// A subset of nums is beautiful if it does not contain two integers with an absolute difference equal to k.
func beautifulSubsets(nums []int, k int) int {
  // include or not include O(2^n)  return 0

  var dfs func (index int, count map[int]int) int
  dfs = func (index int, count map[int]int) int {
    if index == len(nums) { return 1 }
    // skip
    res := dfs(index + 1, count)
    // include
    var add = nums[index] + k 
    var minus = nums[index] - k
    if _, ok := count[add]; !ok {
      if _, ok := count[minus]; !ok {
        count[nums[index]] += 1
        res += dfs(index + 1, count)
        // clean up
        count[nums[index]] -= 1
      }
    }
    return res
  }
  return dfs(0, map[int]int{}) - 1
}

func main() {
  fmt.Println(beautifulSubsets([]int{2,4,6}, 2))
  fmt.Println(beautifulSubsets([]int{1}, 1))
}
