package main

func GetFactorialSeriesFast(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		var last int
		if i == 0 {
			last = 1
		} else {
			last = result[i-1]
		}
		result[i] = last * (i + 1)
	}
	return result
}

func GetFactorialSeriesLooping(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = GetFactorialValue(i + 1)
	}
	return result
}

func GetFactorialValue(x int) int {
	result := 1
	for i := 1; i <= x; i++ {
		result = i * result
	}
	return result
}
