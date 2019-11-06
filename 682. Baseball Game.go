package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []string{"5","-2","4","C","D","9","+","+"}
	fmt.Println(calPoints(arr))
}

func calPoints(ops []string) int {
	sum := 0
	stack := make([]int, 0)
	for i := 0; i < len(ops); i++ {
		if ops[i] == "C" {
			last := stack[len(stack) - 1]
			sum -= last
			stack = stack[:len(stack) - 1]
		} else if ops[i] == "D" {
			last := stack[len(stack) - 1]
			last *=2
			sum += last
			stack = append(stack, last)
		} else if ops[i] == "+" {
			last := stack[len(stack) - 1] + stack[len(stack) - 2]
			sum += last
			stack = append(stack, last)
		} else {
			 num, _ := strconv.Atoi(ops[i])
			 sum += num
			 stack = append(stack, num)
		}
	}
	return sum
}