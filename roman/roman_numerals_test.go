package roman

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		result := romanToInt(test.input)
		if result != test.expected {
			t.Errorf("romanToInt(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, test := range tests {
		result := intToRoman(test.input)
		if result != test.expected {
			t.Errorf("intToRoman(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
