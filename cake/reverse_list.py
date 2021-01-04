def reverse(list_of_chars):
	loc = list_of_chars
	length = len(list_of_chars)

	if length == 1:
		return list_of_chars

	for index in range(length // 2):
		loc[index], loc[length - index - 1] = loc[length - index - 1], loc[index]

	return list_of_chars

def rvrs(list_of_chars):
	left_index = 0
	right_index = len(list_of_chars) -1
	loc = list_of_chars

	while left_index < right_index:
		loc[left_index], loc[right_index] = loc[right_index], loc[left_index]

		left_index += 1
		right_index -= 1

	return list_of_chars