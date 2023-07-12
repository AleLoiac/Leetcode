package main

import "fmt"

func main() {

	s := "XVII"

	fmt.Println(romanToInt(s))

}

func romanToInt(s string) int {

	roman := make(map[string]int)

	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	result := 0

	for i := len(s) - 1; i >= 0; i-- {

		x := roman[string(s[i])]

		if i >= 1 {

			y := roman[string(s[i-1])]

			if x <= y {

				result += x

			} else {

				result += x - y

				i--

			}

		} else {

			result += x

		}

	}

	return result
}
