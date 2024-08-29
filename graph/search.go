package graph

func IslandDfs(r, c int, grid [][]byte) {
	rows := len(grid)
	cols := len(grid[0])
	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	IslandDfs(r-1, c, grid)
	IslandDfs(r+1, c, grid)
	IslandDfs(r, c-1, grid)
	IslandDfs(r, c+1, grid)
}

func IslandBfs(grid [][]byte, r, c int) {
	q := [][2]int{}

	q = append(q, [2]int{r, c})

	grid[r][c] = '5'

	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, offset := range offsets {
			x := curr[0] + offset[0]
			y := curr[1] + offset[1]

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == '1' {
				q = append(q, [2]int{x, y})
				grid[x][y] = 2
			}
		}
	}
}
