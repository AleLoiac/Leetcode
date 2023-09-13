package main

import "fmt"

func main() {

	x := 1995

	fmt.Println(intToRoman(x))

}

func intToRoman(x int) string {
	y := x
	ints := make(map[int]string)
	ints[1] = "I"
	ints[4] = "IV"
	ints[5] = "V"
	ints[9] = "IX"
	ints[10] = "X"
	ints[40] = "XL"
	ints[50] = "L"
	ints[90] = "XC"
	ints[100] = "C"
	ints[400] = "CD"
	ints[500] = "D"
	ints[900] = "CM"
	ints[1000] = "M"

	result := ""

	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

loop:

	if y > 0 {

		for i := len(numbers) - 1; i >= 0; i-- {

			if y >= numbers[i] {

				result = result + ints[numbers[i]]

				y = y - numbers[i]

				goto loop

			}

		}

	}

	return result

}
