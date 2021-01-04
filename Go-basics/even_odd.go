package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range num {
		if n%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
