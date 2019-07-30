package main

import (
	"fmt"
)

/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number,
also in sorted non-decreasing order.



Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
解题思路：
找到中间位置，
一个往前找
一个往后找
双指针比较，排序

Time Complexity O(n)
Space Complexity O(n)
*/

func main() {
	fmt.Println(sortedSquares([]int{2, 2}))
}

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	// [2, 2]
	j := 0
	for _, value := range A {
		if value <= 0 {
			j++
			break
		}
	}

	i := j - 1


	// 0 - i 负数 i - len(A) - 1正数
	t := 0
	for  ; i >= 0 && j < len(A); t++ {
		i2 := A[i] * A[i]

		j2 := A[j] * A[j]
		if i2 > j2 {
			result[t] = j2
			j++
		} else  {
			result[t] = i2
			i--
		}
	}

	for ;i >= 0 ;i--  {
		result[t] = A[i] * A[i]
		t++
	}

	for ;j < len(A) ;j++  {

		result[t] = A[j] * A[j]
		t++
	}

	return result
}

