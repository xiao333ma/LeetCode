package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit1("RLLLLRRRLR"))
}

// 比较二的解法，用了栈，但是这样判断比较复杂
func balancedStringSplit(s string) int {
	res := 0

	if len(s) <= 0 {
		return res
	}
	stack := make([]uint8, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) != 0 {
			last := stack[len(stack)-1]
			new := s[i]
			if last == new {
				stack = append(stack, s[i])
			} else {
				stack = stack[:len(stack) - 1]
				if len(stack) == 0 {
					res++
				}
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return res
}

// 使用一个字典，看数量一样就可以，这个做法很机智，时间复杂度一样，但是这样更清晰
func balancedStringSplit1(s string) int {
	res := 0
	info := map[uint8]int{76: 0, 82: 0 }
	for i := 0; i < len(s); i++  {
		info[s[i]]++
		if info[76] == info[82] {
			res++
			info[76] = 0
			info[82] = 0
		}
	}
	return res
}