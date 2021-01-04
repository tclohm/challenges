"""

	hash function takes data and outputs a hash, a fixed-size string or number
	
	hash as a "fingerprint." File will always have the same hash, but we can't go 
	from the hash back to the original file

	multiple files having the same hash value, hash collision

	Dictionaries: use hash function to translate keys into list indices

	Prevent man-in-the-middle-attacks: When you finish the download, try hashing the file
	and confirm you got the same result. If not, your internet provider or someone else might
	have injected malware or tracking software into your download
"""

"""

	Hash Table, organizes data you can quickly look up values for a given key

	Strengths:
		Fast lookups
		Flexible keys

	Weaknesses:
		lookup O(n), worst case
		unordered
		single-directional lookups, looking up the keys for a given value requires looping
			data set O(n)
		not cache-friendly, many use linked lists, not next to each other in memory

"""