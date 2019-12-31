package other

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxNumberOfBalloons("leetcode"))
}

func maxNumberOfBalloons(text string) int {
	/*
	b 98
	a 97
	l 108
	o 111
	n 110
	*/
	charMap := map[rune]int{98: 10, 97: 0, 108: 0, 111: 0, 110: 0}
	for _, value := range text {
		charMap[value] += 1
	}

	bCount := charMap[98]
	aCount := charMap[97]
	lCount := charMap[108] / 2
	oCount := charMap[111] / 2
	nCount := charMap[110]

	res := math.Min(float64(bCount),math.Min(float64(aCount), math.Min(float64(lCount), math.Min(float64(oCount), float64(nCount)))) )
	return int(res)
}