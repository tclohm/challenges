import math
def divide(dividend: int, divisor: int) -> int:
	result = dividend / divisor
	return math.floor(result) if result > 0 else math.ceil(result)



if __name__ == '__main__':
	print(divide(10, 3))
	print(divide(7, -3))