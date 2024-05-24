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

func optimalSolution(nums []int, k int) int {
  counter := make(map[int]int)

  for _, n := range nums {
    counter[n] += 1
  }

  groups := []map[int]int{}
  visited := make(map[int]bool)
  

  var helper func(n int, g map[int]int) int
  helper = func(n int, g map[int]int) int {
    if _, in := g[n]; !in { return 1 }
    skip := helper(n + k, g)
    include := helper(n + 2*k, g)
    return skip + include
  }

  for key, _ := range counter {
    if _, in := visited[key]; in { continue }
    g := make(map[int]int)

    for counter[key - k] > 0 {
      key -= key
    }

    for counter[key] > 0 {
      g[key] = counter[key]
      visited[key] = true
    }

    groups = append(groups, g)
  }

  result := 1
  // variation of house robber algo
  for g := range groups {
    n := 10000000
    for key := range g {
      if key < n {
        n = key
      }
    }
    result *= helper(n, g)
  }
  return result - 1
}

func main() {
  fmt.Println(beautifulSubsets([]int{2,4,6}, 2))
  fmt.Println(beautifulSubsets([]int{1}, 1))
}
