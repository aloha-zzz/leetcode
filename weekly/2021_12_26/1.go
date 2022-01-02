package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	str := strconv.Itoa(num)

	lastWordArr := strings.Split(str, "")

	if lastWordArr[len(lastWordArr)-1] == "0" {
		return false
	}
	return true

}

func main() {
	res := isSameAfterReversals(123)
	fmt.Println(res)
}
