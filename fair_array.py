def ways_to_make_fair(nums: [int]) -> int:
	# store the sum, O(n)
	# after deleting the index, all numbers, after the index, in the odd position will become even
	# therefore, all the even position will become odd
	# everything before the index will stay the same

	# if we know the sum of odd and even indexed numbers that are before and after each index
	# we will get our answer
	total_sums = [0, 0]
	sums = [sum(nums[0::2]), sum(nums[1::2])]
	result = 0

	for index, digit in enumerate(nums):
		sums[index % 2] -= digit
		result += total_sums[0] + sums[1] == total_sums[1] + sums[0]
		total_sums[index % 2] += digit
	return result

if __name__ == '__main__':
	print(ways_to_make_fair([6, 1, 7, 4, 1]))
	print(ways_to_make_fair([1, 1, 1, 1, 1, 1]))
	print(ways_to_make_fair([1, 2, 3]))
	print(ways_to_make_fair([1, 1, 1]))
	print(ways_to_make_fair([2, 1, 6, 4]))