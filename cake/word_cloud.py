"""
	using class allows us to call methods

	handle duplicate words with different cases, 
	we choose to make a word
	uppercase in our dictionary only if it is always 
	uppercase in the original string
	
	split words by spaces, em dashes, and ellipses
"""

class WordCloudData(object):
	def __init__(self, input_string):
		# count the frequency of each word
		self.words_to_counts = {}
		self.populate_words_to_count(input_string)

	def add_word_to_dictionary(self, word):
		# if the word is already in the dict we increment its count
		if word in self.words_to_counts:
			self.words_to_counts[word] += 1
		# if lowercase version is in the dict, 
		# we know our input word must be uppercase
		# but we only include uppercase words if they're always uppercase
		# so we just increment the lowercase version's count
		elif word.lower() in self.words_to_counts:
			self.words_to_counts[word.lower()] += 1
		# if uppercase version is in the dict,
		# we know our input word must be lowercase
		# since we only include uppercase words if they're always uppercase,
		# we add the lowercase version and give it the uppercase version's count
		elif word.capitalize() in self.words_to_counts:
			self.words_to_counts[word] = 1
			self.words_to_counts[word] += self.words_to_counts[word.capitalize()]
			del self.words_to_counts[word.capitalize()]
		else:
			self.words_to_counts[word] = 1

	def populate_words_to_count(self, input_string):
		current_word_start_index = 0
		current_word_length = 0

		for index, character in enumerate(input_string):
			if index == len(input_string) - 1:
				if character.isalpha():
					current_word_length += 1
				if current_word_length > 0:
					current_word = input_string[current_word_start_index:current_word_start_index + current_word_length]
					self.add_word_to_dictionary(current_word)
					current_word_length = 0
				elif character == ' ' or character == '\u2014':
					if current_word_length > 0:
						current_word = input_string[current_word_start_index:current_word_start_index + current_word_length]
						self.add_word_to_dictionary(current_word)
						current_word_length = 0
				elif character == '.':
					if index < len(input_string) - 1 and input_string[index + 1] == '.':
						if current_word_length > 0:
							current_word = input_string[current_word_start_index:current_word_start_index + current_word_length]
							self.add_word_to_dictionary(current_word)
							current_word_length = 0
				elif character.isalpha() or character == '\'':
					if current_word_length == 0:
						current_word_start_index = index
					current_word_length += 1
				elif character == '-':
					if index > 0 and input_string[index - 1].isalpha() and input_string[index + 1].isalpha():
						if current_word_length == 0:
							current_word_start_index = index
						current_word_length += 1
					else:
						if current_word_length > 0:
							current_word = input_string[current_word_start_index:current_word_start_index + current_word_length]
							self.add_word_to_dictionary(current_word)
							current_word_length = 0

"""

	Good engineers will come up with a solution, great engineers will come
	up with several solution for the given context

	think after the first solution
	how many solutions can you come up with?

"""