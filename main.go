package main

import "fmt"

func main() {
	var input []int = []int{100000, 5, 7, 9, 44, 33, 4, 6, 882222226, 453, 244, 234}
	BubbleSort(input)
	fmt.Println(input)
}

func BubbleSort(input []int) {
	len := len(input)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if input[j] > input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}
}
