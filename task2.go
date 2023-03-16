package main

import "fmt"

func bubbleSortRev(input ...int) []int {
	for rightShift := 0; rightShift < len(input); rightShift++ {
		for i := 1; i < len(input)-rightShift; i++ {
			if input[i] > input[i-1] {
				input[i], input[i-1] = input[i-1], input[i]
			}
		}
	}
	return input
}
func main() {
	fmt.Println(bubbleSortRev(2, 3, 6, 4, 2, 3, 5, 1))
}
