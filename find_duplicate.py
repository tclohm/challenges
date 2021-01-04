def find_duplicate(nums: [int]) -> int:
	# current = len(nums) - 1
	# duplicate = nums[current] if len(nums) > 0 else -1
	# found = False
	# index = len(nums) - 2

	# if duplicate == -1:
	# 	return duplicate

	# while not found:
	# 	if current == 0:
	# 		duplicate = None
	# 		break
	# 	elif nums[index] == duplicate:
	# 		found = True
	# 		break
	# 	elif index == 0:
	# 		current -= 1
	# 		index = current - 1
	# 		duplicate = nums[current]
	# 	else:
	# 		index -= 1
	# return duplicate

	h = {}
	for digit in nums:
		if digit in h:
			return digit
		if digit not in h:
			h[digit] = 1
	return -1

	# Floyds
	# slow and fast pointer
	# slow = nums[0]
	# fast = nums[0]

	# while True:
	# 	slow = nums[slow]
	# 	fast = nums[nums[fast]]

	# 	if slow == fast:
	# 		slow = nums[0]
	# 		while slow != fast:
	# 			fast = nums[fast]
	# 			slow = nums[slow]
	# 		return slow
	
if __name__ == '__main__':
	print(find_duplicate([3, 1, 3, 4, 2]))
	print(find_duplicate([1, 3, 4, 2, 2]))