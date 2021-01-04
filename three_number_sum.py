def three_number_sum(arr, target):
	ht = {}
	result = []
	
	for index, digit in enumerate(arr):
		ht[digit] = index

	for digit in arr:
		for next_digit in arr:
			remainder = abs(digit + next_digit)
			if remainder in ht and digit is not remainder and next_digit is not remainder and digit is not next_digit:
				r = sorted([digit, next_digit, remainder])
				if (digit + next_digit + remainder == target) and r not in result:
					result.append(r)
	return result

if __name__ == '__main__':
	print(three_number_sum([12, 3, 1, 2, -6, 5, -8, 6], 0))