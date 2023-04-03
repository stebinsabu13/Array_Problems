package main

import "fmt"

func main() {
	var arr1 = []int{1, 3, 2, 6, 5, 8}
	var arr2 = []int{3, 2, 3, 6, 8}
	var arr3 []int
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				arr3 = append(arr3, arr1[i])
				break
			}
		}
	}
	fmt.Println(arr3)
}
