def max_profit(prices: [int]) -> int:
	if len(prices) < 2:
		return 0

	min_price = prices[0]
	memo = [0] * len(prices)
	memo[0] = 0

	for i in range(1, len(prices)):
		print(memo)
		# if current price is greater than min_price
		if prices[i] >= min_price:
			# store the maximum of memoized previous or the current price minus min price
			memo[i] = max(memo[i-1], prices[i] - min_price)
		else:
			# else we make current price the minimum price
			min_price = prices[i]
			# and store the previous value at the current index into the memoized array and  
			memo[i] = memo[i - 1]
	# return the last item in memoized array
	return memo[len(prices) - 1]


if __name__ == '__main__':
	print(max_profit([7, 1, 5, 4, 6, 4]))