def search(nums: [int], target: int) -> int:
	# array sorted in ascending order is rotated at some pivot unknown
	# return the target value's index or -1 if not found

	# check what from low to mid, 
	# 	if the value of low is greater than mid
	#		pivot has occurred
	# 			if low is less than target and target is less than high
	# 				middle is high
	# 			else
	# 				middle is low
	#	check the right side

	if len(nums) == 0:
		return -1
	
	low = 0
	high = len(nums) - 1

	if nums[low] == target:
		return low

	if nums[high] == target:
		return high

	while low + 1 < high:

		middle = (high + low) // 2

		# print("middle:", middle, "high:", high, "low:", low)

		if nums[middle] == target:
			return middle

		if nums[low] < nums[middle]:
			if nums[low] <= target and target <= nums[middle]:
				high = middle
			else:
				low = middle
		else:
			if nums[middle] <= target and target <= nums[high]:
				low = middle
			else:
				high = middle
	return -1
		
		

if __name__ == '__main__':
	print(search([4,5,6,7,0,1,2], 0))
	print(search([4, 5, 6, 1, 2, 3], 4))
	print(search([5, 6, 10, 11, 13, 40, 1, 2, 3], 4))
	print(search([1, 3], 2))