import unittest
# given an array of n numbers, our taks is to calculate
# the maximum subarray sum
# largest possible sum of a sequence of consecutive values
# in the array

def max_subarray_sum(arr):
	best, sum_subarray = 0, 0
	for index in range(len(arr)):
		sum_subarray = max(arr[index], sum_subarray+arr[index])
		best = max(best, sum_subarray)
	return best


class Test(unittest.TestCase):
	def test_max_sum(self):
		actual = max_subarray_sum([-1, 2, 4, -3, 5, 2, -5, 2])
		expected = 10
		self.assertEqual(actual, expected)

unittest.main(verbosity=2)