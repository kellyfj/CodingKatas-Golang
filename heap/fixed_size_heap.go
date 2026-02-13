package heap

import (
	"container/heap"
)

// FixedSizeHeap is a min-heap that maintains only the smallest 'size' elements added to it.
//
// HEAP STRUCTURE:
// The heap uses an array (data) to store elements in a binary tree structure.
// For an element at index i:
//   - Left child is at index 2*i + 1
//   - Right child is at index 2*i + 2
//   - Parent is at index (i-1)/2
//
// HEAP PROPERTY (Min-Heap):
// Every parent node is smaller than its children. The array is NOT sorted.
// Example: for data = [5, 10, 15, 20, 25], the structure is:
//
//	     5 (root, index 0)
//	    / \
//	  10   15
//	 / \
//	20 25
//
// OPERATIONS:
// - Push: Appends to end, then "bubbles up" to maintain heap property
// - Pop: Removes smallest (root), moves last element to root, then "bubbles down"
// - The container/heap package handles the bubbling automatically via Less() and Swap()
//
// SIZE LIMITATION:
// When the heap exceeds 'size', the smallest element is removed, keeping only the
// 'size' largest elements. This makes it memory-efficient for top-k problems.
type FixedSizeHeap struct {
	size int
	data []int
}

func NewFixedSizeHeap(size int) *FixedSizeHeap {
	return &FixedSizeHeap{
		size: size,
		data: []int{},
	}
}

func (h *FixedSizeHeap) Add(value int) {
	heap.Push(h, value)
}

func (h *FixedSizeHeap) Len() int {
	return len(h.data)
}

func (h *FixedSizeHeap) Less(i, j int) bool {
	// Less defines the heap property: returns true if data[i] < data[j]
	// For a min-heap, this keeps smaller values at the top
	// The heap package uses this to decide how to bubble elements up/down
	return h.data[i] < h.data[j]
}

func (h *FixedSizeHeap) Swap(i, j int) {
	// Swap exchanges two elements in the array
	// Called by the heap package during bubble-up and bubble-down operations
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *FixedSizeHeap) Push(x any) {
	// Push appends the new value to the end of the array
	// The heap package then calls Swap() and Less() to bubble it up to its correct position
	h.data = append(h.data, x.(int))
	// If we exceed the fixed size, remove the smallest element (which bubbles to the top)
	// This keeps only the 'size' largest elements in memory
	if len(h.data) > h.size {
		heap.Pop(h)
	}
}

func (h *FixedSizeHeap) Pop() any {
	// Pop removes and returns the smallest element (the root)
	// The heap package moves the last element to the root and calls Swap() and Less()
	// to bubble it down to its correct position
	old := h.data
	n := len(old)
	x := old[n-1]         // Get the last element
	h.data = old[0 : n-1] // Remove it from the array
	return x
}

// HeapItem represents an element from one of the sorted arrays being merged.
// It tracks the value and its position (which array and which index within that array).
type HeapItem struct {
	value        int
	arrayIndex   int
	elementIndex int
}

// MinHeap is a min-heap implementation for merging sorted arrays.
// It implements the heap.Interface required by the container/heap package.
type MinHeap []HeapItem

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(HeapItem))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// MergeSortedArrays merges multiple sorted arrays into a single sorted array.
//
// TIME COMPLEXITY: O(N log k) where N is the total number of elements and k is the number of arrays.
// This is more efficient than the two-pointer approach for k > 2 arrays.
//
// ALGORITHM:
// 1. Create a min-heap and insert the first element of each sorted array
// 2. Pop the smallest element from the heap and add it to the result
// 3. If the source array has more elements, push the next element into the heap
// 4. Repeat until the heap is empty
//
// EXAMPLE: MergeSortedArrays([][]int{{1,3},{2},{4,5}}) returns [1,2,3,4,5]
func MergeSortedArrays(arrays [][]int) []int {
	h := &MinHeap{}
	heap.Init(h)

	// Push the first element of each array into the heap
	for i, arr := range arrays {
		if len(arr) > 0 {
			heap.Push(h, HeapItem{value: arr[0], arrayIndex: i, elementIndex: 0})
		}
	}

	result := make([]int, 0)
	for h.Len() > 0 {
		item := heap.Pop(h).(HeapItem)
		result = append(result, item.value)

		// If there are more elements in the same array, push the next one into the heap
		if item.elementIndex+1 < len(arrays[item.arrayIndex]) {
			nextValue := arrays[item.arrayIndex][item.elementIndex+1]
			heap.Push(h, HeapItem{value: nextValue, arrayIndex: item.arrayIndex, elementIndex: item.elementIndex + 1})
		}
	}

	return result
}

/**
 * EPIJ 10.3: Sort an almost-sorted array where each number is at most k away
 * from its correctly sorted position. To solve this problem in a general
 * setting we need to be able to store k+1 numbers and want to be able to
 * efficiently extract the min number and add a new number. Use a min-heap. We
 * add the first k numbers to a min-heap. Now we add additional numbers and
 * extract the minimum from the heap. When the numbers run our we drain the heap.
 * Time Complexity: O(n log k)
 * Space Complexity: O(k)
 */
// EXAMPLE: SortKSortedArray([]int{3,2,1,5,4}, 2) returns [1,2,3,4,5]
func SortKSortedArray(arr []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)

	result := make([]int, 0, len(arr))

	// Add the first k+1 elements to the heap
	for i := 0; i <= k && i < len(arr); i++ {
		heap.Push(h, HeapItem{value: arr[i]})
	}

	for i := k + 1; i < len(arr); i++ {
		minItem := heap.Pop(h).(HeapItem)
		result = append(result, minItem.value)
		heap.Push(h, HeapItem{value: arr[i]})
	}

	// Drain the remaining elements from the heap
	for h.Len() > 0 {
		minItem := heap.Pop(h).(HeapItem)
		result = append(result, minItem.value)
	}

	return result
}

// MaxHeap is a max-heap implementation for finding the k closest stars.
// It implements the heap.Interface required by the container/heap package.
type MaxHeap []Star

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// For a max-heap, we want the farther star to be "less" so it bubbles to the top
	return h[i].Distance() > h[j].Distance()
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Star))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	star := old[n-1]
	*h = old[0 : n-1]
	return star
}

type Star struct {
	X, Y, Z float64
}

func (s Star) Distance() float64 {
	return s.X*s.X + s.Y*s.Y + s.Z*s.Z
}

/**
 * EPIJ 10.4: Compute k Closest stars. Use a Max-heap - always removing the max
 * as we go.
 * Time Complexity: O(n log k)
 * Space Complexity: O(k)
 */
func KClosestStars(stars []Star, k int) []Star {
	h := &MaxHeap{}
	heap.Init(h)

	for _, star := range stars {
		heap.Push(h, star)
		if h.Len() > k {
			heap.Pop(h) // Remove the farthest star
		}
	}

	result := make([]Star, h.Len())
	for i := 0; i < len(result); i++ {
		result[i] = heap.Pop(h).(Star)
	}

	// Reverse the result since heap.Pop returns elements in reverse order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

/**
 * EPIJ 10.5: Compute the Median of online data.
 * The median of collection divides the collection into two equal parts. When a new element is
 * added to the collection the parts can change by at most one element and the element to be moved
 * is the largest of the smaller half or the smallest of the larger half.
 * We can use two heaps, a max-heap for the smaller half and a min-heap for the larger half. We will keep
 * the heaps balanced in size.
 * Time Complexity: O(log n) per entry corresponding to insertion and extraction from heap
 */

// MaxHeapItem is a max-heap implementation for HeapItem (used in online median).
// It implements the heap.Interface required by the container/heap package.
type MaxHeapItem []HeapItem

func (h MaxHeapItem) Len() int {
	return len(h)
}

func (h MaxHeapItem) Less(i, j int) bool {
	// For a max-heap, reverse the comparison
	return h[i].value > h[j].value
}

func (h MaxHeapItem) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeapItem) Push(x any) {
	*h = append(*h, x.(HeapItem))
}

func (h *MaxHeapItem) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func OnlineMedian(stream []int) []float64 {
	smallerHalf := &MaxHeapItem{}
	largerHalf := &MinHeap{}
	heap.Init(smallerHalf)
	heap.Init(largerHalf)

	medians := make([]float64, len(stream))

	for i, num := range stream {
		if smallerHalf.Len() == 0 || num <= (*smallerHalf)[0].value {
			heap.Push(smallerHalf, HeapItem{value: num})
		} else {
			heap.Push(largerHalf, HeapItem{value: num})
		}

		// Balance the heaps
		if smallerHalf.Len() > largerHalf.Len()+1 {
			item := heap.Pop(smallerHalf).(HeapItem)
			heap.Push(largerHalf, item)
		} else if largerHalf.Len() > smallerHalf.Len() {
			item := heap.Pop(largerHalf).(HeapItem)
			heap.Push(smallerHalf, item)
		}

		// Calculate median
		if smallerHalf.Len() > largerHalf.Len() {
			medians[i] = float64((*smallerHalf)[0].value)
		} else {
			medians[i] = float64((*smallerHalf)[0].value+(*largerHalf)[0].value) / 2.0
		}
	}

	return medians
}

/**
 * Finds the N smallest but it's very space inefficient.
 * Need a FixedSizePriorityQueue implementation that only keeps "n" objects.
 * How can I fix this it is returning N Largest instead of N Smallest? I think I need to use a MaxHeap instead of a MinHeap and pop the max when I exceed the size.
 * Time Complexity: O(n log n) due to sorting the entire array, which is inefficient for large n.
 * Space Complexity: O(n) due to storing the entire array in the heap, which is inefficient for large n.
 */
func NSmallest(arr []int, n int) []int {
	h := &FixedSizeHeap{size: n}
	heap.Init(h)

	for _, num := range arr {
		h.Add(num)
	}

	result := make([]int, h.Len())
	for i := 0; i < len(result); i++ {
		result[i] = heap.Pop(h).(int)
	}

	return result
}
