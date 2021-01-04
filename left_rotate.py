def left_rotate(array: [int]) -> [int]:
	if not array:
		return array

	left_most_element = array[0]
	length = len(array)

	# take the current element and next element and switch places
	for i in range(length - 1):
		array[i], array[i + 1] = array[i + 1], array[i]

	array[length - 1] = left_most_element
	return array

def left_rotate_k(array: [int], times: int) -> [int]:
	if not array:
		return array

	length = len(array)
	rotated = []

	for i in range(times, times + length):
		# i % length, give us the correct index
		# index can be determined by the number, times, which
		# represents the number of rotations we want to perform
		# on the given array and then return the result
		rotated.append(array[i % length])
	return rotated

def left_rotate_sm(array: [int], times: int) -> [int]:
	if not array:
		return array
	return [arr[i % len(array)] for i in range(times, times + len(array))]


if __name__ == '__main__':
	print(left_rotate_k([1, 2, 3, 4, 5], 2))