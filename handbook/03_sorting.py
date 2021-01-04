"""

	sorting is a fundamental algorithm design problem
	many efficient algorithms use sorting as a 
		"subroutine"
			because easier to process data if the elements are in a sorted order

	general sorting algorithms work in O(n log n) time complexity

	USEFUL:
		inversion : a pair of array elements ( array[a] array[b] )

		the number of inversions indicate how much work is needed to sort the array
		array is completely sorted when there are no inversions

	sorting can occur in O(n log n) 
		"mergesort"
			1. if a = b, do not do anything, because subarray is already sorted
			2. calculate the position of the middle element: k = (a + b) / 2
			3. recursively sort the subarray array[a...k]
			4. recursively sort the subarray array[k + 1...b]
			5. merge the sorted subarray array[a...k] and array[k + 1...b] into
				a sorted subarray


			the recursions consists of O(log n) levels, processing each level takes
			O(n) time

			merging the subarrays array[a...k] and array[k + 1...b] is possible
			linear time because they are already sorted


	"Counting Sort"
		sorts an array in O(n) time assuming that every element in the array is an
		integer between 0...c and c = O(n)

		algorithm creates a bookkeeping array, whose indices are elements of the original array
		
		very efficient algorithm but it can only be used when the constant is small

"""

# Given an array that contains n elements, sort the elements in increasing order
import unittest

def sorting(arr):
	book_keep = [0] * (max(arr) + 1)
	result = []

	for index in range(len(arr)):
		if book_keep[ arr[index] ] > 0:  
			book_keep[ arr[index] ] += 1
		else:
			book_keep[ arr[index] ] = 1

	for tup in enumerate(book_keep):
		index, count = tup
		while count > 0:
			
			count -= 1

	return result

class Test(unittest.TestCase):
	def test_sorting_elements(self):
		actual = sorting([1, 3, 8, 2, 9, 2, 5, 6])
		expected = [1, 2, 2, 3, 5, 6, 8, 9]
		self.assertEqual(actual, expected)

unittest.main(verbosity=2)