package test_tahap_1

import (
	"fmt"
	"strings"
)

func SolutionOne() {
	var stringTotal int
	var stringInput string
	var stringMatchIndex []int
	var stringMatch string

	fmt.Print("Masukan jumlah string: ")
	fmt.Scan(&stringTotal)
	arrayString := make([]string, stringTotal)

	for i := 0; i < stringTotal; i++ {
		fmt.Scan(&stringInput)
		arrayString[i] = stringInput
	}

	fmt.Println("strings: ", arrayString)

	for i := 0; i < len(arrayString); i++ {

		if len(stringMatch) == 0 {
			for j := i; j < len(arrayString); j++ {
				if i == j {
					continue
				}
				if strings.EqualFold(arrayString[i], arrayString[j]) {
					stringMatch = arrayString[i]
					stringMatchIndex = append(stringMatchIndex, i+1)
					break
				}
			}
		} else if strings.EqualFold(stringMatch, arrayString[i]) {
			stringMatchIndex = append(stringMatchIndex, i+1)
		}
	}

	if len(stringMatchIndex) == 0 {
		fmt.Println(false)
	} else {
		fmt.Println(stringMatchIndex)
	}
}
