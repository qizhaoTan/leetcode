package __8

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, m+1)
	}

	if obstacleGrid[0][0] == 0 {
		grid[1][1] = 1
	}
	for x := 1; x <= n; x++ {
		for y := 1; y <= m; y++ {
			if obstacleGrid[x-1][y-1] == 1 {
				continue
			}

			grid[x][y] += grid[x][y-1] + grid[x-1][y]
		}
	}
	return grid[n][m]
}
