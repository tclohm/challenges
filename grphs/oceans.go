package main

import "fmt"

type set struct {
	a int
	b int
}

func oceans(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
    pac, atl := make(map[set]bool), make(map[set] bool)
    
    var dfs func(int, int, map[set]bool, int)
    dfs = func(r, c int, visit map[set]bool, prevHeight int) {
        if (
            visit[set{r, c}] ||
            r < 0 ||
            c < 0 ||
            r == ROWS ||
            c == COLS ||
            heights[r][c] < prevHeight) {
            return;
        }
        visit[set{r, c}] = true
        dfs(r + 1, c, visit, heights[r][c])
        dfs(r - 1, c, visit, heights[r][c])
        dfs(r, c + 1, visit, heights[r][c])
        dfs(r, c - 1, visit, heights[r][c])
    }
    
    for c := 0; c < COLS; c++ {
        dfs(0, c, pac, heights[0][c])
        dfs(ROWS - 1, c, atl, heights[ROWS - 1][c])
    }
    
    for r := 0; r < ROWS; r++ {
        dfs(r, 0, pac, heights[r][0])
        dfs(r, COLS - 1, atl, heights[r][COLS - 1])
    }
    
    res := make([][]int, 0)
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if pac[set{r, c}] && atl[set{r, c}] {
                res = append(res, []int{r, c})
            }
        }
    }
    return res
}
func main(){
	heights := [][]int{{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}
	fmt.Println(oceans(heights))
}