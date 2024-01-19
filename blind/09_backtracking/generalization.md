# Generalize backtracking

## Subsets
```go
func subset(nums []int) [][]int {
	var list [][]int
	sort.Ints(nums)
	backtrack(list, []int{}, nums, 0)
	return list
}

func backtrack(list [][]int, tmp, nums []int, start int) {
	list = append(list, append([]int{}, tmp...))
	for i := start ; i < len(nums) ; i++ {
		tmp = append(tmp, nums[i])
		backtrack(list, tmp, nums, i + 1)
		tmp = tmp[:len(tmp) - 1]
	}
}
```

## Subsets 2
```go
func subsetWithDuplicates(nums []int) [][]int {
	var list [][]int
	sort.Ints(nums)
	backtrack(list, []int{}, nums, 0)
	return list
}

func backtrack(list [][]int, tmp, nums []int, start int) {
	list = append(list, append([]int{}, tmp...))
	for i := start ; i < len(nums) ; i++ {
		if i > start && nums[i] == nums[i - 1] continue // skip duplicates
		tmp = append(tmp, nums[i])
		backtrack(list, tmp, nums, i + 1)
		tmp = tmp[len(tmp) - 1]
	}
}
```