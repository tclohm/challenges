test_string = "abc"

"""
Case "a" -> {"a"}
Case "ab" -> {"ab", "ba"}
Case "abc" -> {"abc", "acb", "cab", "bac", "bca", "cba"}
"""
def to_string(lst):
	return ''.join(lst)

def permute(arr, left, right, result=[]):
	if left == right:
		result.append(to_string(arr))
	else:
		for index in range(left, right+1):
			arr[left], arr[index] = arr[index], arr[left]
			permute(arr, left + 1, right)
			arr[left], arr[index] = arr[index], arr[left]
	return result


print(permute(list(test_string), 0, len(test_string)-1))