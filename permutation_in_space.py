def permutate(arr, perm):
	# iterate through every element in the given arrays
	for i in range(len(arr)):
		# we look at perm to see what is the new index
		index_to_swap = perm[i]
		# check index if it has already been swapped before
		while index_to_swap < i:
			index_to_swap = perm[index_to_swap]

		# swap the position of elements
		arr[i], arr[index_to_swap] = arr[index_to_swap], arr[i]


if __name__ == '__main__':
	A = ["a", "b", "c", "d", "e"]
	P = [4, 3, 2, 0, 1]
	permutate(A, P)