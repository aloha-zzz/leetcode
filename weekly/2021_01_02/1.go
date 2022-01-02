package main

import "strings"

func checkString(s string) bool {
	findB := false

	list := strings.Split(s, "")

	for _, v := range list {

		if v == "b" {
			findB = true
		}

		if findB && v == "a" {
			return false
		}
	}
	return true
}
func main() {

}
