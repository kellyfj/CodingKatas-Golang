package arrays

import "math/rand"

/**
 * GOAL: Get the min and max using 3N/2 comparisons (as opposed to 2N)
 *
 * Initialize min and max as minimum and maximum of the first two
 * elements respectively. For rest of the elements, pick them in pairs and
 * compare their maximum and minimum with max and min respectively.
 * http://www.geeksforgeeks.org/maximum-and-minimum-in-an-array/
 * TIME COMPLEXITY: O(3n/2)
 */
func getMinMax(arr []int) (min int, max int) {
	if len(arr) == 0 {
		return 0, 0
	}

	// Initialize with first element
	min, max = arr[0], arr[0]
	startIndex := 1

	// Handle first pair if array has even length
	if len(arr)%2 == 0 {
		if arr[0] > arr[1] {
			min, max = arr[1], arr[0]
		} else {
			min, max = arr[0], arr[1]
		}
		startIndex = 2
	}

	// Process pairs
	for i := startIndex; i < len(arr)-1; i += 2 {
		// Compare pair and assign to temp variables
		small, large := arr[i], arr[i+1]
		if small > large {
			small, large = large, small
		}
		// Update min and max
		if small < min {
			min = small
		}
		if large > max {
			max = large
		}
	}
	return min, max
}

/**
 * http://codercareer.blogspot.com/2011/10/no-09-numbers-with-given-sum.html
 * Given an array, please determine whether it contains two numbers whose
 * sum equals to a target.
 * TIME COMPLEXITY: O(nlogn) due to sort
 */
func doesSumOfTwoEqualN_orderNlogN(arr []int, target int) bool {
	// Sort the array
	quickSort(arr, 0, len(arr)-1)

	// Use two pointers to find the sum
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// doesSumOfTwoEqualN_orderNOnly - Different variant of Sum-of-Two-equals-N
// TIME COMPLEXITY: O(N) due to two sequential loops
func doesSumOfTwoEqualN_orderNOnly(arr []int, target int) bool {
	numsSeen := make(map[int]bool)
	for _, num := range arr {
		complement := target - num
		if numsSeen[complement] {
			return true
		}
		numsSeen[num] = true
	}
	return false
}

/**
 * Sum-of-three Algorithm
 * TIME COMPLEXITY: O(n)*O(n log n) = O(n^2)
 */
// public static boolean doesSumOfThreeEqualN(int[] array, int target) {

/**
 * http://stackoverflow.com/questions/2070359/finding-three-elements-in-an-array-whose-sum-is-closest-to-an-given-number
 *
 * TIME COMPLEXITY: Still O(n^2) but less comparisons
 */
// public static boolean doesSumOfThreeEqualN_faster(int[] A, int target)

/**
 * Randomly shuffle object in an array
 * @param a
 */
func shuffle(arr []int) {
	// Implementing Fisher-Yates shuffle
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := randInt(0, i) // Random index from 0 to i
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

/**
 * TIME COMPLEXITY: O(n)
 */
// public static int lengthOfLongestMonotonicSeries(int[] input) {

/**
 * EPIJ 5.1: The Dutch National Flag Problem
 *
 * http://en.wikipedia.org/wiki/Dutch_national_flag_problem
 * TIME COMPLEXITY: O(n)
 * SPACE COMPLEXITY: O(n) for the new array
 */
func dutchNationalFlagV1(arr []int) []int {
	result := make([]int, len(arr))
	lowIndex := 0
	highIndex := len(arr) - 1

	for _, value := range arr {
		if value < 1 {
			result[lowIndex] = value
			lowIndex++
		} else if value > 1 {
			result[highIndex] = value
			highIndex--
		}
	}

	// Fill in the mid values
	for i := lowIndex; i <= highIndex; i++ {
		result[i] = 1
	}
	return result
}

/**
 * TIME COMPLEXITY: O(n)
 * SPACE COMPLEXITY: O(1) since we are swapping in place
 */
func dutchNationalFlagV2(arr []int) {
	numZeroes := 0
	numOnes := 0
	numTwos := 0

	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			numZeroes++
		case 1:
			numOnes++
		case 2:
			numTwos++
		}
	}

	for i := 0; i < numZeroes; i++ {
		arr[i] = 0
	}
	for i := numZeroes; i < numZeroes+numOnes; i++ {
		arr[i] = 1
	}
	for i := numZeroes + numOnes; i < numZeroes+numOnes+numTwos; i++ {
		arr[i] = 2
	}
}

/**
 * Counting sort for integers in the range 0 to max_range
 * TIME COMPLEXITY: O(n)
 * SPACE COMPLEXITY: O(n) + O(k) = O(n)
 */
func countingSort(arr []int, maxRange int) []int {
	counts := make([]int, maxRange+1)
	for _, num := range arr {
		counts[num]++
	}

	sortedIndex := 0
	for num, count := range counts {
		for i := 0; i < count; i++ {
			arr[sortedIndex] = num
			sortedIndex++
		}
	}
	return arr
}

/**
 * EPIJ 11.1: Search a sorted array for first occurrence of an integer
 *
 * TIME COMPLEXITY: O(log n)
 * SPACE COMPLEXITY: O(1)
 */
// public static int findFirstOccurrenceOfItemInSortedArray(int[] a, int lookFor)

func findFirstOccurrenceOfItemInSortedArray(arr []int, lookFor int) int {
	low, high := 0, len(arr)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == lookFor {
			result = mid
			high = mid - 1 // Continue searching in the left half
		} else if arr[mid] < lookFor {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

/**
 * EPIJ 11.3: Search a cyclically sorted array.
 * Cyclically sorted means it is possible to cyclically shift the entries so that it becomes sorted.
 * Example: 378 - 478 - 550 - 631 - 103 - 203 - 220 - 234 - 279 - 368
 * Brute force approach is to iterate through the array comparing against running minimum - complexity
 * O(n). But we should take advantage of properties of the array if A[m] >
 * A[n-1] then the min value must be in the range [m+1, n-1]
 * Time Complexity: O(log n)
 */
func searchSmallestInCyclicallySortedArray(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr)-1

	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > arr[high] {
			// Min must be in A.sublist(mid+1, high+1)
			low = mid + 1
		} else {
			// A.get(mid) < A.get(high)
			// Minimum cannot be in A.sublist(mid+1, high+1) so it must be in
			// A.sublist(low, mid+1)
			high = mid
		}
	}
	return arr[low]
}

// public static int countBinaryOnes(int number) {

// public static String getBinaryRepresentation(int number)
func getBinaryRepresentation(number int) string {
	if number == 0 {
		return "0"
	}
	binary := ""
	for number > 0 {
		binary = string(rune('0'+(number%2))) + binary
		number /= 2
	}
	return binary
}

// public static List<Set<Integer>> powerSet(int[] array) {

/**
 * EPIJ 11.4: Calculate the Integer Square Root
 * Time Complexity: O(log k)
 */
// public static int calcSquareRoot(int k) {

// public static double calcSquareRoot(double num)

// public static List<List<Integer>> findCombosToReachTargetSum(List<Integer> numbers, int target) {
func mergeSortedArrays(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		merged[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		merged[k] = b[j]
		j++
		k++
	}
	return merged
}

// MergeSortedArrays merges nums2 into nums1 in-place, assuming nums1 has enough space.
//
// Algorithmic flow:
// 1. Start from the end of both arrays (nums1 and nums2), using three pointers:
// - i: last valid element in nums1 (m-1)
// - j: last element in nums2 (n-1)
// - k: last position in nums1 (m+n-1)
// 2. Compare nums1[i] and nums2[j], and place the larger at nums1[k]. Move the corresponding pointer(s) backward.
// 3. Repeat until either i or j is exhausted.
// 4. If nums2 has remaining elements, copy them into nums1 (nums1's remaining elements are already in place).
func mergeSortedArraysInPlace(nums1 []int, m int, nums2 []int, n int) {
	// Start from the end of both arrays
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// If nums2 is not exhausted, copy remaining elements
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// public static Integer[] mergeRemovingDupes(int[] a, int[] b) {

/**
     * TIME COMPLEXITY: O(m+n)
     * SPACE COMPLEXITY: O(1)
	 	public static int getKthSmallestElementInTwoSortedArrays(int[] array1,
			int[] array2, int k) {
*/

/**
 * TIME COMPLEXITY: O(n)
 * SPACE COMPLEXITY: O(n)
public static boolean containsDuplicates_orderNspace(int[] array) {
*/

/**
 * EPIJ 11.10: We are given a list of n-1 integers in the range 1 to n with no duplicates.
 * Find the missing integers.
 *
 * TIME COMPLEXITY: O(n)
 * SPACE COMPLEXITY: O(1)
 */

/**
 * EPIJ: 5.2 Increment an Arbitrary Precision Integer
 * Brute force version
 * Time Complexity: O(n)
 * Space Complexity: O(n)
 */

/**
 * EPIJ: 5.2 Increment an Arbitrary Precision Integer
 * More efficient version.
 *
 * Time Complexity: TBD
 * Space Complexity: O(1)
 */

/**
 * EPIJ 5.5: Remove duplicates from a Sorted Array
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1) - return the size of the array with duplicates moved to end
 */

/**
 * EPIJ: 5.6 - Buy and Sell a Stock once
 *
 * Time Complexity: O(n^2)
 */

/**
 * Key Insight
 * Maximum profit can be made by selling on the day determined by the minimum of the stock prices over
 * the previous days so we
 * 	1) Iterate through the list
 *  2) Keep track of the minimum element
 *  3) If the difference between the current element and the min element is greater than max profit
 *      then update the maxProfit
 * Time Complexity: O(n)
 */

/**
 * EPIJ 5.9 Enumerate all primes to n
 *
 * Start with 0,1 as NOT prime, 2 is a prime then sieve out all multiples of 2
 * from a boolean list, 3 is a prime etc. continue . . .
 *
 * Time Complexity: O(n log log n)
 * Space Complexity: O(n)
 */

/**
 * EPIJ 5.12: Sample Offline Data
 *
 * The key to efficiently building a random subset of size k is to first build one of size k-1
 * and then add one more element, selected randomly from the set
 * For k > 1 we begin by choosing one element at random and place that at A[0] and repeat
 * the same process again with the n-1 element subarray A[1, n-1]
 * Eventually the random subset occupies that slots A[0, k-1] and the remaining elements are in the
 * last n-k nodes
 */

/**
 * EPIJ 5.13 Sample Online Data
 *
 * Design a program that takes as input a size (k) and reads packets, continuously maintaining
 * a uniform random subset of size (k) of the read packets.
 *
 * Suppose we read the first n packets and have a random subset of k of them.
 * When we read the (n+1)the packet it should belong to the new subset with the probability k/(n+1).
 * If we choose one of the packets in the subset randomly to remove the resulting collection
 * will be a random subset of the n+1 packets.
 */
