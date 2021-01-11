package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio."
	kalimat = strings.ToLower(kalimat)
	res1 := strings.Count(kalimat, "a")
	res2 := strings.Count(kalimat, "b")
	res3 := strings.Count(kalimat, "z")

	fmt.Println("Karakter A sebanyak: ", res1)
	fmt.Println("Karakter B sebanyak: ", res2)
	fmt.Println("Karakter Z sebanyak: ", res3)

}
