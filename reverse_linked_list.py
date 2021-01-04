class Node:
	def __init__(self, value=0, next=None):
		self.value = value
		self.next = next

	def __repr__(self):
		if self.next:
			return f"{self.value} -> {self.next}"
		else:
			return f"{self.value}"

def reverse_list(head: Node) -> Node:
	if not head or not head.next:
		return head

	current = head
	prev = None
	next = None

	while current:
		next = current.next
		current.next = previous
		previous = current
		current = next
	return previous

def reverse_list_recurse(head: Node) -> Node:

	def recurse(head, prev=None):
		if not head:
			return prev

		curr = head.next
		head.next = prev
		return recurse(curr, head)

	return recurse(head)


if __name__ == '__main__':
	node5 = Node(5)
	node4 = Node(4, node5)
	node3 = Node(3, node4)
	node2 = Node(2, node3)
	node1 = Node(1, node2)
	print(reverse_list_recurse(node1))