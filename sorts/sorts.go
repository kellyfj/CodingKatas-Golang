package sorts

/**
 *  |-----------|----------|----------|----------|---------|----------|
 *  | SORT      | BEST     | AVERAGE  | WORST    | SPACE   | IN PLACE |
 *  |-----------|----------|----------|----------|---------|----------|
 *  | Bubble    | O(n)     | O(n^2)   | O(n^2)   | O(1)    |   YES    |
 *  | Selection | O(n^2)   | O(n^2)   | O(n^2)   | O(1)    |   YES    |
 *  | Quick     | O(nlogn) | O(nlogn) | O(n^2)   | O(logn) |   YES    |
 *  | Merge     | O(nlogn) | O(nlogn) | O(nlogn) | Depends |   NO     |
 *  | Heap      | O(nlogn) | O(nlogn) | O(nlogn) | O(1)    |   YES    |
 *  |-----------|----------|----------|----------|---------|----------|
 *
 * @author kellyfj
 */
func bubbleSort(input []int) {
	n := len(input)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
}

func selectionSort(input []int) {
	n := len(input)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if input[j] < input[minIdx] {
				minIdx = j
			}
		}
		input[i], input[minIdx] = input[minIdx], input[i]
	}
}

/**
 * Largely lifted from
 * http://www.vogella.com/articles/JavaAlgorithmsMergesort/article.html
 *
 * @param input
 */
func mergeSort(input []int) {
	if len(input) < 2 {
		return
	}

	mid := len(input) / 2
	left := make([]int, mid)
	right := make([]int, len(input)-mid)

	copy(left, input[:mid])
	copy(right, input[mid:])

	mergeSort(left)
	mergeSort(right)

	merge(input, left, right)
}

func merge(input, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			input[k] = left[i]
			i++
		} else {
			input[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		input[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		input[k] = right[j]
		j++
		k++
	}
}

// public static void quickSort(int[] input) {

func quickSort(input []int) {
	if len(input) < 2 {
		return
	}

	pivot := input[len(input)/2]
	left := []int{}
	right := []int{}
	for _, v := range input {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(input, append(append(left, pivot), right...))
}
