# string rotation
# assume string has the following chars: a[0], a[1], a[2] ... a[n - 1]
# swap position i (i >= 0 && i < n - 1)

# time: O(n log n)
# space: O(N)
def orderly_queue(S: str, K: int) -> str:
	if K == 1:
		L = len(S)
		S = 2 * S
		return min(S[i:i+1] for i in range(L))
	return "".join(sorted(S))
