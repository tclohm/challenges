def can_form_array(arr: [int], pieces: [[int]]) -> bool:
	hashed = {}

	for piece in pieces:
		hashed[piece[0]] = piece

	build = []

	for digit in arr:
		if digit in hashed:
			build += hashed[digit]

	return build == arr

if __name__ == '__main__':
	print(can_form_array([85], [[85]]))
	print(can_form_array([49, 18, 16], [[16, 18, 49]]))
	print(can_form_array([15, 88], [[88], [15]]))
	print(can_form_array([91, 4, 64, 78], [[78], [4, 64], [91]]))
	print(can_form_array([1, 3, 5, 7], [[2, 4, 6, 8]]))