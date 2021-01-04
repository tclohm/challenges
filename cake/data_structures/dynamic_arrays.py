"""
	
	low level language like C or Java, we have to specify upfront 
	how many indices we want
	our array to have

	init an array
	the array may have 4 items but the underlying array may have a length of 10
	size: 4, capacity: 10

	under the hood
	make a new bigger array if we are appending to the end of array, usually twice as big

	then copy each element from the old array into the new array
	free up old array
	append new element

	dynamic arrays have a time cost of O(1) for appends, only true for the avg case
	or the amortized cost

	dynamic array advantage is don't have to specify the size ahead of time
	disadvantage is some appends can be expensive

"""