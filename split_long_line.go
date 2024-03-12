package paneling

import "strings"

// SplitLongLine takes a single line of text and a width, and splits the line into multiple lines if it exceeds the given width.
// This function ensures that words are not broken in the middle, if possible, and is used to format content to fit within grid boundaries.
//
// Parameters:
// - line: The line of text to be split.
// - width: The maximum allowed width of each line after splitting. Must be a positive integer.
//
// Returns:
// - []string: A slice of strings, each representing a line of text that does not exceed the specified width.
func SplitLongLine(line string, width int) []string {
	var result []string

	words := strings.Fields(line)
	lineBuf := ""
	for i := 0; i < len(words); i++ {
		if len(lineBuf)+len(words[i]) <= width { // if the word fits in the line
			lineBuf += words[i] + " "
			continue
		}

		// whenever the line is full
		result = append(result, strings.TrimSpace(lineBuf))
		lineBuf = ""

		// The word didn't fit in the last line
		if len(words[i]) > width { // if the word is longer than an entire line
			result = append(result, words[i][0:width])
		} else {
			lineBuf += words[i] + " "
		}
	}

	if len(lineBuf) > 0 {
		result = append(result, strings.TrimSpace(lineBuf))
	}

	return result
}
