package main

import "fmt"

const size2 = 8

func main() {
	arr := [size2]int{2, 3, 6, 4, 2, 3, 5, 1}
	arrRevSorted := func(input [size2]int) [size2]int {
		for rightShift := 0; rightShift < len(input); rightShift++ {
			for i := 1; i < len(input)-rightShift; i++ {
				if input[i] > input[i-1] {
					input[i], input[i-1] = input[i-1], input[i]
				}
			}
		}
		return input
	}(arr)
	fmt.Println(arrRevSorted)
}
