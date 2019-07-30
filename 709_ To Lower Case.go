package main

import "fmt"

/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.



Example 1:

Input: "Hello"
Output: "hello"
Example 2:

Input: "here"
Output: "here"
Example 3:

Input: "LOVELY"
Output: "lovely"

解题思路：
直接进行大小写转换即可，字母大小写的 ASCII 差值为 32

a-z：97-122

A-Z：65-90

0-9：48-57

*/

func main() {
	fmt.Println(toLowerCase("ABCDddd"))
}

func toLowerCase(str string) string {
	s := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 'A' && c <= 'Z' {
			c = c + 32
		}
		s[i] = c
	}
	return string(s)
}
