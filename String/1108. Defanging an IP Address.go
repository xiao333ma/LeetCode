package main

import "fmt"

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

func defangIPaddr(address string) string {

	res := ""

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			res += "[.]"
		} else {
			res += string(address[i])
		}
	}
	return res
}