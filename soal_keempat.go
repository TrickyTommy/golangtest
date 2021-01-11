package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomI(low int, high int) int {
	return rand.Intn(high-low) + low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Think of a number between 1 and 100.")
	fmt.Print("Press ENTER when ready.")
	scanner.Scan()

	low := 1
	high := 100
	hasil := randomI(low, high)
	guess := randomI(low, high)
	tries := 1
	// reply := guess
	for {
		fmt.Print("Saya tebak ", guess)
		fmt.Println(" apakah benar? ")

		if guess == hasil {
			fmt.Println(" Benar !!!,hasil random", hasil)
			fmt.Println("Itu membutuhkan program ini sebanyak", tries, "sehingga menemukan angka random yang di tentukan")
			break
		} else {
			fmt.Println("hasil tebakan adalah", guess, "Coba Lagi")
		}
		tries++
		if high <= low {
			fmt.Println("You're confusing me, I give up.")
			break
		}
		guess = randomI(low, high)
	}

}
