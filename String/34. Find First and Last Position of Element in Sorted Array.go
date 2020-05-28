package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5,7,7,8, 8, 8,8,10}, 8))
}

func searchRange(nums []int, target int) []int {

	i, j := 0, len(nums)

	for i < j  {
		h := int(uint(j + i) >> 1)
		if nums[h] == target {
			
		}
	}

}



