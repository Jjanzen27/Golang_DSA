package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftSorted := mergeSort(arr[:mid])
	rightSorted := mergeSort(arr[mid:])

	return merge(leftSorted, rightSorted)
}

func merge(a1 []int, a2 []int) (merged []int) {

	i, j := 0, 0

	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			merged = append(merged, a1[i])
			i++
		} else {
			merged = append(merged, a2[j])
			j++
		}
	}

	for i < len(a1) {
		merged = append(merged, a1[i])
		i++
	}

	for j < len(a2) {
		merged = append(merged, a2[j])
		j++
	}

	return
}

func main() {
	fmt.Println(mergeSort([]int{4, 1, 9, -7, 5, -2, 0}))
	fmt.Println(mergeSort([]int{1, 2, 3, 4, 5, 6}))
}
