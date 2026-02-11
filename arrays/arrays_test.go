package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinMax(t *testing.T) {
	arr := []int{3, 5, 1, 8, 2, 7}
	min, max := getMinMax(arr)
	assert.Equal(t, 1, min)
	assert.Equal(t, 8, max)

	arr = []int{10}
	min, max = getMinMax(arr)
	assert.Equal(t, 10, min)
	assert.Equal(t, 10, max)

	arr = []int{}
	min, max = getMinMax(arr)
	assert.Equal(t, 0, min)
	assert.Equal(t, 0, max)

	arr = []int{-3, -1, -7, -4}
	min, max = getMinMax(arr)
	assert.Equal(t, -7, min)
	assert.Equal(t, -1, max)
}

func TestDoesSumOfTwoEqualN_orderNlogN(t *testing.T) {
	arr := []int{10, 15, 3, 7}
	target := 17
	assert.True(t, doesSumOfTwoEqualN_orderNlogN(arr, target))

	arr = []int{1, 2, 3, 4, 5}
	target = 10
	assert.False(t, doesSumOfTwoEqualN_orderNlogN(arr, target))

	arr = []int{-1, 0, 1, 2}
	target = 1
	assert.True(t, doesSumOfTwoEqualN_orderNlogN(arr, target))

	arr = []int{}
	target = 5
	assert.False(t, doesSumOfTwoEqualN_orderNlogN(arr, target))
}

func TestDoesSumOfTwoEqualN_orderN(t *testing.T) {
	arr := []int{10, 15, 3, 7}
	target := 17
	assert.True(t, doesSumOfTwoEqualN_orderNOnly(arr, target))

	arr = []int{1, 2, 3, 4, 5}
	target = 10
	assert.False(t, doesSumOfTwoEqualN_orderNOnly(arr, target))

	arr = []int{-1, 0, 1, 2}
	target = 1
	assert.True(t, doesSumOfTwoEqualN_orderNOnly(arr, target))

	arr = []int{}
	target = 5
	assert.False(t, doesSumOfTwoEqualN_orderNOnly(arr, target))
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	original := make([]int, len(arr))
	copy(original, arr)

	shuffle(arr)

	// Check that all original elements are still present
	elementCount := make(map[int]int)
	for _, num := range original {
		elementCount[num]++
	}
	for _, num := range arr {
		elementCount[num]--
	}

	for _, count := range elementCount {
		assert.Equal(t, 0, count)
	}

	// There's a small chance the shuffled array is the same as the original
	// but it's very unlikely for larger arrays. Here we just check that at least
	// one element has changed position.
	samePositionCount := 0
	for i := range arr {
		if arr[i] == original[i] {
			samePositionCount++
		}
	}
	assert.Less(t, samePositionCount, len(arr))
}

func TestCountingSort(t *testing.T) {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	sorted := countingSort(arr, 8)
	expected := []int{1, 2, 2, 3, 3, 4, 8}
	assert.Equal(t, expected, sorted)

	arr = []int{}
	sorted = countingSort(arr, 0)
	expected = []int{}
	assert.Equal(t, expected, sorted)
	arr = []int{5, 5, 5, 5}
	sorted = countingSort(arr, 5)
	expected = []int{5, 5, 5, 5}
	assert.Equal(t, expected, sorted)
}

func TestFindFirstIntegerOccurrence(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	index := findFirstOccurrenceOfItemInSortedArray(arr, target)
	assert.Equal(t, 2, index)

	target = 6
	index = findFirstOccurrenceOfItemInSortedArray(arr, target)
	assert.Equal(t, -1, index)

	arr = []int{1, 1, 1, 1}
	target = 1
	index = findFirstOccurrenceOfItemInSortedArray(arr, target)
	assert.Equal(t, 0, index)

	arr = []int{}
	target = 1
	index = findFirstOccurrenceOfItemInSortedArray(arr, target)
	assert.Equal(t, -1, index)
}

func TestDutchNationalFlagV1(t *testing.T) {
	arr := []int{2, 0, 1, 2, 0, 1, 1, 0}
	res := dutchNationalFlagV1(arr)
	expected := []int{0, 0, 0, 1, 1, 1, 2, 2}
	assert.Equal(t, expected, res)

	arr = []int{0, 0, 0}
	res = dutchNationalFlagV1(arr)
	expected = []int{0, 0, 0}
	assert.Equal(t, expected, res)

	arr = []int{2, 2, 2}
	res = dutchNationalFlagV1(arr)
	expected = []int{2, 2, 2}
	assert.Equal(t, expected, res)

	arr = []int{}
	res = dutchNationalFlagV1(arr)
	expected = []int{}
	assert.Equal(t, expected, res)
}

func TestDutchNationalFlagV2(t *testing.T) {
	arr := []int{2, 0, 1, 2, 0, 1, 1, 0}
	dutchNationalFlagV2(arr)
	expected := []int{0, 0, 0, 1, 1, 1, 2, 2}
	assert.Equal(t, expected, arr)

	arr = []int{0, 0, 0}
	dutchNationalFlagV2(arr)
	expected = []int{0, 0, 0}
	assert.Equal(t, expected, arr)

	arr = []int{2, 2, 2}
	dutchNationalFlagV2(arr)
	expected = []int{2, 2, 2}
	assert.Equal(t, expected, arr)

	arr = []int{}
	dutchNationalFlagV2(arr)
	expected = []int{}
	assert.Equal(t, expected, arr)
}

func TestSearchSmallestInCyclicallySortedArray(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	min := searchSmallestInCyclicallySortedArray(arr)
	assert.Equal(t, 0, min)

	arr = []int{1, 2, 3, 4, 5}
	min = searchSmallestInCyclicallySortedArray(arr)
	assert.Equal(t, 1, min)

	arr = []int{2, 3, 4, 5, 1}
	min = searchSmallestInCyclicallySortedArray(arr)
	assert.Equal(t, 1, min)

	arr = []int{1}
	min = searchSmallestInCyclicallySortedArray(arr)
	assert.Equal(t, 1, min)

	arr = []int{}
	min = searchSmallestInCyclicallySortedArray(arr)
	assert.Equal(t, -1, min)
}
