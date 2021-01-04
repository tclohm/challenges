def search(nums: [int], target: int) -> int:
	high = len(nums)-1
	low = 0
	middle = len(nums) // 2
	guess = nums[middle]

	if nums[high] == target:
		return high
	if nums[low] == target:
		return low

	while guess is not target:
		if low == middle or high == middle:
				return -1
		elif guess < target:
			low = middle
			middle = (high + low) // 2
			#print("low", nums[low], "high", nums[high], "middle", nums[middle])
			guess = nums[middle]
		else:
			high = middle
			middle = (high + low) // 2
			#print("high", nums[high], "low", nums[low], "middle", nums[middle])
			guess = nums[middle]
	return middle
			

if __name__ == '__main__':
	numbs = [-1, 0, 3, 5, 9, 12, 14, 32, 34, 45, 46, 78, 99, 100, 101]
	target = 45
	print(search(numbs, target))