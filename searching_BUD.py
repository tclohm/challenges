A = [13, 27, 35, 40, 49, 55, 59]
B = [17, 35, 39, 40, 55, 59, 60]

def b_search(array, target):
	high = len(array) - 1
	low = 0
	middle = 0
	found = False
	while low <= high:
		middle = (high + low) // 2
		print(low, middle, high, '----', target, array)
		if array[middle] == target:
			found = True
			break

		# im annoyed because i was missing incrementing 
		# and decrementing the middle by 1 after checking
		if array[middle] < target:
			low = middle + 1
		if array[middle] > target:
			high = middle - 1
		
	if found:
		return array[middle]
	return False

def search_for_common(arr1, arr2):
	duplicates = []
	for digit in arr1:
		common = b_search(arr2, digit)
		if common:
			duplicates.append(common)
	return duplicates

print(search_for_common(A, B))