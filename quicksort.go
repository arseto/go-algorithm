package main

type splitArray struct {
	less       []int
	greater    []int
	equal      []int
	lessIdx    int
	greaterIdx int
	equalIdx   int
}

func QuickSort(unsorted []int, sortType int) []int {
	if len(unsorted) <= 1 {
		return unsorted
	}
	pivot := getPivot(unsorted)
	split := generateSplitArray(unsorted, pivot)

	if sortType == ASC {
		result := quickSortAsc(split)
		return result
	}

	result := quickSortDesc(split)
	return result

}

func quickSortDesc(split splitArray) []int {
	firstPart := Append(QuickSort(split.greater[0:split.greaterIdx], DESC), split.equal[0:split.equalIdx]...)
	result := Append(firstPart, QuickSort(split.less[0:split.lessIdx], DESC)...)
	return result
}

func quickSortAsc(split splitArray) []int {
	firstPart := Append(QuickSort(split.less[0:split.lessIdx], ASC), split.equal[0:split.equalIdx]...)
	result := Append(firstPart, QuickSort(split.greater[0:split.greaterIdx], ASC)...)
	return result
}

func generateSplitArray(unsorted []int, pivot int) (split splitArray) {

	split.less = make([]int, len(unsorted))
	split.greater = make([]int, len(unsorted))
	split.equal = make([]int, len(unsorted))

	split.lessIdx = 0
	split.greaterIdx = 0
	split.equalIdx = 0

	for i := 0; i < len(unsorted); i++ {
		x := unsorted[i]
		if x < pivot {
			split.less[split.lessIdx] = x
			split.lessIdx++
		} else if x > pivot {
			split.greater[split.greaterIdx] = x
			split.greaterIdx++
		} else {
			split.equal[split.equalIdx] = x
			split.equalIdx++
		}
	}
	return
}

func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

func getPivot(array []int) int {
	return array[len(array)/2]
}
