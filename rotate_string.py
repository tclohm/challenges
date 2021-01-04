def get_rotation(array, times):
	length = len(array)
	rotated = []
	for i in range(times, times + length):
		rotated.append([i % length])
	return rotated

# O(n^2)
def rotate_string(A: str, B: str) -> bool:
	if len(A) != len(B):
		return False

		# convert to lists as strings are immutable and it's always easier to play
		# with lists than strings especially if modifications are involved
		A, B = list(A), list(B)
		length = len(A)
		result = A == B
		for i in range(1, length):
			# if an rotation of A matches B, then return True
			if get_rotation(A, i) == B:
				return True
		return result

# O(n)
def rotate_string_eff(A: str, B: str) -> bool:
	return len(A) == len(B) and B in A*2