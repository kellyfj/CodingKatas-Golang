package strings

import "strings"

func reverse(in string) string {
	return reverseRunes([]rune(in))
}

func reverseInPlace(r []rune) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func reverseRunes(r []rune) string {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(in string) bool {
	r := []rune(in)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func countVowels(in string) int {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	count := 0
	for _, char := range in {
		if vowels[char] {
			count++
		}
	}
	return count
}

func reverseAllWords(in string) string {
	// NOTE: in the case where there is a trailing space `words`
	// will have a space at the last element! Weird!
	words := strings.Split(in, " ")
	newWordsArray := make([]string, len(words))

	for i, word := range words {
		newWordsArray[(len(words)-1)-i] = word
	}
	return strings.Join(newWordsArray, " ")
}

// Reverse without using Split
func reverseWordsWithoutSplit(in string) string {
	r := []rune(in)
	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			// reverse word from start to i-1
			for j, k := start, i-1; j < k; j, k = j+1, k-1 {
				r[j], r[k] = r[k], r[j]
			}
			start = i + 1
		}
	}
	// reverse the entire string
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Reverse all words Order(1) space

// func isMatch(regex string, s string)

//  Compute Levenshtein Distance

// Get all permutations of a string

// Get the powerset of characters

// Remove duplicate characters

// Find longest palindrome

//Longest common substring

//Add number strings (with converting type)

//Compute all mnemonics for a phone number
