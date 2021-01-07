# You are climbing a staircase. It takes n steps to reach the top.
# Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

def climb(n):
	if n == 1:
		return 1
	first = 1
	second = 2
	for index in range(3, n+1):
		third = first + second
		first = second
		second = third
	return second


# OR DP

def climb_stairs(n):
	if n == 1:
		return 1
	dp = [0] * (n + 1)
	dp[1] = 1
	dp[2] = 2

	for index in range(3, n+1):
		dp[index] = dp[index - 1] + dp[index - 2]

	return dp[n]