package other

import "fmt"

/*

We have an array A of integers, and an array queries of queries.

For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].  Then, the answer to the i-th query is the sum of the even values of A.

(Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)

Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.



Example 1:

Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
Output: [8,6,2,4]
Explanation:
At the beginning, the array is [1,2,3,4].
After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.
After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.
After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.
After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.


Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
1 <= queries.length <= 10000
-10000 <= queries[i][0] <= 10000
0 <= queries[i][1] < A.length
*/
func main() {
	fmt.Println(sumEvenAfterQueries([]int{1,2,3,4}, [][]int{{1,0}, {-3, 1}, {-4, 0}, {2, 3}}))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	// 计算所有偶数的和
	for _, value := range A {
		if value % 2 == 0 {
			sum += value
		}
	}

	result := make([]int, 0)

	for i := 0; i < len(queries); i++  {

		val := queries[i][0]
		index := queries[i][1]

		// 做减法，先看原来的位置是什么，然后在处理
		if A[index] % 2 == 0 { // 如果这个位置原来是个偶数的话，先去掉改值
			sum -= A[index]
		}
		A[index] += val // 赋新的值
		if A[index] % 2 == 0 { // 如果新的值是偶数
			sum += A[index]
		}

		// 麻烦，学会做减法
		//aIndex := A[index]
		//
		//if aIndex % 2 == 0 { // 原来是偶数
		//	if val % 2 == 0 { //现在也是偶数
		//		sum += val
		//	} else { //
		//		sum -= aIndex
		//	}
		//} else { // 原来是奇数
		//	if val % 2 == 0 {
		//
		//	} else {
		//		sum += val + aIndex
		//	}
		//}
		//A[index] = aIndex + val

		result = append(result, sum)

	}
	return result
}