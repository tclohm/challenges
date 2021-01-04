def numJewelInStones(j: str, s: str) -> int:
	count = 0
	for index in range(len(s)):
		if s[index] in j:
			count += 1
	return count
	# return [index if s[index] in j else None for index in range(len(s))]

if __name__ == '__main__':
	jewels, stone = 'aA', 'aAAbbbb'
	print(numJewelInStones(jewels, stone))