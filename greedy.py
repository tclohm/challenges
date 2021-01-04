"""
	GREEDY:
		always makes the choice that seems to be the best at the moment

		locally optimal choice, hope to lead to a globally-optimal solution

		objective function (maximized or minimized)

		has only one shot to compute the optimal solution so that it never goes
		back and reverses the decision
"""

# basic interval scheduling problem

# init, let R be the set of all requests, and let A be empty
# while R is not empty:
# 	choose a request i in R that has the smallest finishing time
# 	add request i to A
# 	delete all requests from R that are not compatible with i
# end while
# return set A as the set of accepted requests

# make algo run in O(n log n)
# 	sort the n request in order of finishing time and labeling them in this order
# 	assume, f(i) <= f(j) when i < j. O(n log n)
# 	O(n) time, construct array, S[1...n] with the property that S[i] contains the value s(i)
# 	select requests by processing the intervals in the order of increasing f(i)

states_needed = set(["mt", "wa", "or", "id", "nv", "ut", "ca", "az"])

stations = {}
stations["kone"] = set(["id", "nv", "ut"])
stations["ktwo"] = set(["wa", "id", "mt"])
stations["kthree"] = set(["or", "nv", "ca"])
stations["kfour"] = set(["nv", "ut"])
stations["kfive"] = set(["ca", "az"])

final_stations = set()

while states_needed:
	best_station = None
	states_covered = set()
	for station, states in stations.items():
		# intersection
		covered = states_needed & states
		if len(covered) > len(states_covered):
			best_station = station
			states_covered = covered
		
	states_needed -= states_covered
	final_stations.add(best_station)

print(final_stations)
		