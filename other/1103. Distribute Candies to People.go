package other

import "fmt"

func main() {
	fmt.Println( distributeCandies1(63, 4))
}

func distributeCandies1(candies int, num_people int) []int {

	// 看可以分多少轮

	round := 0
	sum := 0

	for ; round >= 0 ;  {

		sum += ((1 + round *num_people) + ((round + 1)*num_people)) * num_people / 2
		if sum >= candies {
			break
		}
		round++
	}

	res := make([]int, num_people)
	if round > 0 {
		for i := 1; i <= num_people; i++ {
			res[i-1] = i * round + (round - 1) * num_people / 2
			candies -= res[i-1]
		}
	}


	for i := 0;i < num_people ; i++  {
		next := i + 1 + round  * num_people
		if next >= candies {
			res[i] += candies
			break
		} else {
			res[i] += next
			candies -= next
			if candies <= 0 {
				break
			}
		}
	}

	return res
}

