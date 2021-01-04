import unittest
"""
	
	given a set of coins and our task is to form a sum of money n using the coins

	each coin can be used as many times as we want. 
	What is the minimum amount of coins needed?

"""

def least_amount_of_coinage(money_bag: [int], target: int) -> int:

	results = []
	amount = target

	turns = 0

	while len(money_bag):
		turns += 1
		print(turns)

		value = max(money_bag)

		if amount == 0:

			return results

		if amount > 0:

			amount -= value
			results.append(value)

		if amount < 0:

			amount += value
			removed = money_bag.pop()
			results.pop()
			value = max(money_bag)




class Test(unittest.TestCase):

	def test_it(self):
		coins = [1, 2, 5, 10, 20, 50, 100, 200]
		guess = least_amount_of_coinage(coins, 520)
		expected = [200, 200, 100, 20]
		self.assertEqual(guess, expected)

	def test_ninety_nine(self):
		coins = [1, 2, 5, 10, 20, 50, 100, 200]
		guess = least_amount_of_coinage(coins, 99)
		expected = [50, 20, 20, 5, 2, 2]
		self.assertEqual(guess, expected)
unittest.main(verbosity=2)