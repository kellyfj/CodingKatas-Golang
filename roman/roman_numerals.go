package roman

/**
 * EPIJ 6.7: Convert Roman numerals to integers
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Algorithm: Iterate through the string, adding the value of each numeral to the total.
	// If a numeral is followed by a larger numeral, subtract its value instead of adding it.
	total := 0
	for i := 0; i < len(s); i++ {
		value := romanMap[s[i]]
		if i+1 < len(s) && value < romanMap[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}

/**
 * EPIJ 6.8: Convert integers to Roman numerals
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
func intToRoman(num int) string {
	// Define the values and symbols for Roman numerals in descending order to facilitate the conversion.
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Algorithm: Iterate through the values and symbols, repeatedly subtracting the value from the number and
	// appending the symbol to the result string until the number is reduced to zero.
	result := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			result += syms[i]
			num -= vals[i]
		}
	}
	return result
}
