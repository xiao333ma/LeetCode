package main

import "fmt"

func main() {

	fmt.Println(commonChars([]string{"cook","locko","cook"}))
	fmt.Println((1 << 31) - 1)
}


//func commonChars(A []string) []string {
//	resultArray := make([]string, 0)
//	arrayWithMap := make([]map[rune]int, len(A))
//
//	for idx, str := range A {
//
//		charMap := make(map[rune]int, 0)
//
//		for _, char := range str {
//			num, ok := charMap[char]
//			 if ok {
//				 charMap[char] = num + 1
//			 } else  {
//			 	charMap[char] = 1
//			 }
//		}
//		arrayWithMap[idx] = charMap
//	}
//
//	for _, char := range A[0] {
//		exist := true
//		for _, charMap := range arrayWithMap {
//			num, ok := charMap[char]
//			if ok && num > 0 {
//				charMap[char] = num - 1
//			} else  {
//				delete(charMap, char)
//				exist = false
//				break
//			}
//		}
//		if exist {
//			resultArray = append(resultArray, string(char))
//		}
//	}
//	return resultArray
//}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonChars(A []string) []string {
	m := make([]int, 26)

	for i := 0; i < 26; i++ {
		m[i] = (1 << 31) - 1
	}

	for _, str := range A {
		tmp := make([]int, 26)
		for _, c := range str {
			tmp[byte(c)-'a']++
		}
		for c, num := range tmp {
			m[c] = min(num, m[c])
		}
	}

	out := make([]string, 0)
	for idx, count := range m {
		if count > 0 {
			for i := count; i > 0; i-- {
				out = append(out, string(byte(idx + 'a')))
			}
		}

	}
	return out
}
