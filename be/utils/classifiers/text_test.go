package classifiers

import (
	"testing"
)

func TestExtractTextContext(t *testing.T) {
	const filename = "../../assets/txt/1752698096945427000-test.txt"

	expected := map[string]string{
		"3":  "Beautiful is better than ugly.",
		"4":  "Explicit is better than implicit.",
		"5":  "Simple is better than complex.",
		"6":  "Complex is better than complicated.",
		"7":  "Flat is better than nested.",
		"8":  "Sparse is better than dense.",
		"17": "Now is better than never.",
		"18": "Although never is often better than *right* now.",
	}

	results := ExtractTextContext(filename, "better")

	for number, line := range expected {
		if line != results[number] {
			t.Errorf("Expected Context %q doesn't match with actual on line %s -vs- %s", line, line, results[number])
		}
	}

}
