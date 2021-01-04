class node:
	def __init__(self, data):
		self.data = data
		self.left = None
		self.right = None

def checkBST(root):
	"""
		we only need maximum of all nodes on the left of subtree
		and the minimum of all nodes on the right side of a subtree

		starting at the root, there's no limitations on the value of the root
		so the range we start off with is (-inf, inf)
		whenever we traverse down the right, update the range
		if we go left, update the right side of the range to be the parent's value - 1
		and the left side of the range doesn't change

		if we go right, update the left side of the range to be the parent's value + 1
		and the right side of the range doesn't change

		check the current node's value to make sure it

	"""


if __name__ == '__main__':
	checkBST()
