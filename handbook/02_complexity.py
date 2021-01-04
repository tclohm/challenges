"""
	O(1) 		-- constant time
	O(log n) 	-- logarithmic, havles the input size at each step
	O(sqrt(n))	-- slower than O(log n), faster than O(n)
	O(n)		-- linear
	O(n log n) 	-- sorts the input or data structure
	O(n**2)		-- quadratic, two nested loops
	O(n**3)		-- cubic, three nested loops
	O(2**n)		-- iterates through all subset of the input elements
					ex:
						subset of { 1, 2, 3 }
							{1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}
	O(n!) 		-- all permutations of the input elements,
					ex:
						{ 1, 2, 3 }
							(1, 2, 3), (1, 3, 2), 
							(2, 1, 3), (2, 3, 1), 
							(3, 1, 2), (3, 2, 1)
"""

"""
	An algorithm is polynomial if its time complexity is at most O(n**k) where
	k is a constant

	
"""