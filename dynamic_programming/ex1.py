import Math
"""
figure out the overlapping subproblems
"""
def steal_recursive(nums: [int], ith: int=len(nums)-1) -> [int]:
	if i < 0:
		return 0
	# recursion to cover all posibilities
	return Math.max(nums[i] + steal_recursive(nums, i - 2), steal_recursive(nums, i - 1))

# memoization (top down)
# save a result
#
# 
# time complexity: O(n)
# space complexity: O(n)
memo = []
def steal_top_down(nums: [int], ith: int=len(nums) - 1) -> [int]:
	if i < 0:
		return 0
	result = Math.max(nums[i], + steal_top_down(nums, i - 2), steal_top_down(nums, i - 1))
	memo[i] = result
	return result

# tabulation
def steal_bottom_up(nums: [int]) -> [int]:
	tab = [0, nums[0]]
	for index, digit in enumerate(nums):
		tab[index + 1] = Math.max(nums[index] + tab[index - 1], tab[i])
	return tab[len(nums)-1]