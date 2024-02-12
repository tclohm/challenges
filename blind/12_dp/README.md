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
Related directly to the conditions required by the answer we are seeking
Question: Identify the base cases.
Answer: We need 2 base cases. One for when the current state is valid and one for when the current state is invalid.

5. Code it
6. Optimize


## Examples
0/1 Knapsack

    Characteristics: You have a set of items, each with a weight and a value. You can choose each item only once. The goal is to maximize the value while staying within a weight limit.

    Core Logic: Decide for each item whether to include it in the knapsack or not.

    Identification: If you have a resource limit (like weight) and you need to optimize some value (like cost), it's likely a 0/1 Knapsack problem.

Unbounded Knapsack

    Characteristics: Similar to 0/1 Knapsack, but you can include each item multiple times.

    Core Logic: Decide for each item how many times to include it.

    Identification: If you have a resource limit but can use the same item multiple times, you're dealing with Unbounded Knapsack.

Shortest Path (e.g., Unique Paths I/II)

    Characteristics: You have a grid or a graph, and you need to find the shortest path from one point to another.

    Core Logic: Calculate the minimum cost to reach each point from the starting point.

    Identification: If you need to find the least costly way to get from point A to point B in a grid or network, consider this category.

Fibonacci Sequence (e.g., House Thief, Jump Game)

    Characteristics: Problems where each state depends on one or more previous states, like Fibonacci numbers, where each number depends on the previous two.

    Core Logic: Store previous results to calculate the current result.

    Identification: When you can express the current state in terms of previous states, it's often in this category.

Longest Common Substring/Subsequence

    Characteristics: You have two sequences, and you need to find the longest common subsequence (order matters) or substring (continuous sequence).

    Core Logic: Compare elements in both sequences to find common elements and store them in a way that allows backtracking the longest sequence.

    Identification: If you're looking for the longest shared characteristics between two sequences, this is likely your category.

Differentiating Characteristics:

    Resource Constraints: 0/1 and Unbounded Knapsack problems have a resource limit (e.g., weight, volume).

    Reusability: Unbounded Knapsack allows reusing items; 0/1 does not.

    Network or Grid: Shortest Path problems usually involve a grid or network.

    Dependence on Previous States: Fibonacci Sequence problems have a direct dependence on previous states.

    Multiple Sequences: Longest Common Subsequence/Substring problems typically involve more than one sequence to compare.

Understanding these core aspects will help you identify the type of DP problem.


1. Fibonacci Numbers
Climbing Stairs
House Robber
Fibonacci Number
Maximum Alternating Subsequence Sum

2. Zero / One Knapsack
Partition Equal Subset Sum
Target Sum

3. Unbounded Knapsack
Coin Change
Coin Change II
Minimum Cost for Tickets

4. Longest Common Subsequence
Longest Common Subsequence
Longest Increasing Subsequence
Edit Distance
Distinct Subsequences

5. Palindromes
Longest Palindromic Substring
Palindromic Substrings
Longest Palindromic Subsequence