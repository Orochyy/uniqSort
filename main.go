package main

import (
	"fmt"
)

func unique(arr []string) []string {
	occurred := map[string]bool{}
	result := []string{}
	for e := range arr {

		// check if already the mapped
		// variable is set to true or not
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true
			// Append to result slice.
			result = append(result, arr[e])
		}
	}

	return result
}

func main() {
	mass1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	mass2 := []string{"a", "b", "f", "a", "n", "j", "y", "u", "v", "z", "q", "w", "e", "r", "t", "y", "i", "i", "p", "v", "a", "v", "w", "a", "g", "a"}

	var mass3 []string

	for _, m1 := range mass1 {
		for _, m2 := range mass2 {
			if m1 == m2 {
				mass3 = append(mass3, m2)
			}
		}
	}

	fmt.Println(mass3)
	uniqItems := unique(mass3)
	fmt.Println(uniqItems)

}
