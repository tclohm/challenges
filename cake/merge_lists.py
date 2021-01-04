import unittest


def merge_lists(my_list, alices_list):

    # Combine the sorted lists into one large sorted list
    merged_list_size = len(my_list) + len(alices_list)
    merged = [0] * merged_list_size
    
    al = alices_list
    ml = my_list
    
    ai, mi, index = 0, 0, 0
    
    while index < merged_list_size:
        is_ml_exhausted = mi >= len(ml)
        is_al_exhausted = ai >= len(al)
        
        if ( not is_ml_exhausted and (is_al_exhausted or ml[mi] < al[ai]) ):
            merged[index] = ml[mi]
            mi += 1
        else:
            merged[index] = al[ai]
            ai += 1
            
        index += 1
    

    return merged















# Tests

class Test(unittest.TestCase):

    def test_both_lists_are_empty(self):
        actual = merge_lists([], [])
        expected = []
        self.assertEqual(actual, expected)

    def test_first_list_is_empty(self):
        actual = merge_lists([], [1, 2, 3])
        expected = [1, 2, 3]
        self.assertEqual(actual, expected)

    def test_second_list_is_empty(self):
        actual = merge_lists([5, 6, 7], [])
        expected = [5, 6, 7]
        self.assertEqual(actual, expected)

    def test_both_lists_have_some_numbers(self):
        actual = merge_lists([2, 4, 6], [1, 3, 7])
        expected = [1, 2, 3, 4, 6, 7]
        self.assertEqual(actual, expected)

    def test_lists_are_different_lengths(self):
        actual = merge_lists([2, 4, 6, 8], [1, 7])
        expected = [1, 2, 4, 6, 7, 8]
        self.assertEqual(actual, expected)


unittest.main(verbosity=2)