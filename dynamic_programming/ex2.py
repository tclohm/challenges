"""
DP, solving problems by breaking them down into a collection of simpler subproblems,
solving each of those subproblems just once and storing their solution

recognizing the problem is the most difficult step

can this be expressed as a function of solutions to similar smaller problems?


"""

def can_stop_iterative(runway, init_speed, start_index=0):
	# maximum speed cannot be larger than length of the runway
	max_speed = len(runway)
	if start_index >= len(runway) or start_index < 0 or init_speed < 0 or init_speed > maximum or not runway[start_index]:
		return False
	# { position i : set of speeds for which we can stop from position i }
	memo = {}
	# Base cases, we can stop when a position is not a spoke and speed is zero
	for position in range(len(runway)):
		if runway[position]:
			memo[position] = set([0])
	# outer loop to go over positions from last one to first one
	for position in reversed(range(len(runway))):
		# skip position which contains spikes
		if not runway[position]:
			continue
		# for each position, go over all possible speeds
		for speed in range(1, max_speed + 1):
			# recurrence relation is the same as in the recursive version
			for adjusted_speed in [speed, speed - 1, speed + 1]:
				if (position + adjusted_speed in memo and adjusted_speed in memo[position + adjusted_speed] and adjusted_speed in memo[position + adjusted_speed]):
					memo[position].add(speed)
					break
	return init_speed in memo[start_index]