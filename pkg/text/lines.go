package text

import "strings"

func Lines(maxLineWidth int, content string) []string {
	if len(content) == 0 {
		return make([]string, 0)
	}

	lines := make([]string, 0)

	for _, block := range strings.Split(content, "\n") {
		lines = append(lines, blockToLines(maxLineWidth, block)...)
	}

	return lines
}

// blockToLines converts content without newlines into lines.
func blockToLines(maxLineWidth int, content string) []string {
	if len(content) == 0 {
		return make([]string, 0)
	}

	splitted := strings.Split(content, " ")

	words := make([]string, 0)

	for _, s := range splitted {
		for len(s) > maxLineWidth {
			words = append(words, s[:maxLineWidth])
			s = s[maxLineWidth:]
		}
		words = append(words, s)
	}

	lines := make([]string, 0)
	currentLine := ""

	for _, word := range words {
		// CurrentLine is empty, just set it to word.
		if currentLine == "" {
			currentLine = word
			continue
		}

		// If word is small enough to fit into the line, add it.
		if len(currentLine)+len(word)+1 <= maxLineWidth {
			currentLine = currentLine + " " + word
			continue
		}

		// Word does not fit into line, so create a new line.
		lines = append(lines, currentLine)
		currentLine = word
	}

	lines = append(lines, currentLine)

	return lines
}
