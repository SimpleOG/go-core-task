package main

import "fmt"

func FindIntersections(arr1, arr2 []int) (bool, []int) {
	firstValues := make(map[int]struct{}, len(arr1))
	result := make([]int, 0, len(arr1)/2)
	var haveIntersections bool
	for _, v := range arr1 {
		if _, ok := firstValues[v]; !ok {
			firstValues[v] = struct{}{}
		}
	}
	for _, v := range arr2 {
		if _, ok := firstValues[v]; ok {
			result = append(result, v)
			//Чтобы избежать записи одинаковых пересечений несколько раз
			delete(firstValues, v)
		}
	}
	if len(result) > 0 {
		haveIntersections = true
	}
	return haveIntersections, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(FindIntersections(a, b))
}
