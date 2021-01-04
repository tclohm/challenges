"""
priority queue, take an item out in the priority order

heap is a tree structure and represented as a binary tree

	1	2	3	4	5	6	7	8	9	10

|	2 	2	3	4	7	8	10	7	16	14	|

			1
		   / \
		 2    3
		/ \"  / \
	  4    7 9   10
	 / \" /
	8  16 14

MIN HEAP, each value of nodes is less than or equal to the value of child nodes

parent node always has the smallest value

max heap, parent node has the largest value

access a parent node or a child nodes in the array with indices
- root node 	| i = 1, first item of the array
- parent node 	| parent(i) = i/2
- left child 	| left(i) = 2i
- right child	| right(i) = 2i + 1
"""