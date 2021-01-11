package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio."

	fmt.Println("Hitung Karakter A-Z")

	for char := 'a'; char <= 'z'; char++ {
		fmt.Print(string(char), "=", strings.Count(strings.ToLower(str), string(char)), ", ")
	}

	fmt.Println("\n\nGeser 5 karakter")

	for _, char := range strings.ToLower(str) {
		if int(char) >= 97 && int(char) <= 122 {
			fmt.Print(string(char + 5))
		} else {
			fmt.Print(string(char))
		}
	}

}
