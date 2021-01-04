# reverse a linked list from position m to n, in one pass
# 1 <= m <= n <= length of list
"""
	reversing a linked list, is like reversing an array
	however reversing an array we have the advantage of having the index available
	get two pointers and move both pointers toward the center of the array

	two pointers, one at n and the other at m
	swap the data between the nodes and move the pointers towards each other

	if you use recursion, we can to simulate the backward pointer.
	backtracking process in a recursion will help us in simulating the backward movement
	of the pointer from the nth node in the linked list towards the center
"""

"""
	process all the nodes at-most twice. Once during the normal recursion process and once
	during the backtracking process
	
"""
def reverse_between(head, m, n):
	if not head:
		return None

	left, right = head, head
	stop = False

	def recurse(right, m, n):
		nonlocal left, stop

		if n == 1:
			return

		right = right.next

		if m > 1:
			left = left.next

		recurse(right, m - 1, n - 1)

		if left == right or right.next == left:
			stop = True

		if not stop:
			left.val, right.val = right.val, left.val

			# right pointer moves one step back via backtracking
			left = left.next

	recurse(right, m, n)
	return head

def reverse_between_iter(head, m, n):

	if not head:
		return None

	prev, curr = None, head

	# push curr until we get m, prev follows
	while m > 1:
		curr = curr.next
		prev = curr
		m, n = m - 1, n - 1

	# pointer to fix the final connections
	# tail is what will be the tail in the reverse list
	# con points to the node one before tail of reverse sublist
	# this connects to the new head of the reversed sublist
	tail, con = curr, prev

	while n:
		third = curr.next
		curr.next = prev
		prev = curr
		curr = third
		n -= 1

	if con:
		con.next = prev
	else:
		head = prev
	tail.next = curr
	return head