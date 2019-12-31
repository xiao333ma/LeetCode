package main

import "fmt"

func main()  {
	s := []byte{1,2,3,4,5}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte)  {

	i := 0
	j := len(s) - 1

	for i <= j  {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}