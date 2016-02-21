package main

import (
	"fmt"
)

func MergeSort(unsorted []int, sortType int) []int {
	splits := split(unsorted)
	for len(splits) > 1 {
		fmt.Printf("%v\n", splits)
		var mergedList [][]int
		for index, element := range splits {
			if index%2 == 0 {
				if index+1 < len(splits) {
					ar1 := element
					ar2 := splits[index+1]

					merged := compareAndMerge(ar1, ar2, sortType)
					mergedList = append(mergedList, merged)
				} else {
					mergedList = append(mergedList, element)
				}
			}
		}
		splits = mergedList
	}
	return splits[0]
}

func compareAndMerge(first []int, second []int, sortType int) []int {
	result := make([]int, len(first)+len(second))
	firstIndex, secondIndex := 0, 0
	flag := 0

	for resultIndex := 0; resultIndex < len(result); resultIndex++ {
		if firstIndex >= len(first) {
			flag = 2
		} else if secondIndex >= len(second) {
			flag = 1
		} else {
			comparison := getComparison(first[firstIndex], second[secondIndex], sortType)
			if comparison {
				flag = 1
			} else {
				flag = 2
			}
		}
		switch flag {
		case 1:
			result[resultIndex] = first[firstIndex]
			firstIndex++
			break
		case 2:
			result[resultIndex] = second[secondIndex]
			secondIndex++
			break
		default:
			panic("Unexpected merge condition")
		}
	}

	return result
}

func getComparison(first int, second int, sortType int) bool {
	if sortType == ASC {
		if first < second {
			return true
		} else {
			return false
		}
	} else {
		if first > second {
			return true
		} else {
			return false
		}
	}
}

func split(intArray []int) [][]int {
	result := make([][]int, len(intArray))
	for index, element := range intArray {
		resultItem := []int{element}
		result[index] = resultItem
	}
	return result
}
