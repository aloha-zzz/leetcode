package main

import (
	"fmt"
	"sort"
)

func findEvenNumbers(digits []int) []int {
	if len(digits) <= 2 {
		return nil
	}
	filterMap := map[int]bool{}
	result := []int{}
	for k, v := range digits {
		if v%2 == 0 {
			selectedNumber := removeArr(digits, k)
			result = processNumber(v, selectedNumber, filterMap, result)
			fmt.Println(selectedNumber, v, result)

		}
	}

	sort.Ints(result)
	return result
}

func removeArr(arr []int, k int) []int {
	newInt := []int{}
	if k == len(arr)-1 {
		newInt = append(newInt, arr[:k]...)
		return newInt
	}

	newInt = append(newInt, arr[:k]...)
	newInt = append(newInt, arr[k+1:]...)
	return newInt
}

func processNumber(last int, selectedNumber []int, filterMap map[int]bool, prev []int) []int {
	if len(selectedNumber) < 2 {
		return prev
	}
	for i := 0; i < len(selectedNumber); i++ {
		if i == len(selectedNumber)-1 {
			return prev
		}

		for j := i + 1; j < len(selectedNumber); j++ {
			if selectedNumber[i] != 0 {
				result := selectedNumber[i]*100 + selectedNumber[j]*10 + last

				if !filterMap[result] {
					filterMap[result] = true
					prev = append(prev, result)
				}
			}
			if selectedNumber[j] != 0 {
				result := selectedNumber[i]*10 + selectedNumber[j]*100 + last
				if !filterMap[result] {
					filterMap[result] = true
					prev = append(prev, result)
				}
			}

		}
	}
	return prev
}

func main() {
	digits := []int{2, 1, 3, 0}

	res := findEvenNumbers(digits)
	fmt.Println(res, digits)
}
