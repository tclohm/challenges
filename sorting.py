# bubble sort, O(n^2)
# array is completely sorted when there are no inversions
def bubble_sort(array: [int]) -> [int]:
	for index in range(len(array)):
		for j_index in range(len(array) - 1):
			if array[j_index] > array[j_index + 1]:
				array[j_index], array[j_index + 1] = array[j_index + 1], array[j_index]
	return array

# merge sort, O(n log n)
# if a = b, do nothing, sub array is already sorted
# calculate the positiion of the middle element, k = | (a + b) / 2 |
# recursively sort the subarray array[a...k]
# recursively sort the subarray array[k + 1...b]
# merge the sorted subarrays array[a...k] and array[k + 1...b] into a sorted subarray array[a...b]

# efficient because halves the size of the subarray at each step


if __name__ == '__main__':
	print(bubble_sort([1,3,8,2,9,2,5,6]))