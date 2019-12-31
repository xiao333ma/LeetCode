package other

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1,1,1,1,2,2,2,3,3,3}))
}

func distributeCandies(candies []int) int {

	dict := make(map[int]int, l)
	for _, value := range candies {
		dict[value]++
	}

	candiesKind := len(dict)

	everyOneCandiesKind := len(candies) / 2

	if candiesKind >= everyOneCandiesKind {
		return everyOneCandiesKind
	} else {
		return candiesKind
	}
}

