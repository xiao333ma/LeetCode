package main

import "fmt"

/*
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.



Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

解题思路

原地置换，交换交换奇偶数，把奇数放后边，偶数放前边

Time Complexity O(n)
Space Complexity O(1)

*/
func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {

	i := 0
	j := len(A) - 1

	for ;i < j ;  {
		if A[i] % 2 > A[j] % 2 {
			temp := A[i]
			A[i] = A[j]
			A[j] = temp
		}
		if A[i] % 2 == 0  {
			i++
		}
		if A[j] % 2 == 1 {
			j--
		}
	}
	return A
}