def substractProductAndSum(n: int) -> int:
	pod, sod, string_number = 1, 0, str(n)
	for digit in string_number:
		pod *= int(digit)
		sod += int(digit)
	return pod - sod

if __name__ == '__main__':
	number = 234
	print(substractProductAndSum(number))

	number2 = 4421
	print(substractProductAndSum(number2))