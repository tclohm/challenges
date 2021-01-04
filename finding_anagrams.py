"""
	
	Finding Anagrams : create a hashmap
							ht = {}
							keys: words formed with alphabets in sorted order
"""

#!/usr/bin/python3
anagram = {}
word_list = { "apt", "car", "pat", "dog", "tap", "god", "arc", "airmen", "marine", "remain"}
for word in word_list:
	word = word.strip()
	hashkey = "".join(sorted(word))
	if hashkey not in anagram:
		anagram[hashkey] = []
	anagram[hashkey].append(word)

for value in anagram.items():
	if (len(value) > 1):
		print(value)