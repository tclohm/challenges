package main

import "fmt"

type Queue struct {
	container []int
}

func (q *Queue) push(n int) {
	q.container = append(q.container, n)
}

func (q *Queue) pop() {
	q.container = q.container[1:]
}

func (q *Queue) peek() int {
	return q.container[0]
}

func (q *Queue) peekLast() int {
	return q.container[len(q.container) - 1]
}

func (q *Queue) length() int {
	return len(q.container)
}

func (q *Queue) max() int {
	max := -190
	for i := 0 ; i < len(q.container) ; i++ {
		if max < q.container[i] {
			max = q.container[i]
		}
	}
	return max
}

func maxSlidingWindow(nums []int, k int) []int {
    result := []int{}

    for i := 0 ; i + k <= len(nums) ; i++ {
    	max := max(nums[i:i+k])
    	result = append(result, max)
    }
    return result
}

func max(n []int) int {
	if len(n) == 1 { return n[0] }
	big := n[0]
	for i := 0 ; i < len(n) ; i++ {
		if big < n[i] {
			big = n[i]
		}
	}
	return big
}

func maxSlidingWithQueue(nums []int, k int) []int {
	results := []int{}
	left, right := 0, 0
	queue := Queue{}

	for right < len(nums) {
		for queue.length() != 0 && nums[queue.peekLast()] < nums[right] {
			queue.pop()
		}

		queue.push(right)

		if left > queue.peek() {
			queue.pop()
		}

		if (right + 1) >= k {
			results = append(results, nums[queue.peek()])
			left++
		}
		right++
	}

	return results
}


func main() {
	fmt.Println(maxSlidingWithQueue([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWithQueue([]int{1}, 1))
}