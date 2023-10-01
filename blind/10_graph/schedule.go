package main 

type doulbeArray [][]int

func canFinish(numberOfCourses int, prerequisites doulbeArray) bool {
	graph := make(map[int][]int)

	for _, course := range prerequisites {

		first_course, next_course := course[1], course[0]

		graph[first_course] = append(graph[first_course], next_course)

	}

	visit := make(map[int]struct{})

	var dfs = func(course int) bool
	dfs = func(course int) bool {

		if _, ok := visit[course]; ok {
			return false
		}

		if len(graph[course]) == 0 {
			return true
		}

		visit[course] = struct{}{}

		for _, prerequisite := range graph[course] {
			if !dfs(prerequisite) { return false }
		}

		// remove the course
		delete(visit, course)
		graph[course] = []int{}

		return true

	}
	
	for i := 0 ; i < numberOfCourses ; i++ {
		if !dfs(i) { return false }
	}
	return true
}

func main() {}