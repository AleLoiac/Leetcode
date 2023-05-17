package main

import (
	"fmt"
	"math"
)

func main() {

	//x := 123
	//x := 3205600
	x := -32

	fmt.Println(reverse(x))

}

func reverse(x int) int {

	rev := make([]int, 0)

	for x != 0 {
		num := x % 10
		rev = append(rev, num)
		x = x / 10
	}

	result := 0

	for _, num := range rev {
		result = result*10 + num // this conversion takes also care about the "initial 0"
	}

	if result > math.MaxInt32 {
		return 0
	}

	if result < math.MinInt32 {
		return 0
	}

	return result
}
