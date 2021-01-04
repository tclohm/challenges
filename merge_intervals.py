"""
	given a collection of intervals, merge all overlapping intervals

	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

	Input: intervals = [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.
"""

def merge(intervals: List[List[int]]) -> List[List[int]]:
	sorted_list = sorted(intervals, key=lambda x:x[0])
	index = 1
	length = len(sorted_list)

	while index < length:
		if sorted_list[index - 1][0] and sorted_list[index - 1][1] >= sorted_list[index][0]:
			sorted_list[index - 1][1] = max(sorted_list[index][1], sorted_list[index - 1][1])
			sorted_list.pop(index)
			length -= 1
		else:
			index += 1
	return sorted_list