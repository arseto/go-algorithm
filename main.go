package main

import (
	"fmt"
	"math/rand"
)

const (
	ASC  = 1
	DESC = 2

	BUBBLE = 3
	QUICK  = 4
)

func main() {
	mode := selectMode()
	switch mode {
	case 1:
		showFibonacci()
		break
	case 2:
		showSort(BUBBLE)
		break
	case 3:
		showSort(QUICK)
		break
	default:
		panic("Invalid mode specified")
	}
}

func selectMode() (mode int) {
	fmt.Println("Please select one:")
	fmt.Println("1: Fibonacci")
	fmt.Println("2: Bubble Sort")
	fmt.Println("3: Quick Sort")
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
	fmt.Println("Please select one:")
	fmt.Printf("%d: Ascending\n", ASC)
	fmt.Printf("%d: Descending\n", DESC)
	fmt.Println()

	sortType = inputInt("selection")
	return
}

func showSort(alg int) {
	size := inputInt("size")
	result := generateArray(size)
	fmt.Println()

	sortType := selectSortType()

	fmt.Println()
	fmt.Println("Original Array:")
	fmt.Println(result)
	fmt.Println()

	if alg == QUICK {
		result = QuickSort(result, sortType)
	} else if alg == BUBBLE {
		BubbleSort(result, sortType)
	} else {
		panic("Invalid sort algorithm")
	}

	fmt.Println("Sorted Array:")
	fmt.Println(result)
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
