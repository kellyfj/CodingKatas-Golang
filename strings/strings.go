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

func findAnagrams(s []string) [][]string {
	anagramMap := make(map[string][]string)

	// Algorithm: For each word, sort the characters to get a canonical form.
	// Use this sorted string as a key in a map,
	for _, word := range s {
		lowerWord := strings.ToLower(word)
		sorted := sortString(lowerWord)
		anagramMap[sorted] = append(anagramMap[sorted], word)
	}

	// The values return in the result will be the groups of anagrams.
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		for j := 0; j < len(r)-i-1; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
	return string(r)
}

//  Compute Levenshtein Distance

// Get all permutations of a string

// Get the powerset of characters

// Remove duplicate characters

// Find longest palindrome

// Longest common substring
func longestCommonSubstring(s1, s2 string) string {

	// Algorithm: Create a map of all substrings of s1. Then iterate through all substrings of
	// s2 and check if they are in the map.
	m := make(map[string]bool)
	for i := 0; i < len(s1); i++ {
		for j := i + 1; j <= len(s1); j++ {
			m[s1[i:j]] = true
		}
	}

	longest := ""
	for i := 0; i < len(s2); i++ {
		for j := i + 1; j <= len(s2); j++ {
			substr := s2[i:j]
			if m[substr] && len(substr) > len(longest) {
				longest = substr
			}
		}
	}
	return longest
}

//Add number strings (with converting type)

//Compute all mnemonics for a phone number
