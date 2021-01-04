def rotate(nums: [int], k: int) -> None:
	n = len(nums)
	k %= n

	start = count = 0
	while count < n:
		curr, prev = start, nums[start]
		while True:
			next_idx = (curr + k) % n
			nums[next_idx], prev = prev, nums[next_idx]
			curr = next_idx
			count += 1

			if start == curr:
				break
		start += 1
