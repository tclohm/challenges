def numberOfSteps(num: int, steps: int=0) -> int:
	if num == 0:
		return steps
	else:
		if num % 2 == 0:
			steps += 1
			num /= 2
			return numberOfSteps(num, steps)
		else:
			steps += 1
			num -= 1
			return numberOfSteps(num, steps)


if __name__ == "__main__":
	number = 123
	print(numberOfSteps(number))
