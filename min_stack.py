# Every operation needs to be in constant time
class MinStack:

	def __init__(self):
		"""
		initizalize data structure here

		what are we going to use for the storage of elements?
		
		"""
		self.storage = []

	def push(self, x: int) -> None:
		if len(self.storage) > 0:
			current_min = min(x, self.storage[-1][1])
		else:
			current_min = x

		self.storage.append( ([x, current_min]) )

	def pop(self) -> None:
		self.storage.pop()

	def top(self) -> int:
		return self.storage[-1][0]

	def getMin(self) -> int:
		return self.storage[-1][1]