package main

import "strings"


func toGoatLatin(S string) string {

	arr := strings.Split(S, " ")

	a := ""
	for idx, value := range arr {
		a += "a"
		s := []byte(value)
		if s[0] == 'a' || s[0] == 'A' ||
			s[0] == 'e' || s[0] == 'E' ||
			s[0] == 'i' || s[0] == 'I' ||
			s[0] == 'o' || s[0] == 'O' ||
			s[0] == 'u' || s[0] == 'U' {
			value += "ma" + a

		} else {
			c := s[0]
			s = s[1:]
			s = append(s, c)
			value = string(s)  + "ma" + a
		}
		arr[idx] = value
	}

	return  strings.Join(arr, " ")

}