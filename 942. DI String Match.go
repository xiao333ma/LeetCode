package main

import "fmt"

/*
Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

If S[i] == "I", then A[i] < A[i+1]
If S[i] == "D", then A[i] > A[i+1]


Example 1:

Input: "IDID"
Output: [0,4,1,3,2]

Example 2:

Input: "III"
Output: [0,1,2,3]

Example 3:

Input: "DDI"
Output: [3,2,0,1]


Note:

1 <= S.length <= 10000
S only contains characters "I" or "D".

解题思路：

双指针法，
如果是 I 就把 res[i] 设为最小的
如果是 D 就把 res[i] 设为最大的
然后依次遍历，直到结束

Time Complexity O(N)
Space Complexity O(N)

*/

func main() {
	fmt.Println(diStringMatch("DDI"))
}

func diStringMatch(S string) []int {

	i := 0
	j := len(S)
	result := make([]int, len(S) + 1)

	for t := 0; t < len(S) ; t++  {
		if S[t] == 'I' { // 当前的是最小的
			result[t] = i
			i++
		} else  { // 当前的是大的
			result[t] = j
			j--
		}
	}
	result[len(S)] = j
 	return result

}