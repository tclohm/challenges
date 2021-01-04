import re

def decode_string(s: str) -> str:
	"""
		3[a]2[bc]
		3*a+2*bc

		3[a2[c]]
		3*a2c++

		2[abc]3[cd]ef
		2*abc+3*cd+ef

		if our character is not a digit,
			add it to our current string
		if our character is a digit,
			add it to our current number, another number can come after it,
			take the current number and times by 10 then added the number
			(i.e.: )
	"""
	stack = []
	current_number = 0
	current_string = ""

	for char in s:
		if char == "[":
			# print("curr", current_string)
			stack.append(current_string)
			stack.append(current_number)
			current_string = ""
			current_number = 0
		elif char == "]":
			number = stack.pop()
			string = stack.pop()
			current_string = string + (number * current_string)
		# check to see if the character is a digit, then add it to the current_number * 10
		# (ie, 1 * 10 = 10, 2 * 10 = 20)
		elif char.isdigit():
			current_number = (current_number * 10) + int(char)
		else:
			# add the characters to the current string
			current_string += char
	return current_string 




if __name__ == '__main__':
	print(decode_string("3[a]2[bc]"))
	print(decode_string("4[a2[c]]"))
	print(decode_string("10[ab]"))
