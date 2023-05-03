package main

import "fmt"

func main() {

	var s = make([]int, 0)
	s = append(s, 3)
	s = append(s, 2)
	s = append(s, 4)
	s = append(s, 8)

	//s = append(s, 6)
	//s = append(s, 1)
	//s = append(s, 2)

	fmt.Printf("%v", twoSum(s, 10))
}

func two(nums []int, target int) []int {

	out := make([]int, 0)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				out = append(out, i)
				out = append(out, j)
				return out
				break
			}
		}
	}
	return out
}

func twoSum(nums []int, target int) []int {

	count := make(map[int]int)
	//fmt.Printf("%v \n", count)
	for i, num := range nums {
		var ok bool
		j, ok := count[num]
		if ok {
			//fmt.Println("Inside if")
			return []int{j, i}
		}
		count[target-num] = i
		//fmt.Printf("After loop: %v \n", count)
	}
	return []int{}
}
