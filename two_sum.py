def twoSum(numbers: [int], target: int) -> [int]:
	cache = {}
	for index, digit in enumerate(numbers):
		remainder = target - digit
		print(remainder)
		if remainder in cache:
			return [cache[remainder], index]
		else: 
			cache[digit] = index

# runs in O(n * log(n)) time and O(1) space
def two_sum(numbers: [int], target: int) -> [int]:
	sort = sorted(numbers)
	left = 0
	right = len(numbers) - 1

	targets = None

	while left < right:
		curr_sum = sort[left] + sort[right]
		if curr_sum == target:
			targets = [left, right]
			break
		elif curr_sum > target:
			right += 1
		else:
			left += 1

	if not targets:
		return None

	return [numbers.index(sort[targets[0]]), numbers.index(sort[targets[1]])]


if __name__ == '__main__':
	nums = [2, 7, 11, 15]
	tar = 9
	print(twoSum(nums, tar))