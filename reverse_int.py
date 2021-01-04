def reverse(x: int) -> int:
	int_to_string = str(x)[::-1]
	if int_to_string[0] == '0':
		int_to_string = int_to_string[1:]
	if int_to_string[-1] == '-':
		int_to_string = '-' + int_to_string[:-2]
	return int(int_to_string)

if __name__ == '__main__':
	print(reverse(123))
	print(reverse(-123))
	print(reverse(120))
	print(reverse(13497294000000))