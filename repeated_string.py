# could cycle through the initial string we are given
# count the number of "a" while we're doing that

def repeated(string: str, number: int) -> int:
	# pointer = 0
	# counter = 0
	# result = 0

	# if string == "a":
	# 	return number

	# if "a" not in string:
	# 	return result

	# while pointer < number:

	# 	if counter == len(string):
	# 		counter = 0

	# 	if string[counter] == "a":
	# 		result += 1

	# 	pointer += 1
	# 	counter += 1

	# return result
	if "a" not in string:
		return 0

	if string == "a":
		return number

	fullcycle = (number // len(string)) * string.count('a')
	remainder = string[:number % len(string)].count('a')
	return fullcycle + remainder



if __name__ == '__main__':
	print(repeated("aba", 10)) # aba aba aba a

	print(repeated("a", 1000000))

	print(repeated("abcac", 10))