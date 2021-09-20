package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0
	n := []int{10000, 50000, 100000, 500000, 1000000, 5000000}

	// algoritma 1
	duration_1 := []float64{}
	for _, n := range n {
		start := time.Now()
		for i := 1; i <= n; i += 2 {
			for j := 1; j <= n/2; j++ {
				a += 1
			}
		}
		elapsed := time.Since(start).Seconds()
		fmt.Println("Value a:", a)
		fmt.Println("n:", n)
		fmt.Println("Execution time:", elapsed)
		duration_1 = append(duration_1, elapsed)
	}
	fmt.Println(duration_1)

	// algoritma 2
	duration_2 := []float64{}
	for _, n := range n {
		start := time.Now()
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j = j * 2 {
				a += 1
			}
		}
		elapsed := time.Since(start).Seconds()
		fmt.Println("Value a:", a)
		fmt.Println("n:", n)
		fmt.Println("Execution time:", elapsed)
		duration_2 = append(duration_2, elapsed)
	}
	fmt.Println(duration_2)
}
