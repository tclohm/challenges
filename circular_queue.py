class CircularQueue:
	def __init__(self, k: int):
		self.storage = [None] * k
		self.head = -1 # head, reads
		self.tail = -1 # tail, writes
		self.size = k

	def __str__(self):
		return f"{self.storage}"

	def enqueue(self, value: int) -> bool:
		"""
		insert an element into the circular queue
		if the queue is full, can only dequeue
		
		tail = self.tail + 1 % self.size e.g: (0 + 1) % 3 = 1 (remainder)
											  (1 + 1) % 3 = 2 (remainder)
											  (2 + 1) % 3 = 0
		"""
		if self.is_full():
			return False

		if self.is_empty():
			self.tail = -1
			self.head = 0

		self.tail = (self.tail + 1) % self.size
		self.storage[self.tail] = value

		return True
		

	def dequeue(self) -> bool:
		"""
        Delete an element from the circular queue. Return true if the operation is successful.
        """
		print(self.head, "dequeue")
		if self.is_empty():
			return False

		if self.head == self.tail:
			self.storage[self.head] = None
			self.head = -1
			self.tail = -1
			return True

		self.storage[self.head] = None
		self.head = (self.head + 1) % self.size

		return True


	def peek(self) -> int:
		"""
        Get the front item from the queue.
        """
		if self.is_empty():
			return -1
		return self.storage[self.head]

	def peek_tail(self) -> int:
		"""
        Get the back item from the queue.
        """
		if self.is_empty():
			return -1
		return self.storage[self.tail]

	def is_empty(self) -> bool:
		"""
        queue empty?
        """
		return self.head == -1

	def is_full(self) -> bool:
		"""
        queue full?
        """
		return ((self.tail + 1) % self.size) == self.head


if __name__ == '__main__':
	q = CircularQueue(3)
	q.enqueue("a")
	q.enqueue("b")
	q.enqueue("c")
	q.dequeue()
	q.dequeue()
	q.enqueue("d")
	q.enqueue("e")
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.enqueue("f")
	q.dequeue()
	q.enqueue("g")
	q.enqueue("h")
	q.enqueue("i")
	print(q)
	q.dequeue()
	q.dequeue()
	q.dequeue()
	print(q)