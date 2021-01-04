matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]

def unroll(matrix):
	result = []
	if len(matrix) == 0:
		return matrix

	# get top row
	for index, column in enumerate(matrix[0]):
		result.append(column)

	print(result)
	return -1
	# get last element of all remaining rows
	for row in matrix:
		result.append(row.pop())

	# get all remaining on bottom row
	for row in matrix[-1]:
		result.append(row.pop())

	# get remaining on the left
	for row in matrix:
		result.append(row.pop().reverse())

	# get remaining on first element of rows
	for index, row in enumerate(matrix):
		result.append(row.pop(index))

	return unroll(matrix)


if __name__ == '__main__':
	print(unroll(matrix))