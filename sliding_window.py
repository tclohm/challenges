# given an array of integers of size `n`
# aim is to calculate the max sum possible for `k` consecutive elements in the array

def max_sum(array: [int], k: int) -> int:
	m = 0
	for index in range(len(array) - k + 1):
		current_sum = 0
		for j_index in range(k):
			current_sum = current_sum + array[index + j_index]
			print("current sum", current_sum, "\nindex", index, "\nj index", j_index, "\n------")
		m = max(current_sum, m)
	return m

# better efficiency,
# 
def max_sum_eff(array: [int], k: int) -> int:
	m = ws = 0
	# computed sum of 1st window, (size k), in order to get the sum of the next overlapping
	# window. We need to elave out the left most item value and add the new (rightmost) item value
	
	# calculate sum of 1st window
	for i in k:
		ws += array[i]

	# slide window from start to end in array
	for i in range(k ,len(array)):
		ws += array[i] - array[i - k]
		m = max(m, ws)

	return m

if __name__ == '__main__':
	print(max_sum([100, 200, 300, 400], 2))
