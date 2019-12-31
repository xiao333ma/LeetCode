package other

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	strArr := strings.Split(s, " ")
	for i, world := range strArr {
		strArr[i] = revers(world)
	}
	return strings.Join(strArr, " ")
}

func revers(str string) string   {
	bytes := []byte(str)
	i, j := 0, len(bytes) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
