package main

import "fmt"

// [a^i, b^i] --> take b^i before a^i
// [0, 1] --> take course 1 before taking course 0
func canFinish(numCourse int, prerequisites [][]int) bool {
	// adjancency list
	graph := make(map[int][]int)

	// build adjancency list
	for _, courses := range prerequisites {
		pre_course, next_course := courses[0], courses[1]
		graph[next_course] = append(graph[next_course], pre_course)
	}

	// no redundancy
	visit := make(map[int]struct{})

	var dfs func(course int) bool
	dfs = func(course int) bool {
		// if course is in visit, there is a cycle
		if _, ok := visit[course]; ok {
			return false
		}
		// if the courses value is zero, we have removed the ints
		if len(graph[course]) == 0 {
			return true
		}
		// add the cource into the visit map
		visit[course] = struct{}{}
		// add the prereqs to the visit by running through our adjacency list
		for _, prereq := range graph[course] {
			if !dfs(prereq) { return false }
		}
		// after delete the course from visit
		delete(visit, course)
		// make the course an empty array when we go 
		graph[course] = []int{}

		return true
	}
	
	for i := 0 ; i < numCourse ; i++ {
		if !dfs(i) { return false }
	}

	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1,0}}))
	fmt.Println(canFinish(2, [][]int{{1,0},{0,1}}))
	fmt.Println(canFinish(2, [][]int{{0,1}, {0,2}, {1,3}, {1,4}, {3,4}}))
}