func numIslands(grid [][]byte) int {
	totalIslands := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				totalIslands++
				sink(grid, i, j)
			}
		}
	}
    return totalIslands
}

func sink(grid [][]byte, r int, c int)  {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	sink(grid,r+1,c)
	sink(grid,r-1,c)
	sink(grid,r,c+1)
	sink(grid,r,c-1)
}