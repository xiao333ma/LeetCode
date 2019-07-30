package main

/*

In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.

Return the element repeated N times.



Example 1:

Input: [1,2,3,3]
Output: 3
Example 2:

Input: [2,1,2,5,3,2]
Output: 2
Example 3:

Input: [5,1,5,2,5,3,5,4]
Output: 5


Note:

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length is even

解题思路：
找到重复的返回，利用 map 去重

Time Complexity O(n)

Space Complexity O(n)

*/

func main() {
	repeatedNTimes([]int{2, 1, 2, 5, 3, 2})
}

func repeatedNTimes(A []int) int {

	result := make(map[int]bool, len(A))

	for _, value := range A {
		_, ok := result[value]
		if ok {
			return value
		} else {
			result[value] = true
		}
	}
	return 0

}