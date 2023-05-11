package main

import "fmt"

func main() {
	//s := "()[]{}"
	s := "([])"
	fmt.Print(isValid(s))
}

func isValid(s string) bool {

	par := make(map[string]string)
	par["("] = ")"
	par["["] = "]"
	par["{"] = "}"

	open := make([]string, 0)
	for _, p := range s {
		_, prs := par[string(p)] //check for existence of the key, if true open parenthesis, if false closing parenthesis

		if prs == true {
			open = append(open, string(p))
		}

		if prs == false {

			l := len(open)

			if l <= 0 {
				return false
				break
			}

			lastElement := open[l-1]

			check, _ := par[lastElement]

			if check == string(p) {
				open = open[:len(open)-1]
			} else {
				return false
				break
			}

		}
	}
	//fmt.Printf("%s", open)
	if len(open) > 0 { //if the slice is not empty not all open parenthesis were closed
		return false
	}

	return true
}
