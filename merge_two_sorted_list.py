class ListNode:
	def __init__(self, value=0, next=None):
		self.value = value
		self.next = next

	def __repr__(self):
		if self.next is None:
			return f"{self.value}"
		return f"{self.value} -> " + str(self.next) + " "

def mergeTwoLists(list1: ListNode, list2: ListNode) -> ListNode:
	# create a new list node
	# check to see who has the the lower value in list1 or list2
	# create new node List and place the lower value into the node
	# move to new node and point to next
	head = ListNode(None)
	current_node = head
	pointer_one = list1
	pointer_two = list2

	while pointer_one is not None and pointer_two is not None:

		if pointer_one.value is None:
			print("pointer 1 value is none")
			new_node = ListNode(pointer_two.value)
			current_node.next = new_node
			current_node = current_node.next
			pointer_two = pointer_two.next

		elif pointer_two.value is None:
			print("pointer 2 value is none")
			new_node = ListNode(pointer_one.value)
			current_node.next = new_node
			current_node = current_node.next
			pointer_one= pointer_one.next

		elif pointer_one.value <= pointer_two.value:
			new_node = ListNode(pointer_one.value)
			current_node.next = new_node
			current_node = current_node.next
			pointer_one = pointer_one.next
			print("pointer 1 is injected", current_node)
			print("pointer 1 next is ", pointer_one)

		elif pointer_one.value > pointer_two.value:
			new_node = ListNode(pointer_two.value)
			current_node.next = new_node
			current_node = current_node.next
			pointer_two = pointer_two.next
			print("pointer two is injected", current_node)
			print("pointer 2 next is", pointer_two)

	return head.next


if __name__ == "__main__":
	node4 = ListNode(4)
	node2 = ListNode(2, node4)
	node1 = ListNode(1, node2)

	node_4 = ListNode(4)
	node_3 = ListNode(3, node_4)
	node_1 = ListNode(1, node_3)

	print(mergeTwoLists(node1, node_1))

