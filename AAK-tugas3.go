package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 3000000
	fmt.Println("Quick sort\nJumlah data : ", n)
	for i := 0; i < 5; i++ {
		slice := generateSlice(n)
		// fmt.Println("\n--- Unsorted --- \n\n", slice)
		start := time.Now()
		quicksort(slice)
		elapsed := time.Since(start).Seconds()
		// fmt.Println("\n--- Sorted ---\n\n", slice)
		fmt.Println("Time Complexity : ", elapsed)
	}

	fmt.Println("\nSelection sort\nJumlah data : ", n)
	for i := 0; i < 5; i++ {
		slice := generateSlice(n)
		// fmt.Println("\n--- Unsorted --- \n\n", slice)
		start := time.Now()
		selectionsort(slice)
		elapsed := time.Since(start).Seconds()
		// fmt.Println("\n--- Sorted ---\n\n", slice)
		fmt.Println("Time Complexity : ", elapsed)
	}
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
