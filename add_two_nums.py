class ListNode:
 	def __init__(self, value=0, next=None):
 		self.value = value
 		self.next = next

 	def __repr__(self):
 		if self.next is None:
 			return str(self.value)
 		return str(self.value) + " -> " + str(self.next)

# -- MARK : First Pass Solution --- O(n)^5
def add_two_numbers(l1: ListNode, l2: ListNode) -> ListNode:
	# curr_one, curr_two = l1, l2
	# result_one, result_two = "", ""

	# # O(n)
	# while curr_one is not None and curr_two is not None:
	# 	result_one += str(curr_one.value)
	# 	result_two += str(curr_two.value)
	# 	curr_one, curr_two = curr_one.next, curr_two.next

	# while curr_one is None and curr_two is not None:
	# 	result_two += str(curr_two.value)
	# 	curr_two = curr_two.next

	# while curr_one is not None and curr_two is None:
	# 	result_one += str(curr_one.value)
	# 	curr_one = curr_one.next

	# reverse_one = result_one[::-1] # O(n)
	# reverse_two = result_two[::-1] # O(n)
	# final_result = str(int(reverse_one) + int(reverse_two))[::-1] # O(n)
	# # print(final_result)
	# i = 0

	# new_list = ListNode(final_result[i])
	# head = new_list
	# # O(n)
	# while True:
	# 	if i == len(final_result)-1:
	# 		return head
	# 	i += 1
	# 	new_list.next = ListNode(final_result[i])
	# 	new_list = new_list.next

	carry = 0
	total = 0
	dummy_node = ListNode()
	curr = dummy_node

	while l1 or l2 or carry:

		if (not l1 or not l2) and not carry:
			curr.next = l1 or l2
			break

		total = carry
		if l1:
			total += l1.value
			l1 = l1.next
		if l2:
			total += l2.value
			l2 = l2.next

		carry = total // 10
		total = total % 10
		curr.next = ListNode(total)
		curr = curr.next

	return dummy_node.next


if __name__ == '__main__':
	node3 = ListNode(3)
	node4 = ListNode(4, node3)
	node2 = ListNode(2, node4) # head
	
	
	node4_too = ListNode(4)
	node6 = ListNode(6, node4_too)
	node5 = ListNode(5, node6) # head

	print(add_two_numbers(node2, node5))

	node_5 = ListNode(5)
	node_5_too = ListNode(5)

	print(add_two_numbers(node_5, node_5_too))

	node8 = ListNode(8)
	node1 = ListNode(1, node8)

	node0 = ListNode(0)

	print(add_two_numbers(node1, node0))