class ListNode:
	def __init__(self, x, next=None):
		self.value = x
		self.next = next

	def __str__(self):
		if self.next is None:
			return f"{self.value}"
		return f"{self.value} -> " + str(self.next)

def delete(head, target_node):
	# delete the node given
	# current and find the target node
	current = head
	while current is not None:
		if current.next.value == target_node.value:
			current.next = current.next.next
			return head
		else:
			current = current.next


if __name__ == '__main__':
	node9 = ListNode(9)
	node1 = ListNode(1, node9)
	node5 = ListNode(5, node1)
	node4 = ListNode(4, node5)
	# print(delete(node4, node1))

	print(delete(node4, node5))