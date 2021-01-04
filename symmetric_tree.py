class Node:
	def __init__(self, value=0, left=None, right=None):
		self.value = value
		self.left = left
		self.right = right

def is_symmetric(root: Node) -> bool:
	if root is None:
		return True
	q = [(root.left, root.right)]

	while q:
		curr = q.pop(0)
		left, right = curr
		if not left and not right:
			continue
		elif (not left or not right) or (left.value != right.value):
			return False
		q.append((left.right, right.left))
		q.append((left.left, right.right))
	return True