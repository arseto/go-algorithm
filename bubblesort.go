package main

const (
	ASC = iota
	DESC
)

func swap(unsorted []int, i int, j int) {
	temp := unsorted[i]
	unsorted[i] = unsorted[j]
	unsorted[j] = temp
}

func getSwapCondition(sortType int, unsorted []int, index int) (swapCondition bool) {
	if sortType == ASC {
		swapCondition = unsorted[index] < unsorted[index-1]
	} else {
		swapCondition = unsorted[index] > unsorted[index-1]
	}
	return
}

func BubbleSort(unsorted []int, sortType int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(unsorted); i++ {
			swapCondition := getSwapCondition(sortType, unsorted, i)
			if swapCondition {
				swap(unsorted, i, i-1)
				swapped = true
			}
		}

	}
}
