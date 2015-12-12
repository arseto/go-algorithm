package main

import (
	"fmt"
	"math/rand"
)

func main() {
	showBubbleSort()
}

func showQuickSort() {
	array := generateArray()
	result := QuickSort(array, ASC)
	fmt.Println(result)
}

func generateArray() []int {
	result := make([]int, 50)
	r := rand.New(rand.NewSource(10))
	for i := 0; i < len(result); i++ {
		result[i] = r.Intn(50)
	}
	fmt.Println(result)
	return result
}

func showBubbleSort() {
	result := generateArray()
	BubbleSort(result, ASC)
	fmt.Println(result)
}

func showFibonacci() {
	fmt.Println(GetFibonacciSeriesFast(10))
}
