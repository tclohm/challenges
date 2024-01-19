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

## Combination Sum
```go
func combinationSum(nums []int, target int) [][]int {
	var list [][]int
	backtrack(list, []int{}, nums, target, 0)
	return list
}

func backtrack(list [][]int, tmp, nums []int, remain, start int) {
	if remain < 0 { return }
	else if remain == 0 {
		list = append(list, append([]int{}, tmp...))
	} else {
		for i := start ; i < len(nums) ; i++ {
			tmp = append(tmp, nums[i])
			backtrack(list, tmp, nums, remain - nums[i], i + 1)
			tmp = tmp[:len(tmp) - 1]
		}
	}
}
```

## Combination Sum 2
```go
func combinationSum2(nums []int, target int) [][]int {
	var list [][]int
	backtrack(list, []int{}, nums, target, 0)
	return list
}

func backtrack(list [][]int, tmp, nums []int, remain, start int) {
	if remain < 0 { return }
	else if remain == 0 {
		list = append(list, append([]int{}, tmp...))
	} else {
		for i := start ; i < len(nums) ; i++ {
			if i > start && nums[i] == nums[i - 1] { continue }
			tmp = append(tmp, nums[i])
			backtrack(list, tmp, nums, remain - nums[i], i + 1)
			tmp = tmp[:len(tmp) - 1]
		}
	}
}
```

## Palindrome Partition
```go
func partition(s string) [][]string {
	var list [][]string
	backtrack(list, []string{}, s, 0)
	return list
}

func backtrack(list [][]string, tmp []string, s string, start int) {
	if start == len(s) {
		list = append(list, append([]string{}, tmp...))
	} else {
		for i := start ; i < len(s) ; i++ {
			if isPalindrome(s, start, i) {
				tmp = append(tmp, s[start:i + 1])
				backtrack(list, tmp, s, i + 1)
				tmp = tmp[:len(tmp) - 1]
			}
		}
	}
}

func isPalindrome(s string, low int, high int) {
	for low < high {
		if s[low] != s[high] { return false }
		low++
		high--
	}
	return true
}
```