package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	input := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	bubbleSort(input)
	for i := range input {
		if input[i] != expected[i] {
			t.Errorf("bubbleSort failed at index %d: got %d, expected %d", i, input[i], expected[i])
		}
	}
}

func TestSelectionSort(t *testing.T) {
	input := []int{64, 25, 12, 22, 11}
	expected := []int{11, 12, 22, 25, 64}
	selectionSort(input)
	for i := range input {
		if input[i] != expected[i] {
			t.Errorf("selectionSort failed at index %d: got %d, expected %d", i, input[i], expected[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	input := []int{38, 27, 43, 3, 9, 82, 10}
	expected := []int{3, 9, 10, 27, 38, 43, 82}
	mergeSort(input)
	for i := range input {
		if input[i] != expected[i] {
			t.Errorf("mergeSort failed at index %d: got %d, expected %d", i, input[i], expected[i])
		}
	}
}

func TestQuickSort(t *testing.T) {
	input := []int{10, 7, 8, 9, 1, 5}
	expected := []int{1, 5, 7, 8, 9, 10}
	quickSort(input)
	for i := range input {
		if input[i] != expected[i] {
			t.Errorf("quickSort failed at index %d: got %d, expected %d", i, input[i], expected[i])
		}
	}
}
