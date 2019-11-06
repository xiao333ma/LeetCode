package main

import (
	"fmt"
	"strings"
)

func main() {
	strA := "apple apple"
	strB := "banana"
	fmt.Println(uncommonFromSentences(strA, strB))
}

func uncommonFromSentences(A string, B string) []string {

	//mapA := make(map[string]int, 0)
	//mapB := make(map[string]int, 0)
	//
	//aArray := strings.Split(A, " ")
	//bArray := strings.Split(B, " ")
	//
	//for _, value := range aArray {
	//	mapA[value] += 1
	//}
	//
	//for _, value := range bArray {
	//	mapB[value] += 1
	//}
	//
	//result := make([]string, 0)
	//
	//for str, count := range mapA {
	//	_, ok := mapB[str]
	//	if count == 1 && !ok {
	//		result = append(result, str)
	//	}
	//}
	//
	//for str, count := range mapB {
	//	_, ok := mapA[str]
	//	if count == 1 && !ok {
	//		result = append(result, str)
	//	}
	//}
	//
	//return result

	newStr := A + " " + B
	arr := strings.Split(newStr, " ")

	dict := make(map[string]int, 0)
	for _, value := range arr {
		dict[value] += 1
	}

	result := make([]string, 0)

	for key, value := range dict {
		if value == 1 {
			result = append(result, key)
		}
	}

	return result
}