package other

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(nums1, nums2))

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	res := make([]int, 0)
	nums2Map := make(map[int]int, 0)
	for i := 0; i < len(nums2); i++  {
		nums2Map[nums2[i]] = i
	}

	for i := 0; i < len(nums1); i++ {
		j := nums2Map[nums1[i]]
		next := -1
		for ;j < len(nums2); j++  {
			if nums2[j] > nums1[i] {
				next = nums2[j]
				break
			}
		}
		res = append(res, next)

	}
	return res
}