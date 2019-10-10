package main

import "strings"

func findOcurrences(text string, first string, second string) []string {

	wordArr := strings.Split(text, " ")

	res := make([]string, 0)
	for i := 0; i < len(wordArr) - 2; i++  {
		if wordArr[i] == first && wordArr[i + 1] == second {
			res = append(res, wordArr[i + 2])
		}
	}

	return res
}