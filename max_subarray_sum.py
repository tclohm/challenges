# given an array of n numbers, task is to calculate the maximum subarray sum
# the largest possible sum of a sequence of consecutive values in the array

# we can assume that an empty subarray is allowed, so the maximum subarray sum is
# always at least 0

# O(n^3)
def max_sum(array: [int]) -> int:
	best = 0
	# index and k_index fix the first and last index of the subarray, and the sum value
	# is calculated to the variable sum
	for index in range(len(array)):
		for k_index in range(index, len(array)):
			summ = 0
			for j_index in range(index, k_index):
				summ += array[j_index]
			best = max(best, summ)
	return best

# changes to O(n^2)
def max_sum_efficient(array: [int]) -> int:
	best = 0
	for index in range(len(array)):
		summ = 0
		for k_index in range(index, len(array)):
			summ += array[k_index]
			best = max(best, summ)
	return best

# For O(n), calculate for each array position, the max sum of a subarray that ends at that position
# 2 possibilities
# subarray only contains the element at position k
# subarray consists of a subarray that ends at position k - 1, followed by the element at position k

def max_sum_more_efficient(array: [int]) -> int:
	best = summ = 0
	for index in range(len(array)):
		summ = (max(array[index], summ + array[index]))
		best = max(best, summ)
	return best

if __name__ == '__main__':
	rr = [-1, 2, 4, -3, 5, 2, -5, 2]
	print(max_sum_more_efficient(rr))