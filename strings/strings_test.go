package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, reverse("hello"), "olleh")
	assert.Equal(t, reverse("Go is awesome!"), "!emosewa si oG")
	assert.Equal(t, reverse("racecar"), "racecar")
	assert.Equal(t, reverse("a"), "a")
	assert.Equal(t, reverse(""), "")
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("racecar"))
	assert.True(t, isPalindrome("madam"))
	assert.False(t, isPalindrome("hello"))
	assert.True(t, isPalindrome("a"))
	assert.True(t, isPalindrome(""))
}

func TestCountVowels(t *testing.T) {
	assert.Equal(t, countVowels("hello"), 2)
	assert.Equal(t, countVowels("Go is awesome!"), 6)
	assert.Equal(t, countVowels("rhythm"), 0)
	assert.Equal(t, countVowels("AEIOU"), 5)
	assert.Equal(t, countVowels("aaeeiioouuAAEEIIOOUU"), 20)
	assert.Equal(t, countVowels(""), 0)
}

func TestReverseInPlace(t *testing.T) {
	r := []rune("hello")
	reverseInPlace(r)
	assert.Equal(t, string(r), "olleh")

	r = []rune("Go is awesome!")
	reverseInPlace(r)
	assert.Equal(t, string(r), "!emosewa si oG")

	r = []rune("racecar")
	reverseInPlace(r)
	assert.Equal(t, string(r), "racecar")

	r = []rune("a")
	reverseInPlace(r)
	assert.Equal(t, string(r), "a")

	r = []rune("")
	reverseInPlace(r)
	assert.Equal(t, string(r), "")
}

func TestReverseWords(t *testing.T) {
	rev := reverseAllWords("able was I ere I saw elba")
	assert.Equal(t, "elba saw I ere I was able", rev)

	assert.Equal(t, "", reverseAllWords(""))
	assert.Equal(t, "one", reverseAllWords("one"))
	assert.Equal(t, "two one", reverseAllWords("one two"))
	assert.Equal(t, " space extra one", reverseAllWords("one extra space "))
}
