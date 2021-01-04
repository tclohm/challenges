def smallerThanCurrent(nums: [int]) -> [int]:
	result = [0] * len(nums)
	for index, num in enumerate(nums):
		count = 0
		for iteration in nums:
			if iteration < num:
				count += 1
		result[index] = count
	return result

if __name__ == '__main__':
	numbers = [2, 4, 1, 5]
	print(smallerThanCurrent(numbers))

	numbers2 = [7, 7, 7, 7]
	print(smallerThanCurrent(numbers2))

	numbers3 = [500, 100, 123, 421, 329, 12]
	print(smallerThanCurrent(numbers3))