package other

import "fmt"

func main() {
	fmt.Println(projectionArea([][]int{{1,4},{1,1}}))
}


//func projectionArea(grid [][]int) int {
//
//	var res int = 0
//
//	// 找出每一行的最大值，列的值
//	// 找出每一列的最大值，行的值
//
//	for i := 0; i < len(grid); i++  {
//		var rowMax int = 0
//		var colMax int = 0
//		for j := 0; j < len(grid[i]) ;j++  {
//			if grid[i][j] > 0 {
//				res++
//			}
//			if grid[i][j] > rowMax {
//				rowMax = grid[i][j]
//			}
//			if grid[j][i] > colMax {
//				colMax = grid[j][i]
//			}
//		}
//		res = res + rowMax + colMax
//	}
//	return res
//}

func projectionArea(grid [][]int) int {

	var res int = 0

	// 找出每一行的最大值，列的值
	// 找出每一列的最大值，行的值

	for I, iV := range grid {

		var col = 0
		var row = 0

		for J, v := range iV {

			if v > 0 {
				res++
			}
			if v > col {
				col = v
			}

			if grid[J][I] > row {
				row = grid[J][I]
			}
		}
		res = res + row + col
	}
	return res
}