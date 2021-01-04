"""
	O(n lg n) is the time to beat

	scores already sorted we'd have to do a full walk through the list to conform
	it was in fact fully sorted

	spend at least O(n) time on our sorting function

	O(n) runtime

	we can use greedy or counting

	we can build a score_counts where the indices represent scores and the values
	represent how many times the score appears
"""
import unittest

def sort_scores(unsorted_scores, highest_possible_score):
	# list of 0s at indices 0..highest_possible_score
	score_counts = [0] * (highest_possible_score + 1)

	# populate score_counts
	for score in unsorted_scores:
		score_counts[score] += 1

	# populate the final sorted list
	sorted_scores = []

	# for each item in score_counts
	for score in range(len(score_counts) - 1, -1, -1):
		count = score_counts[score]
		# For the number of times the item occurs
		for time in range(count):
			# Add it to the sorted list
			# print(score)
			sorted_scores.append(score)
	# print(score_counts)
	return sorted_scores


class Test(unittest.TestCase):

    def test_no_scores(self):
        actual = sort_scores([], 100)
        expected = []
        self.assertEqual(actual, expected)

    def test_one_score(self):
        actual = sort_scores([55], 100)
        expected = [55]
        self.assertEqual(actual, expected)

    def test_two_scores(self):
        actual = sort_scores([30, 60], 100)
        expected = [60, 30]
        self.assertEqual(actual, expected)

    def test_many_scores(self):
        actual = sort_scores([37, 89, 41, 65, 91, 53], 100)
        expected = [91, 89, 65, 53, 41, 37]
        self.assertEqual(actual, expected)

    def test_repeated_scores(self):
        actual = sort_scores([20, 10, 30, 30, 10, 20], 100)
        expected = [30, 30, 20, 20, 10, 10]
        self.assertEqual(actual, expected)


unittest.main(verbosity=2)