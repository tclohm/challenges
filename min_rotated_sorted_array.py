# an array sorted ascending order is rotated
# at some pivot unknown to you beforehand

# sorted, rotated, doesn't contain any duplicate elements

# search the entire array and find the minimum element

# O(N), if there are N elements in the given array
# O(1)
def find_min(nums: [int]) -> int:
	minimum_element = float('inf')

	for n in nums:
		minimum_element = min(minimum_element, n)
	return minimum_element


# find mid element of the array
# mid element > first element of array, this means that we need to look
# for the inflection point on the right of mid
# if mid element < first element of array this that we need to look for
# the inflection point on the left of mid
# time: O(log n)
# space: O(1)
def find_min_eff(nums: [int]) -> int:
	if len(nums) == 1:
		return nums[0]

		left = 0
		right = len(nums) - 1

		if nums[right] > nums[0]:
			return nums[0]

		# Binary search way
		while right >= left:
			mid = (left + right) / 2

			if nums[mid] > nums[mid + 1]:
				return nums[mid]

			if nums[mid - 1] > nums[mid]:
				return nums[mid]

			if nums[mid] > nums[0]:
				left = mid + 1
			else:
				right = mid - 1