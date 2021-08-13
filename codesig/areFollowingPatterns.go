func areFollowingPatterns(strings []string, patterns []string) bool {
	if len(strings) != len(patterns) {
		return false
	}
    
	matchS := compare(strings, patterns)
	matchP := compare(patterns, strings)
	return matchS && matchP
}

func compare(a []string, b []string) bool {
	match := make(map[string]string)

	for idx, v := range a {
		if match[v] != "" && match[v] != b[idx] {
			return false
		}
		match[v] = b[idx]
	}

	return true
}