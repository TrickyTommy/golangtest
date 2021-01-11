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
		fmt.Println("I guess", guess)
		fmt.Println("Is that:")

		if guess == hasil {
			fmt.Println("It took me", tries, "tries to guess correctly.")
			break
		} else {
			fmt.Println("You entered,", guess, "Please try again")
		}
		tries++
		if high <= low {
			fmt.Println("You're confusing me, I give up.")
			break
		}
		guess = randomI(low, high)
	}

}
