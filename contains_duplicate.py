# array of integers, does it contain duplicates
# if yes, true
# otherwise false

def contains_duplicate(numbers: [int]) -> bool:
	cache = set() # could use dictionary with the value being the occurences
	for number in numbers:
		if number in cache:
			return True
		elif number not in cache:
			cache.add(number)
	return False

if __name__ == '__main__':
	print(contains_duplicate([1, 2, 3, 1]))
	print(contains_duplicate([1, 2, 3, 4]))
	print(contains_duplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2]))