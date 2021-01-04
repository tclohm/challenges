"""
	company built an in-house calendar
	tool called HiCal. Want to add a feature
	to see the times in a day when everyone is
	available

	don't assume the meetings are in order
"""

def merge_ranges(meetings):
	# sort the meetings
	sorted_meetings = sorted(meetings)
	# put the first meeting into the merged_meetings
	merged_meetings = [sorted_meetings[0]]
	# now go through the rest of the sorted meetings
	for current_merged_start, current_merged_end in sorted_meetings[1:]:
		# take the last start and last end in the merged meetings
		last_merged_start, last_merged_end = merged_meetings[-1]
		# if the current merged start is less than last merged end
		if (current_merged_start <= last_merged_end):
			# change the merged array last element with the last merged start 
			# and max last merged end or current merged end
			merged_meetings[-1] = (last_merged_start, max(last_merged_end, current_merged_end))
		else:
			# if current merged start is greater than last merged start,
			# append the current start and current end
			merged_meetings.append( (current_merged_start, current_merged_end) )

	return merged_meetings