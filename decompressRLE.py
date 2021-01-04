def decompressRLE(nums: [int]) -> [int]:
	result = []
	for index in range(0, len(nums)-1, 2):
		value, frequency = nums[index+1], nums[index]
		result += [value] * frequency
	return result

if __name__ == '__main__':
	numbers = [1, 2, 3, 4]
	print(decompressRLE(numbers))