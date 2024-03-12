package paneling

import "strings"

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
