package main

import (
	"fmt"
	"strings"
)

func numberOfBeams(bank []string) int {
	valueNum := []int{}

	for _, list := range bank {
		tmp := 0

		strArr := strings.Split(list, "")

		for _, v := range strArr {
			if v == "1" {
				tmp++
			}
		}
		valueNum = append(valueNum, tmp)
	}

	if len(valueNum) <= 1 {
		return 0
	}

	sum := 0

	prev := 0

	for _, v := range valueNum {
		if v != 0 {
			sum += prev * v
			prev = v
		}
	}
	return sum
}

func main() {
	bank := []string{
		"011001", "000000", "010100", "001000",
	}

	res := numberOfBeams(bank)
	fmt.Println(res)
}
