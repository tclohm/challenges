import unittest

def reverse(list_of_chars):

	loc = list_of_chars
	length = len(list_of_chars)

	if length == 1:
		return list_of_chars

	for index in range(length // 2):
		loc[index], loc[length - index - 1] = loc[length - index - 1], loc[index]

def reverse_other_solution(list_of_chars):
	left_index = 0
	right_index = len(list_of_chars) - 1

	while left_index < right_index:
		list_of_chars[left_index], list_of_chars[right_index] = \
		list_of_chars[right_index], list_of_chars[left_index]

		left_index += 1
		right_index -= 1


class Test(unittest.TestCase):

    def test_empty_string(self):
        list_of_chars = []
        reverse(list_of_chars)
        expected = []
        self.assertEqual(list_of_chars, expected)

    def test_single_character_string(self):
        list_of_chars = ['A']
        reverse(list_of_chars)
        expected = ['A']
        self.assertEqual(list_of_chars, expected)

    def test_longer_string(self):
        list_of_chars = ['A', 'B', 'C', 'D', 'E']
        reverse(list_of_chars)
        expected = ['E', 'D', 'C', 'B', 'A']
        self.assertEqual(list_of_chars, expected)


unittest.main(verbosity=2)