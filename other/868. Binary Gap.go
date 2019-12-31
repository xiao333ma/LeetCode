package other

import "fmt"

func main() {
	fmt.Println(binaryGap(8))
}

func binaryGap(N int) int {
	lastOne := 0
	dis := 0
	for i := 1; N != 0; i++  {
		if N & 1 == 1 { // 1
			if lastOne != 0 {
				newDis := i - lastOne
				if newDis > dis {
					dis = newDis
				}
			}
			lastOne = i
		}
		N = N >> 1
	}
	return  dis
}