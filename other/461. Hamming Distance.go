package other

import (
	"fmt"
	"strconv"
	"strings"
)

/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.

解题思路：

^ 按位异或，不同则为 1
异或完成后，转为二进制
统计1的个数

*/

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	s := strconv.FormatInt( int64(x ^ y), 2)
	count := strings.Count(s, "1")
	return count

}