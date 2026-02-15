package heap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSingleElement(t *testing.T) {
	h := NewFixedSizeHeap(5)
	h.Add(10)
	assert.Equal(t, 1, h.Len())
}

func TestAddMultipleElements(t *testing.T) {
	h := NewFixedSizeHeap(5)
	h.Add(10)
	h.Add(20)
	h.Add(5)
	h.Add(15)
	h.Add(25)

	assert.Equal(t, 5, h.Len())
	// Just verify all elements are present, not their specific order
	result := make([]int, len(h.data))
	copy(result, h.data)
	sort.Ints(result)
	expected := []int{5, 10, 15, 20, 25}
	assert.Equal(t, expected, result)
}

func TestAddExceedingSize(t *testing.T) {
	h := NewFixedSizeHeap(3)
	h.Add(10)
	h.Add(20)
	h.Add(5)
	h.Add(15) // This should cause the smallest (5) to be removed

	assert.Equal(t, 3, h.Len())
	result := make([]int, len(h.data))
	copy(result, h.data)
	sort.Ints(result)
	expected := []int{10, 15, 20} // 5 should be removed
	assert.Equal(t, expected, result)
}

func TestMergeSortedArraysTwoArrays(t *testing.T) {
	arrays := [][]int{{1, 3}, {2}}
	result := MergeSortedArrays(arrays)
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, result)
}

func TestMergeSortedArraysMultipleArrays(t *testing.T) {
	arrays := [][]int{{1, 3}, {2}, {4, 5}}
	result := MergeSortedArrays(arrays)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestMergeSortedArraysEmpty(t *testing.T) {
	arrays := [][]int{}
	result := MergeSortedArrays(arrays)
	expected := []int{}
	assert.Equal(t, expected, result)
}

func TestMergeSortedArraysSomeEmpty(t *testing.T) {
	arrays := [][]int{{}, {1, 2}, {}, {3}}
	result := MergeSortedArrays(arrays)
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, result)
}

func TestMergeSortedArraysLarge(t *testing.T) {
	arrays := [][]int{
		{1, 5, 9},
		{2, 6, 10},
		{3, 7, 11},
		{4, 8, 12},
	}
	result := MergeSortedArrays(arrays)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	assert.Equal(t, expected, result)
}

func TestSortKSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	k := 2
	result := SortKSortedArray(arr, k)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestSortKSortAlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	k := 2
	result := SortKSortedArray(arr, k)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestSortKSortReverseSorted(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	k := 4
	result := SortKSortedArray(arr, k)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestSortKSortSingleElement(t *testing.T) {
	arr := []int{1}
	k := 0
	result := SortKSortedArray(arr, k)
	expected := []int{1}
	assert.Equal(t, expected, result)
}

func TestSortKSortEmpty(t *testing.T) {
	arr := []int{}
	k := 2
	result := SortKSortedArray(arr, k)
	expected := []int{}
	assert.Equal(t, expected, result)
}

func TestKClosestStars(t *testing.T) {
	stars := []Star{
		{X: 1, Y: 2, Z: 3},
		{X: 4, Y: 5, Z: 6},
		{X: 0, Y: 0, Z: 0},
		{X: 7, Y: 8, Z: 9},
	}
	k := 2
	result := KClosestStars(stars, k)
	expected := []Star{
		{X: 0, Y: 0, Z: 0},
		{X: 1, Y: 2, Z: 3},
	}
	assert.Equal(t, expected, result)
}

func TestOnlineMedian(t *testing.T) {
	values := []int{5, 15, 1, 3}
	medians := []float64{5, 10, 5, 4}

	for i := range values {
		stream := make([]int, len(values[:i+1]))
		copy(stream, values[:i+1])
		result := OnlineMedian(stream)

		expectedMedians := make([]float64, len(medians[:i+1]))
		copy(expectedMedians, medians[:i+1])
		assert.Equal(t, expectedMedians, result)
	}
}

func TestNSmallestElements(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	n := 3
	result := getNSmallestHeap(arr, n)
	expected := []int{3, 4, 7}
	assert.Equal(t, expected, result)
}

func TestNSmallestElementsMoreThanN(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	n := 10
	result := getNSmallestHeap(arr, n)
	expected := []int{3, 4, 7, 10, 15, 20}
	assert.Equal(t, expected, result)
}

func TestNSmallestElementsEmpty(t *testing.T) {
	arr := []int{}
	n := 3
	result := getNSmallestHeap(arr, n)
	expected := []int{}
	assert.Equal(t, expected, result)
}

func TestNSmallestElementsNZero(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	n := 0
	result := getNSmallestHeap(arr, n)
	expected := []int{}
	assert.Equal(t, expected, result)
}

func TestNLargest(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	n := 3
	result := getNLargestHeap(arr, n)
	expected := []int{10, 15, 20}
	assert.Equal(t, expected, result)
}

func TestNLargestMoreThanN(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	n := 10
	result := getNLargestHeap(arr, n)
	expected := []int{3, 4, 7, 10, 15, 20}
	assert.Equal(t, expected, result)
}

func TestNLargestEmpty(t *testing.T) {
	arr := []int{}
	n := 3
	result := getNLargestHeap(arr, n)
	expected := []int{}
	assert.Equal(t, expected, result)
}

func TestNLargestNZero(t *testing.T) {
	arr := []int{7, 10, 4, 3, 20, 15}
	n := 0
	result := getNLargestHeap(arr, n)
	expected := []int{}
	assert.Equal(t, expected, result)
}
