"""

	what power must we raise this base to, in order to get this answer

	log10(100)

	10 is the base
	100 is the answer

	log base 10 of 100 -- what power do we need to raise this base (10) to, to get this answer (100)

	what x gets us our result of 100? Answer is 2

	10**2 = 100

	so,

	log10(100) = 2

	10**x = 100

	logarithms are used to solve for x when x is an exponent

	log10(10**x) = log10(100)

	left side, "what power must we raise 10 to in order to get 10?", answer = x

	x = log10(100)

	Log rules
	simplification: logb(b**x)  = x ... bring varible dwon from exponent
	multiplication: logb(x * y) = logb(x) + logb(y)
	division:		logb(x / y) = logb(x) - logb(y)
	powers:			logb(x**y)  = y * logb(x)
	change of base: logb(x) = logc(x) / logc(b) ... useful for changing the base of a log from b to c

"""

"""
	binary search
	start with mid: big or small than target?
					left or right
					repeat
"""

def b_search(tar, nums):
	# floor index and ceil index as walls around the possible positions
	# so floor_index = -1 (to the left of 0th index) and ceil_index = len(nums) (to the right ond the last index)
	floor = -1
	ceil = len(nums)

	# if there isn't at least 1 index between floor and ceiling,
	# we've run out of guesses and the number must not be present
	while floor_index + 1 < ceil_index:
		#	find the index ~halfway between floor and ceiling
		#	use integer division, we'll never get a "half index"
		distance = ceil_index - floor_index
		half = distance // 2
		guess_index = floor_index + half

		guess = nums[guess_index]
		if guess == tar:
			return True

		if guess > tar:
			ceil_index = guess_index
		else:
			floor_index = guess_index

	return False

# n * 1/2 * 1/2 * 1/2 * ... = 1
# n * (1/2)**x = 1
# log2(n) = log2(2**x)
# log2(n) = 2
# O(log2(n))

def merge_sort(list_to_sort):
	# base case: lists with fewer than 2 elements are sorted
	if len(list) < 2:
		return list_to_sort

	# divide the list in half
	# use integer division, we'll never get "half index"
	mid_index = len(list_to_sort) // 2
	left = list_to_sort[:mid_index]
	right = list_to_sort[mid_index:]

	# sort each hald
	sorted_left = merge_sort(left)
	sorted_right = merge_sort(right)

	# merge sorted halves
	sorted_list = []
	current_index_left = 0
	current_index_right = 0

	# sorted left first element comes next
	# if it's less than sorted right's first
	# element or if sorted right is exhausted
	while len(sorted_list) < len(left) + len(right):
		if (( current_index_left < len(left) ) and 
		(current_index_right == len(right) or 
			sorted_left[current_index_left] < sorted_right[current_index_right])):
			
			sorted_list.append(sorted_left[current_index_left])
			current_index_left += 1
		else:
			sorted_list.append(sorted_right[current_index_right])
			current_index_right += 1
			
	return sorted_list
