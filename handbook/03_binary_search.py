import unittest

# O(log n) -- halves the size at each step
def search_b(arr, target):
	length = len(arr)
	high, low = length - 1, 0

	while (low <= high):
		middle = high + low // 2

		if arr[middle] == target:
			return middle
		elif arr[middle] > target:
			high = middle - 1
		else:
			low = middle + 1

	return -1


# Method 2: make jumps and slow the speed when we get closer to the target
# 			goes left to right, and the initial jump is n/2 then, n/4, n/8, n/16, etc...
# 			O(log n)

def jump_search(arr, target):
	# how far will we jump?
	jump_length = 0
	length = len(arr)
	half = length // 2
	# while our halfing stays above 1
	while half >= 1:
		# and while (jump + half) is less than our array's length and our (jump + half) 
		# is less than our target (this while loop is performed at most twice for each 
		# jump length)
		while (jump_length + half) < length and arr[jump_length + half] <= target:
			# we will add half to our jump
			jump_length += half
		# when the jump+half goes over our length or target, we will half the half and repeat
		# the cycle. Our halves are decreasing every iteration (goes from left to right) 
		half //= 2
	# check
	if arr[jump_length] == target:
		return jump_length

	return -1


# An important use for binary search is to find the position where the value of a 
# function changes

class Test(unittest.TestCase):
	def test_search(self):
		actual = search_b([1, 2, 3, 4, 40, 50, 60, 74, 80, 90, 100], 60)
		expected = 6
		self.assertEqual(actual, expected)

	def test_jump(self):
		actual = jump_search([22, 32, 44, 55, 65, 78, 99, 101, 110, 130, 145, 200], 78)
		expected = 5
		self.assertEqual(actual, expected)

unittest.main(verbosity=2)