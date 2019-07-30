package main

import "fmt"

/*
Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.

Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.

You may return any answer array that satisfies this condition.



Example 1:

Input: [4,2,5,7]
Output: [4,5,2,7]
Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.


Note:

2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000

解题思路：
1. 新建一个数组，遍历原数组，把奇数偶数放在对应的位置

Time Complexity O(N)
Space Complexity O(N)

2. 在原有的基础上，遍历数组，在所有的下标是偶数位而值为奇数位的时候，找一个下标是奇数而值是偶数的元素进行交换
Time Complexity O(N)
Space Complexity O(1)

*/

func main() {
	fmt.Println(sortArrayByParityII([]int{4,2,5,7}))
}

func sortArrayByParityII(A []int) []int {

	even := 0
	odd := 1
	result := make([]int, len(A))
	for i := 0; i < len(A); i++  {
		if A[i] % 2 == 0 { // 偶数
			result[even] = A[i]
			even += 2
		} else  {
			result[odd] = A[i]
			odd += 2
		}
	}
	return result

}