package main

func GetFibonacciSeriesFast(length int) []int {
	result := make([]int, length)

	for i := 0; i < length; i++ {
		if i == 0 {
			result[i] = 0
		} else if i == 1 {
			result[i] = 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}

	return result
}

func getFibonacciSeries(startIndex int, endIndex int) []int {
	length := endIndex - startIndex
	result := make([]int, length)

	resultIdx := 0
	for i := startIndex; i < endIndex; i++ {
		result[resultIdx] = getFibonacciItem(i)
		resultIdx++
	}

	return result
}

func getFibonacciItem(index int) int {
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return getFibonacciItem(index-2) + getFibonacciItem(index-1)
}
