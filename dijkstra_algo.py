"""

	update the costs and parents hash tables as the aglo
	progresses

"""

graph = {}
graph["start"] = {}
graph["start"]["a"] = 6
graph["start"]["b"] = 2
graph["a"] = {}
graph["a"]["fin"] = 1
graph["b"] = {}
graph["b"]["a"] = 3
graph["b"]["fin"] = 5

graph["fin"] = {} # the finish node doesn't have neighbors

"""
	[start]--- 6 ----[A]---- 1 ----[fin]
		|			  |				|
		|			  3				|
		|			  |				|
		|----- 2 ----[B]---- 5 ------
"""

infinity = float("inf")
costs = {}
costs["a"] = 6
costs["b"] = 2
costs["fin"] = infinity

parents = {}
parents["a"] = "start"
parents["b"] = "start"
parents["fin"] = None

processed = []

"""
	-while we have nodes to process
	-grab the node that is closest to the start
	-update costs for its neighbors
	-if any of the neighbors costs were updated, update the parents too
	-make this node processed
"""

node = find_lowest_cost_node(costs) # find the lowest cost node we haven't processed
while node is not None: # if processed all nodes, while loop is done
	cost = costs[node]
	neighbors = graph[node]
	for n in neighbors.keys(): # go through all the neighbots of this node
		new_cost = cost + neighbors[n] # if it's cheaper to get this neighbor
		if costs[n] > new_cost: # by going through this node
			costs[n] = new_cost # update the cost for this node
			parents[n] = node # this node becomes the new parent for this neighbor
	processed.append(node) # mark the node as processed
	node = find_lowest_cost_node(costs) # find the next node to process, and loop


def find_lowest_cost_node(costs):
	lowest_cost = float("inf")
	lowest_cost_node = None
	for node in costs: # go through each node
		cost = cost[node]
		if cost < lowest_cost and node not in processed: # if it's the lowest cost so far
														 # and hasn't been processed yet
			lowest_cost = cost # set it as new lowest-cost node
			lowest_cost_node = node
	return lowest_cost_node