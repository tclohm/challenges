"""
	instead of storing the strings right inside our array
	let's put the strings wherever we can fit them in memory
	then we have each element in our array hold the address
	in memory of its corresponding string

	each address is an integer, so really our outer array
	is just an array of integers, these are the pointers 
	(points to another spot in memory)

	fixes disadvantages of arrays:
		1. items dont to be the same length -- each string can be as long or as short
		2. dont need enough uninterrupted free memory to store all our strings next
		to each other -- we can place each of them separately, 
		wherever there's space in RAM

	original array is very cache-friendly, extra speedup from processor cache

	pointers in array are not cache-friendly

	pointer-based array requires less uniterrupted memory and can
	accommodate elements that aren't all the same size, but it's slower
	because it's not cache friendly
"""