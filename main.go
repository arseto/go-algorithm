package main

import (
	"fmt"
	"math/rand"
)

const (
	ASC  = 1
	DESC = 2

	BUBBLE = 1
	QUICK  = 2
)

func main() {
	mode := selectMode()
	switch mode {
	case 1:
		showFibonacci()
		break
	case 2:
		showSort()
		break
	default:
		panic("Invalid mode specified")
	}
}

func selectMode() (mode int) {
	fmt.Println("Please select mode:")
	fmt.Println("1: Fibonacci")
	fmt.Println("2: Sort")
	fmt.Println()

	mode = inputInt("selection")
	return
}

func inputInt(verb string) int {
	fmt.Printf("Please input %v: ", verb)
	var mode int
	_, err := fmt.Scanf("%d", &mode)

	if err != nil {
		panic(err)
	}
	return mode
}

func selectSortType() (sortType int) {
	fmt.Println("Please select sort direction:")
	fmt.Printf("%d: Ascending\n", ASC)
	fmt.Printf("%d: Descending\n", DESC)
	fmt.Println()

	sortType = inputInt("sort direction")
	return
}

func selectSortAlgorithm() (alg int) {
	fmt.Println("Please select sort algorithm:")
	fmt.Printf("%d: Bubble Sort\n", BUBBLE)
	fmt.Printf("%d: Quick Sort\n", QUICK)
	fmt.Println()

	alg = inputInt("algorithm")
	return
}

func showSort() {
	alg := selectSortAlgorithm()
	sortType := selectSortType()

	size := inputInt("size")
	unsorted := generateArray(size)
	fmt.Println()

	fmt.Println("Original Array:")
	fmt.Println(unsorted)
	fmt.Println()

	sorted := doSort(unsorted, alg, sortType)

	fmt.Println("Sorted Array:")
	fmt.Println(sorted)
}

func doSort(result []int, alg int, sortType int) []int {
	switch alg {
	case QUICK:
		result = QuickSort(result, sortType)
		break
	case BUBBLE:
		BubbleSort(result, sortType)
		break
	default:
		panic("Invalid sort algorithm")
	}
	return result
}

func generateArray(size int) []int {
	result := make([]int, size)
	r := rand.New(rand.NewSource(int64(size)))
	for i := 0; i < len(result); i++ {
		result[i] = r.Intn(size)
	}
	return result
}

func showFibonacci() {
	size := inputInt("size")
	fmt.Println(GetFibonacciSeriesFast(size))
}
