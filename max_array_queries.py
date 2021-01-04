def max_array_queries(n, queries):
	arr = [0] * (n + 1)
	for q in queries:
		start, end, k = q

		# increment the value at the start of the sequence
		arr[start-1] += k

		# decrement the value at the first position after the sequence
		arr[end] -= k

	tmp, max = 0, 0

	for x in arr:
		
		tmp += x

		if tmp > max:
			max = tmp
			
	return max

if __name__ == '__main__':
	print(max_array_queries(10, [[1, 5, 3], [2, 4, 10], [4, 5, 2]]))