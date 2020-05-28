package main

func main() {
	println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}

func reverseOnlyLetters(S string) string {
	arr := make([]int32, 0)
	for _, value := range S {
		arr = append(arr, value)
	}
	lo := 0
	hi := len(arr) - 1
	for lo < hi {
		loo := isChar(arr[lo])
		hii := isChar(arr[hi])
		if loo && hii {
			arr[lo], arr[hi] = arr[hi], arr[lo]
			lo++
			hi--
		} else if !loo {
			lo++
		} else {
			hi--
		}
	}

	s := ""
	for _, value := range arr {
		s += string(value)
	}

	return s
}

func isChar(c int32) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}