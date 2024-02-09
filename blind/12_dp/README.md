# Thinking in DP

1. Category
Most DP question can be boiled down to a few categories

Question: Identify this problem as one of the categories below before continuing
Options:
0/1 Knapsack (Answer)
Unbounded Knapsack
Shortest Path
Fibonacci Sequence
Longest Common Substring / Subsequence

Our Capacity is the target we want to reach 'S'
Our items are the numbers in the input subset and
the weights of the items are the values

2. States

Question: Determine state variables
Answer: Index and Current Sum
General rule, index is a required state in nearly all DP problems
except for shortest path

3. Decisions
DP is about making the optimal decisions
Decisions will have to bring us closer to the base case
and lead us towards the question we want to answer

Question: What decisions do we have to make each recursive call?
Answer: problem requires we take all items in our input subset
so at every step we will be adding an item to our knapsack
The decision is:
	a. add the current number positive value
	b. add the current number neagtive value

4. Base Case