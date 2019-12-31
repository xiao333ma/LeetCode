package other

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countPrimeSetBits(0, 4))
}

var primMap = map[int]int{2:2, 3: 3, 5: 5, 7: 7, 11: 11, 13: 13, 17: 17, 19: 19, 23: 23, 29: 29}

func countPrimeSetBits(L int, R int) int {

	count := 0
	for i := L; i <= R; i++ {
		count += prim(bits.OnesCount(uint(i)))
	}
	return count
}
/*
采用 与 操作符， 和 1 做 与 操作，这样的话，可以直接统计个数，然后 右移，知道结束
*/
func getOneCount(num int) int {
	count := 0
	for ; num != 0 ;  {
		count += num & 1
		num >>= 1
	}
	return count
}
/*
题目中最大是 [1, 10^6] 所以，这个 而 int 最大是 32 位，也就是最多有 32 个 1，
那么 32 以内的素数就是 primMap

判断一个数是否是素数的话，简单点貌似有 6x +- 1

*/
func prim(num int) int  {
	_, ok := primMap[num]
	if ok {
		return  1
	} else {
		return 0
	}
}