package main

import (
	"fmt"
	"strings"
)

func executeInstructions(n int, startPos []int, s string) []int {
	res := []int{}
	cmd := strings.Split(s, "")

	for k := range cmd {
		num := execCommand(n, startPos, cmd[k:])
		res = append(res, num)
	}
	return res
}

func execCommand(n int, startPos []int, cmd []string) int {
	num := 0
	pos := startPos
	for len(cmd) > 0 {
		isSuccess, newPos := execCommandStep(n, pos, cmd)
		if isSuccess {
			num++
			cmd = cmd[1:]
			pos = newPos
		} else {
			return num
		}
	}
	return num
}

func execCommandStep(n int, Pos []int, cmd []string) (bool, []int) {
	switch cmd[0] {
	case "U":
		if Pos[0] <= 0 {
			return false, nil
		}
		return true, []int{Pos[0] - 1, Pos[1]}
	case "D":
		if Pos[0] >= n-1 {
			return false, nil
		}
		return true, []int{Pos[0] + 1, Pos[1]}
	case "L":
		if Pos[1] <= 0 {
			return false, nil
		}
		return true, []int{Pos[0], Pos[1] - 1}
	case "R":
		if Pos[1] >= n-1 {
			return false, nil
		}
		return true, []int{Pos[0], Pos[1] + 1}
	}
	return true, nil
}

func main() {
	n := 3
	startPos := []int{0, 1}
	s := "RRDDLU"
	res := executeInstructions(n, startPos, s)
	fmt.Println(res)

}
