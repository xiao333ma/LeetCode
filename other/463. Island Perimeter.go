package other

import "fmt"

func main() {

	a := [][]int{
		{0,1},

	}

	fmt.Println(islandPerimeter(a))
}


func islandPerimeter(grid [][]int) int {
	sum := 0

	row := len(grid)
	col := len(grid[0])

	for i := 0; i < row ; i++  {
		for j := 0; j < col; j++  {
			tem := 0

			if grid[i][j] == 0 {
				continue
			}
			// 上边
			if i == 0 || grid[i-1][j] == 0 {
				tem += 1
			}

			// 下边
			if i == row - 1 || grid[i+1][j] == 0 {
				tem += 1
			}

			// 左边

			if j == 0 || grid[i][j - 1] == 0 {
				tem += 1
			}

			// 右边

			if j == col - 1 || grid[i][j + 1] ==0 {
				tem += 1
			}


			sum += tem

		}
	}
	return sum
}