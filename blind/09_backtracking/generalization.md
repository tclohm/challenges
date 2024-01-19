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

## Permutations
```go
func permute(nums []int) {
	var list [][]int
	backtrack(list, []int{}, nums)
	return list
}

func contains(element int, list []int) bool {
	for _, el := range list {
		if el == element { return true }
	}
	return false
}

func backtrack(result [][]int, tmp, nums []int) {
	if len(tmp) == len(nums) {
		list = append(list, append([]int{}, tmp...))
	} else {
		for i := 0 ; i < len(nums) ; i++ {
			if contains(nums[i], tmp) continue
			tmp = append(tmp, nums[i])
			backtrack(result, tmp, nums)
			tmp = tmp[:len(tmp) - 1]
		}
	}
}
```

## Permutations 2
```go
func permute(nums []int) {
	var list [][]int
	var used make([]bool, len(nums), len(nums))
	backtrack(list, []int{}, nums, used)
	return list
}

func backtrack(result [][]int, tmp, nums []int, used []bool) {
	if len(tmp) == len(nums) {
		result = append(result, append([]int{}, tmp...))
	} else {
		for i := 0 ; i < len(nums) ; i++ {
			if used[i] || i > 0 && nums[i] == nums[i - 1] && !used[i - 1] continue
			used[i] = true
			tmp = append(tmp, nums[i])
			backtrack(list, tmp, nums, used)
			tmp = tmp[:len(tmp) - 1]
		}
	}
}
```