"""
input string : s
pattern 	 : p


s can be empty and contains only lowercase letters
p can be empty and contains only lowercase letters and chars like ? and *
"""
def isMatch(s: str, p: str) -> bool:
	dynamic_programming = [[False for i in range(len(p)+1)] for j in range(len(s)+1)]
	dynamic_programming[0][0] = True
	for index in range(1, len(p) + 1):
		if p[index - 1] == '*':
			dynamic_programming[0][index] = dynamic_programming[0][index - 1]
	for index in range(1, len(s) + 1):
		for jndex in range(1, len(p) + 1):
			if s[index - 1] == p[jndex - 1] or p[jndex - 1] == '?':
				dynamic_programming[index][jndex] = dynamic_programming[index - 1][jndex - 1]
			elif p[jndex - 1] == "*":
				dynamic_programming[index][jndex] = dynamic_programming[index - 1][jndex] | dynamic_programming[index][jndex - 1]
	return dynamic_programming[-1][-1]



if __name__ == '__main__':
	print(isMatch("aa", "*")) # true

	print(isMatch("cb", "?a")) # false

	print(isMatch("adceb", "*a*b")) # true

	print(isMatch("acdcb", "a*c?b")) # false
