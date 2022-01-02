package main

import "fmt"

func getDistances(arr []int) []int64 {
	idxMap := buildMap(arr)
	valList := []int64{}
	preMap, suffixMap := buildPreAndSuffix(idxMap)
	mapCount := map[int]int{}
	for _, v := range arr {
		distance := calcDistance(v, preMap, suffixMap, mapCount)
		valList = append(valList, int64(distance))
	}
	return valList
}

// key 是 元素， value list 是 idx
func buildMap(arr []int) map[int][]int {
	calcMap := map[int][]int{}

	for k, v := range arr {
		calcMap[v] = append(calcMap[v], k)
	}
	return calcMap
}

func calcDistance(v int, pre map[int][]int64, after map[int][]int64, mapCount map[int]int) int64 {

	length := len(pre[v]) - 1

	idx := mapCount[v]
	mapCount[v]++

	return pre[v][idx] + after[v][length-idx]

}

func buildPreAndSuffix(calcMap map[int][]int) (map[int][]int64, map[int][]int64) {
	// 前缀和
	pre := map[int][]int64{}
	for num, idxList := range calcMap {
		prev := -1
		sum := 0
		for k, v := range idxList {
			addValue := (v - prev) * k
			pre[num] = append(pre[num], int64(sum+addValue))
			sum += addValue
			prev = v
		}

	}
	fmt.Println(pre)
	// 后缀和
	after := map[int][]int64{}

	for num, idxList := range calcMap {
		length := len(idxList)
		sum := 0
		idx := 0
		prevIdx := 9999999
		for i := length - 1; i >= 0; i-- {
			addValue := (prevIdx - idxList[i]) * idx

			after[num] = append(after[num], int64(sum+addValue))
			idx++
			prevIdx = idxList[i]
			sum += addValue
		}
	}

	fmt.Println(after)
	return pre, after
}

func main() {
	arr := []int{10, 1, 10, 10, 10, 12, 10, 10}
	res := getDistances(arr)
	fmt.Printf("res: %+v", res)
	fmt.Printf("res: %+v", res)

}
