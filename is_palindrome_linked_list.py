class Node(object):
	"""docstring for Node"""
	def __init__(self, value=0, next=None):
		self.value = value
		self.next = next

def reverse(head: Node) -> Node:
	curr = head
	prev = None
	while curr != None:
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	return prev

def is_palindrom(head: Node) -> bool:
	if head is None:
		return True
	rewind, forward = head, head

	while forward and forward.next:
		rewind = rewind.next
		forward = forward.next.next

	rev = self.reverse(rewind)
	forward = head

	while rev != None:
		if forward.val != rev.val:
			return False
		forward = forward.next
		rev = rev.next
	return True
		