package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	//Initialisation
	e := []int{9,8,7,6,5}

	//Execution
	BubbleSort(e)

	//Validation

	assert.NotNil(t, e)
	assert.EqualValues(t, 5, len(e))
	assert.EqualValues(t, e[0], 5)
	assert.EqualValues(t, e[1], 6)
	assert.EqualValues(t, e[2], 7)
	assert.EqualValues(t, e[3], 8)
	assert.EqualValues(t, e[4], 9)
}

func TestBubbleSortBestCase(t *testing.T) {
	//Initialisation
	e := []int{5,6,7,8,9}

	//Execution
	BubbleSort(e)

	//Validation

	assert.NotNil(t, e)
	assert.EqualValues(t, 5, len(e))
	assert.EqualValues(t, e[0], 5)
	assert.EqualValues(t, e[1], 6)
	assert.EqualValues(t, e[2], 7)
	assert.EqualValues(t, e[3], 8)
	assert.EqualValues(t, e[4], 9)
}

func TestBubbleSortNilCase(t *testing.T) {
	//Execution
	BubbleSort(nil)
}

func getElements(n int) []int {
	result := make([]int, n)
	i:=0
	for j:=n-1; j>=0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	e := getElements(5)
	assert.NotNil(t, e)
	assert.EqualValues(t, 5, len(e))
	assert.EqualValues(t, 4, e[0])
	assert.EqualValues(t, 3, e[1])
	assert.EqualValues(t, 2, e[2])
	assert.EqualValues(t, 1, e[3])
	assert.EqualValues(t, 0, e[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i:=0; i<b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort10(b *testing.B) {
	elements := getElements(10)
	for i:=0; i<b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i:=0; i<b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort1000(b *testing.B) {
	elements := getElements(1000)
	for i:=0; i<b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	elements := getElements(100000)
	for i:=0; i<b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elements := getElements(100000)
	for i:=0; i<b.N; i++ {
		Sort(elements)
	}
}