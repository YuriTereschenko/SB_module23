package main

import "fmt"

const size = 10

func insertionSort(input [size]int) [size]int {
	for i := 1; i < size; i++ {
		tmp := i
		for j := i - 1; j >= 0; j-- {
			if input[tmp] >= input[j] {
				break
			}
			input[tmp], input[j] = input[j], input[tmp]
			tmp = j
		}
	}
	return input
}
func main() {
	arr := [size]int{4, 2, 5, 7, 1, 3, 8, 6, 0, 9}
	fmt.Println(insertionSort(arr))
}
