package other

import "fmt"

func main() {
	fmt.Print(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {

	res := make([]byte, 0)
	for i := 0; i < len(S); i++  {
		if len(res) > 0 && res[len(res) - 1] == S[i] {
			res = res[:len(res) - 1]
		} else {
			res = append(res, S[i])
		}
	}
	return string(res[:])
}