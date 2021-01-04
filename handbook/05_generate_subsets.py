"""
	complete search, solve almost any algorithm problem

	generate all possible solution or count the number of solutions

"""
powerset = []
def search(item: int, powerset: list):
	subset = []
	if item == 4:
		# process subset
		print(powerset)
		return powerset
	else:
		search(item+1, powerset)
		powerset = powerset + list(subset) + [item]
		search(item+1, powerset)

def get_power_set(s):
	power_set = [[]]
	for elem in s:
		for subset in power_set:
			power_set = power_set + [list(subset) + [elem]]
	return power_set


def searchQueen(y: int):
	n = 4
	if y == n:
		count += 1
		return
	for x in range(n):
		if column[x] or diag1[x+y] or diag2[x-y+n-1]:
			continue
		column[x] = diag1[x+y] = diag2[x-y+n-1] = 1
		searchQueen(y+1)
		column[x] = diag1[x+y] = diag2[x-y+n-1] = 0

if __name__ == '__main__':
	search(1, powerset)
	print(get_power_set(["rock", "paper", "scissors"]))