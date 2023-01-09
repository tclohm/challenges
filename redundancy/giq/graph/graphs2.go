package main

import (
	"fmt"
)

type Queue struct {
	Container []int
}

func (q *Queue) Push(item int) {
	q.Container = append(q.Container, item)
}

func (q *Queue) Pop() int {
	var node = q.Container[0]
	q.Container = q.Container[1:]
	return node
}

func (q *Queue) Length() int {
	return len(q.Container)
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

type adjList [][]int

func makeAdjacentList(n int, managers []int) [][]int {
	var al = make(adjList, len(managers))
	for e := 0 ; e < n ; e++ {
		var manager = managers[e]
		if manager == -1 { continue }
		al[manager] = append(al[manager], e)
	}
	return al
}

func makeAL(n int, courses adjList) [][]int {
	var al = make(adjList, n)
	for i := 0 ; i < len(courses) ; i++ {
		current, next := courses[i][0], courses[i][1]
		al[next] = append(al[next], current)
	}
	return al
}

func dfs(currentId int, adjacentList adjList, informTime []int) int {
	if len(adjacentList[currentId]) == 0 { return 0 }
	var time = 0
	var subordinates = adjacentList[currentId]
	for i := 0 ; i < len(subordinates) ; i++ {
		var otherTime = dfs(subordinates[i], adjacentList, informTime)
		time = max(time, otherTime)
	}
	return time + informTime[currentId]
}

// cyclic? unconnected? weighted? directed?
func numOfMinutes(n, headId int, managers, informTime []int) int {
	al := makeAdjacentList(n, managers)
	return dfs(headId, al, informTime)
}

// Course Scheduler
func canFinish(n int, prerequisites adjList) bool {
	var al = makeAL(n, prerequisites)
	
	for v := 0 ; v < n ; v++ {
		q := Queue{}
		seen := map[int]bool{}

		for i := 0 ; i < len(al[v]) ; i++ {
			q.Push(al[v][i])
		}

		for q.Length() > 0 {
			var current = q.Pop()
			seen[current] = true

			if current == v { return false }

			var adjacent = al[current]

			for i := 0 ; i < len(adjacent) ; i++ {
				var next = adjacent[i]
				if _, ok := seen[next]; !ok {
					q.Push(next)
				}
			}
		}
	}

	return true
}


func main() {
	// a company has n employees with unique IDs from 0 to n-1
	// head of the company has then ID headID

	// each employee has one direct manager
	// company head has no manager, -1
	// tree structure
	totalEmployee := 8
	// employees := []int{0, 1, 2, 3, 4, 5, 6, 7}
	headOfCompany := 4
	// 				  0, 1, 2, 3,  4, 5, 6, 7
	managers := []int{2, 2, 4, 6, -1, 4, 4, 5}

	// 				4
	//		 	  / | \
	//			 5	2  6
	//		 	/  / \  \
	//		   7  0   1  3

	// inform subordinates
	//					0, 1, 2, 3, 4, 5, 6, 7
	informTime := []int{0, 0, 4, 0, 7, 3, 6, 0}

	// number of minutes it takes to inform all employees of the news
	// 			4
	// 	  	  / | \
	// 	 	 5	2  6
	// 		/  / \  \
	//     7  0   1  3

	fmt.Println(numOfMinutes(totalEmployee, headOfCompany, managers, informTime))

	fmt.Println("-----")

	prereqs := [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}
	fmt.Println(canFinish(6, prereqs))
}