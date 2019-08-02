package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestRangeI([]int{2, 7, 2}, 1))
}
// 4 6 9
// -2 0 3
func smallestRangeI(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	minB := A[0]
	maxB := A[0]

	for _, value := range A {

		if value < minB {
			 minB = value
		}
		if value > maxB{
			 maxB = value
		}
	}

	res := maxB - minB - 2 * K
	if res < 0 {
		return 0
	}
	return res
}