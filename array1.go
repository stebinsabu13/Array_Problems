package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int
	var k, size, element int
	fmt.Print("Enter the size of array:")
	fmt.Scan(&size)
	fmt.Println("Enter the array elements:")
	for i := 0; i < size; i++ {
		fmt.Scan(&element)
		arr = append(arr, element)
	}
	sort.Ints(arr)
	fmt.Println("Array:", arr)
	fmt.Print("Number of largest and smallest:")
	fmt.Scan(&k)
	fmt.Println(k, "th smallest element:", arr[k-1])
	fmt.Println(k, "th largest element:", arr[len(arr)-k])
}
