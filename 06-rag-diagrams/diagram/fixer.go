package diagram

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// runeLen returns the number of visual characters (runes) in a string
func runeLen(s string) int {
	return utf8.RuneCountInString(s)
}

// Box represents a single box in the diagram with its position
type Box struct {
	StartLine int
	EndLine   int
	StartCol  int
	Lines     []string
}

// FixAlignment auto-corrects ASCII diagram alignment issues while preserving multi-column layouts
func FixAlignment(diagram string) string {
	lines := strings.Split(diagram, "\n")

	// Detect all boxes in the diagram
	boxes := detectBoxes(lines)

	// Fix each box individually
	for i := range boxes {
		boxes[i].Lines = fixBox(boxes[i].Lines)
	}

	// Reconstruct the diagram preserving positions
	return reconstructDiagram(lines, boxes)
}

// detectBoxes finds all boxes in the diagram and their positions
func detectBoxes(lines []string) []Box {
	var boxes []Box

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Find all ┌ characters in this line (multiple boxes can be on same line)
		boxStarts := findBoxStarts(line)

		for _, startCol := range boxStarts {
			// Extract this box
			box := extractBox(lines, i, startCol)
			if box != nil {
				boxes = append(boxes, *box)
			}
		}
	}

	return boxes
}

// findBoxStarts returns column positions of all ┌ characters in a line
func findBoxStarts(line string) []int {
	var positions []int
	col := 0

	for _, r := range line {
		if r == '┌' {
			positions = append(positions, col)
		}
		col++
	}

	return positions
}

// extractBox extracts a single box starting at the given position
func extractBox(lines []string, startLine, startCol int) *Box {
	if startLine >= len(lines) {
		return nil
	}

	box := &Box{
		StartLine: startLine,
		StartCol:  startCol,
		Lines:     []string{},
	}

	// Extract the box line by line
	for i := startLine; i < len(lines); i++ {
		line := lines[i]

		// Extract the portion of the line that belongs to this box
		boxPart := extractBoxPart(line, startCol)
		if boxPart == "" {
			break
		}

		box.Lines = append(box.Lines, boxPart)

		// Check if this is the last line of the box
		if strings.Contains(boxPart, "└") {
			box.EndLine = i
			break
		}
	}

	if len(box.Lines) == 0 {
		return nil
	}

	return box
}

// extractBoxPart extracts the box portion from a line starting at startCol
func extractBoxPart(line string, startCol int) string {
	runes := []rune(line)

	if startCol >= len(runes) {
		return ""
	}

	// Find where this box ends on this line
	var result []rune
	inBox := false

	for i := startCol; i < len(runes); i++ {
		r := runes[i]

		// Start of box
		if r == '┌' || r == '├' || r == '│' || r == '└' {
			inBox = true
		}

		if inBox {
			result = append(result, r)

			// End of box line
			if r == '┐' || r == '┤' || r == '┘' {
				break
			}

			// If we hit another box or empty space after │, stop
			if i > startCol && r == ' ' && i+1 < len(runes) && runes[i+1] == '┌' {
				break
			}
		}
	}

	return strings.TrimRight(string(result), " ")
}

// reconstructDiagram rebuilds the diagram with fixed boxes at their original positions
func reconstructDiagram(originalLines []string, boxes []Box) string {
	// Create a copy of original lines
	result := make([]string, len(originalLines))
	copy(result, originalLines)

	// Track which lines/columns have been modified by boxes
	modified := make(map[int]map[int]bool) // line -> column -> modified

	// Replace each box in its original position
	for _, box := range boxes {
		for i, boxLine := range box.Lines {
			lineIdx := box.StartLine + i
			if lineIdx >= len(result) {
				continue
			}

			if modified[lineIdx] == nil {
				modified[lineIdx] = make(map[int]bool)
			}

			// Replace the box part in the original line
			result[lineIdx] = replaceBoxPart(result[lineIdx], box.StartCol, boxLine, modified[lineIdx])
		}
	}

	return strings.Join(result, "\n")
}

// replaceBoxPart replaces a portion of a line with the fixed box part
func replaceBoxPart(originalLine string, startCol int, boxPart string, modifiedCols map[int]bool) string {
	runes := []rune(originalLine)
	boxRunes := []rune(boxPart)

	// Ensure the line is long enough
	for len(runes) < startCol+len(boxRunes) {
		runes = append(runes, ' ')
	}

	// Replace the box portion
	for i, r := range boxRunes {
		col := startCol + i
		if !modifiedCols[col] {
			runes[col] = r
			modifiedCols[col] = true
		}
	}

	return string(runes)
}

// fixBox ensures all lines in a box have the same visual length
func fixBox(boxLines []string) []string {
	if len(boxLines) == 0 {
		return boxLines
	}

	// Calculate the required width based on longest content
	maxContentLen := 0
	for _, line := range boxLines {
		// Check content lines (those with │ on both sides)
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "│") && strings.HasSuffix(trimmed, "│") {
			// Extract content between │ symbols
			re := regexp.MustCompile(`│\s*(.+?)\s*│`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				content := strings.TrimSpace(matches[1])
				contentLen := runeLen(content)
				if contentLen > maxContentLen {
					maxContentLen = contentLen
				}
			}
		}
	}

	// Calculate target width: │ + space + content + space + │
	targetWidth := maxContentLen + 4

	// Now fix each line to match this width
	fixed := make([]string, len(boxLines))
	for idx, line := range boxLines {
		fixed[idx] = fixLine(line, targetWidth)
	}

	return fixed
}

// fixLine adjusts a single line to match the target length
func fixLine(line string, targetLen int) string {
	trimmed := strings.TrimRight(line, " ")

	// Border lines (┌, ├, └)
	if strings.Contains(trimmed, "┌") || strings.Contains(trimmed, "├") || strings.Contains(trimmed, "└") {
		return fixBorderLine(trimmed, targetLen)
	}

	// Content lines (│ ... │)
	if strings.HasPrefix(strings.TrimSpace(trimmed), "│") && strings.HasSuffix(trimmed, "│") {
		return fixContentLine(trimmed, targetLen)
	}

	// Other lines (like inheritance symbols)
	return line
}

// fixBorderLine ensures border has correct length
func fixBorderLine(line string, targetLen int) string {
	// Detect the border pattern
	var leftChar, rightChar, fillChar string

	if strings.Contains(line, "┌") {
		leftChar = "┌"
		rightChar = "┐"
		fillChar = "─"
	} else if strings.Contains(line, "├") {
		leftChar = "├"
		rightChar = "┤"
		fillChar = "─"
	} else if strings.Contains(line, "└") {
		leftChar = "└"
		rightChar = "┘"
		fillChar = "─"
	} else {
		return line
	}

	// Calculate how many fill characters we need
	fillCount := targetLen - 2 // -2 for left and right chars
	if fillCount < 0 {
		fillCount = 0
	}

	return leftChar + strings.Repeat(fillChar, fillCount) + rightChar
}

// fixContentLine pads content to reach target length
func fixContentLine(line string, targetLen int) string {
	// Extract content between │ symbols
	re := regexp.MustCompile(`│\s*(.+?)\s*│`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 2 {
		// Fallback: just pad with spaces
		if len(line) < targetLen {
			return line + strings.Repeat(" ", targetLen-len(line))
		}
		return line
	}

	content := strings.TrimSpace(matches[1])

	// Calculate padding needed using RUNE count (visual characters)
	// Format: │ + space + content + spaces + │ = targetLen
	// So: padding = targetLen - 2(for │) - 1(space before) - runeLen(content)
	contentLen := runeLen(content)
	totalPadding := targetLen - contentLen - 3
	if totalPadding < 1 {
		totalPadding = 1 // At least one space before closing │
	}

	// Build the line: │ + space + content + padding + │
	return "│ " + content + strings.Repeat(" ", totalPadding) + "│"
}
