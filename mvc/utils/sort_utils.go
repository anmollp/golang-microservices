package utils

import "sort"

func BubbleSort(elements []int) {
	loopthrough := true
	for loopthrough {
		loopthrough = false
		for i :=0; i< len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				loopthrough = true
			}
		}
	}
}

func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
	} else {
		sort.Ints(elements)
	}
}
