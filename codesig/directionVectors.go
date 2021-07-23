func move(matrix [][]int) bool {
	dr := []int{-1, +1, 0, 0}
	dc := []int{0, 0, +1, -1}

	for i := 0 ; i < 4 ; i++ {
		rr := r + dr[i]
		cc := c + dc[i]

		if rr < 0 || cc < 0 {
			continue
		}
		if rr >= R || cc >= C {
			continue
		}

		// finish later
	}
}