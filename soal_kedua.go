package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3}
	// fmt.Println(len(arr))
	// var arri []int

	limit := 11
	total := 0
	for i := 0; i < len(arr); i += limit {
		batch := arr[i:min(i+limit, len(arr))]
		sort.Sort(sort.Reverse(sort.IntSlice(batch)))
		// for j := 0; j < limit; j++ {

		// 	// adding the values of
		// 	// array to the variable sum
		// 	sum = sum + (batch[j])
		// }

		// // declaring a variable
		// // avg to find the average
		// avg := (float64(sum)) / (float64(limit))
		for _, number := range batch {
			total = total + number
		}
		average := total / len(batch) // len  function return array size
		min, max := findMinAndMax(batch)

		fmt.Println(batch)
		fmt.Println("\nAverage = ", average)
		fmt.Println("Min: ", min)
		fmt.Println("Max: ", max)

	}

}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
