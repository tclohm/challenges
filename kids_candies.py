def kidsWithCandies(candies: [int], extraCandies: int):
        result = [None] * len(candies)
        highest = -1

        for element in candies:
        	if element > highest:
        		highest = element

        for index, number in enumerate(candies):
            if highest <= number + extraCandies:
                result[index] = "true"
            else:
                result[index] = "false"
        return result

if __name__ == "__main__":
	candies = [2,3,5,1,3]
	extra = 3

	print(kidsWithCandies(candies, extra))