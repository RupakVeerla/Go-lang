package main

import "fmt"

func main() {
	fmt.Println("Hello Rupak!")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("Existed.....Bye")
}
