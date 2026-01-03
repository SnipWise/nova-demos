package diagram

import "strings"

// ExtractCodeBlock extracts diagram from markdown code block
func ExtractCodeBlock(text string) string {
	lines := strings.Split(text, "\n")
	var diagramLines []string
	inCodeBlock := false

	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inCodeBlock = !inCodeBlock
			continue
		}
		if inCodeBlock {
			diagramLines = append(diagramLines, line)
		}
	}

	return strings.Join(diagramLines, "\n")
}
