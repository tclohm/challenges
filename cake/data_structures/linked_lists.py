"""
	node ( can be seen as a two-item array )
	character itself and a pointer to the next character
	
	series of nodes connected is a linked-list

	no need for order in memory

	first node in linked list is the head
	last node is the tail

	linked lists worst-case O(1) time appends much better than dynamic arrays appends O(n)

	linked lists prepend faster than dynamic arrays

"""

def get_ith_item_in_linked_list(head, i):
	if i < 0:
		raise ValueError(f"I can't be negative: {i}")

	curr = head
	curr_pos = 0
	while curr:
		if curr_pos == i:
			return curr

		curr = curr.next
		curr_pos += 1

	raise ValueError(f"List has fewer than i + 1 {i + 1} nodes")