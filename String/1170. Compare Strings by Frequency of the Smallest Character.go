package main

import (
	"fmt"
	"sort"
)


func numSmallerByFrequency(queries []string, words []string) []int {

	arr := make([]int, 0)

	for _, s := range words {
		arr = append(arr, getFrequency(s))
	}

	sort.Ints(arr)

	res := make([]int, 0)
	for _, s := range queries {
		count := getFrequency(s)
		res = append(res, len(arr) - sort.SearchInts(arr, count + 1))
	}

	return res
}

func getFrequency(s string) int  {
	m := map[rune]int{}
	c := 'z'
	for _, value := range s {
		if value < c {
			c = value
		}
		m[value]++
	}
	return m[c]
}

func main() {
	fmt.Println(binSearch([]int{0, 1, 2,4, 4, 5, 6}, 4))
}

func binSearch(arr []int, target int) int  {

	i, j := 0, len(arr)
	for i < j {
		m := int(uint(i + j) >> 1)
		if !(arr[m] >= target) {
			i = m + 1
		} else {
			j = m
		}
	}
	return i
}