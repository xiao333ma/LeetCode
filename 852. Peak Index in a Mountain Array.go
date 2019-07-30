package main

import (
	"fmt"
)

/*
Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

Input: [0,1,0]
Output: 1
Example 2:

Input: [0,2,1,0]
Output: 1
Note:

3 <= A.length <= 10000
0 <= A[i] <= 10^6
A is a mountain, as defined above.
*/

func main() {
	fmt.Println(peakIndexInMountainArray([]int{1, 2, 3, 4, 0}))
}

func peakIndexInMountainArray(A []int) int {

	// 直接遍历
	//for i := 0; i < len(A) - 1 ; i++  {
	//	if A[i] > A[i + 1] {
	//		return i
	//	}
	//}
	//return 0

	// 二分法

	low := 0
	height := len(A) - 1

	for ;low < height ;  {
		mid := low + (height - low) / 2
		if A[mid] < A[mid + 1] {
			low = mid + 1
		} else {
			height = mid
		}
	}
	return low
}