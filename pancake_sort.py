# Reverse arr[0..i]
def flip(arr, index):
	start = 0
	while start < index:
		tmp = arr[start]
		arr[start] = arr[index]
		arr[index] = tmp
		start += 1
		index -= 1

# Returns index of the maximum 
# element in arr[0..n-1]
def find_max(arr, n):
	maximum = 0
	for i in range(0, n):
		if arr[i] > arr[maximum]:
			maximum = i
	return maximum

# Main function that
# sorts given array
# using flip operations
def pancake_sort(arr, n):
	# start from the complete
	# array and one by one
	# reduce current size
	# by one
	curr_size = n
	while curr_size > 1:
		# find index of maximum
		# element in arr[0..curr_size - 1]
		maximum = find_max(arr, curr_size)

		# move maximum element
		# to end of current array
		# if it's not already at
		# the end
		if maximum != curr_size - 1:
			# to move at the end,
			# first move maximum
			# number to beginning
			flip(arr, maximum)

			# now move the maximum
			# number to end by
			# reversing current array
			flip(arr, curr_size - 1)
		curr_size -= 1