package other

func main() {
	numJewelsInStones("aA", "aAAbbbb")
}

/*

You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3Ò
Example 2:

Input: J = "z", S = "ZZ"
Output: 0
Note:

S and J will consist of letters and have length at most 50.
The characters in J are distinct.


解题思路

利用 map 查找 实现查找字符 O(1) 的复杂度
把 J 处理成一个 map
然后在查找 S 中的字符，是否在 J 中出现，出现 +1

时间复杂度 O(N + M)

*/
func numJewelsInStones(J string, S string) int {

	count := 0
	jLength := len(J)
	jmap := make(map[byte]int, jLength)
	for i := 0; i < jLength; i++ {
		jmap[J[i]] = 1
	}
	sLength := len(S)
	for i := 0; i < sLength; i++ {
		if jmap[S[i]] == 1 {
			count++
		}
	}
	return count
}
