"""

hash collision, if both words sum up to a certain number

eg
	l	i 	e 	s
	108 105 101	155

	f	o	e 	s
	102 111 101 155
	===============

	both equal 429

	common solution:

	use pointer in array slot to a linked list holding the counts for all the words that hash to the index

	collisions are rare enough that average lookups in a hash table are O(1)

"""