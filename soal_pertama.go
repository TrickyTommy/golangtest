package main

import "fmt"

func main() {
	for i := 1; i <= 7; i++ {
		fmt.Print(i)
		for j := 6; j >= i; j-- {

			fmt.Print(0)

		}
		fmt.Println("")

	}

}
