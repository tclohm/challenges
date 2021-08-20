import "math/rand"

func quicksort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
            more = append(more, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			less = append(less, item)
		}
	}

	less, more = quicksort(less), quicksort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func kthLargestElement(nums []int, k int) int {
   sorted := quicksort(nums)
   return sorted[k - 1]
}
