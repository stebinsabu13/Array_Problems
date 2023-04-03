//Program to check if sum of any two elements of an array is 10

package main

import "fmt"

func main() {
	var (
		arr           []int
		size, element int
	)
	fmt.Println("enter the size of array:")
	fmt.Scan(&size)
	fmt.Println("enter the array elements:")
	for i := 0; i < size; i++ {
		fmt.Scan(&element)
		arr = append(arr, element)
	}
	fmt.Println(arr)
	var check bool
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i]+arr[j] == 10 {
				check = true
				fmt.Println("Sum of the elements", arr[i], "and", arr[j], "is 10")
			}
		}
	}
	if check == false {
		fmt.Println("No elements with sum is 10")
	}
}
