class Node:
    def __init__(self, value=0, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right

def level_order_bottom(root: Node) -> [[int]]:
	if root is None:
		return []

	queue, result = [root], []

	while queue:
		level = []
		for _ in range(len(queue)):
			node = queue.pop(0)
			if node.left:
				queue.append(node)
			if node.right:
				queue.append(node)
			level.append(node.value)
		result.append(level)
	return result[::-1]