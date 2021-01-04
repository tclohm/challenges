import re
def alternate_character(string: str) -> int:
	# check every two characters
	# if they match, delete one, increment the count of deletions
	# should we delete the characters? question is asking for the number of deletions

	# r = re.sub(r"(.)\1+", r"\1", s)
	# count = len(s) - len(r)
	# return count


	deletions = 0
	for index in range(1, len(string)):
		current = string[index - 1]
		neighbor = string[index]
		#print(current, "neighbor:", neighbor)
		if current == neighbor:
			deletions += 1
	return deletions


if __name__ == '__main__':
	print(alternate_character('ABABA'))
	print(alternate_character('AABBAA'))
	print(alternate_character('AAB'))
	print(alternate_character('BBA'))