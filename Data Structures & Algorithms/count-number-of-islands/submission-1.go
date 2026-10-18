func numIslands(grid [][]byte) int {
	count := 0
    for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1'{
				count++
				sink(grid, i, j)
			}
		}
	}
	return count
}

func sink(grid [][]byte, r int, c int) {
	if r < 0 || r > len(grid) - 1 || c < 0 || c > len(grid[r]) - 1  || grid[r][c] != '1'{
		return
	}

	grid[r][c] = '0'
	sink(grid, r + 1 , c)
	sink(grid, r - 1, c)
	sink(grid, r , c + 1)
	sink(grid, r , c - 1)
}