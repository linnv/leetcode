package demo

func islandPerimeter(grid [][]int) int {
	var oneCount int
	var duplicatedEdgeCount int
	for row, v := range grid {
		for column, e := range v {
			if e == 1 {
				oneCount++
				if row != 0 {
					if grid[row-1][column] == 1 {
						duplicatedEdgeCount++
					}
				}
				if column != 0 {
					if grid[row][column-1] == 1 {
						duplicatedEdgeCount++
					}
				}
			}
		}
	}
	return oneCount*4 - 2*duplicatedEdgeCount
}
