package validator

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// runeLen returns the number of visual characters (runes) in a string
func runeLen(s string) int {
	return utf8.RuneCountInString(s)
}

// ValidateAlignment checks if all box lines have consistent visual length (runes)
func ValidateAlignment(diagram string) (bool, map[int][]int) {
	lines := strings.Split(diagram, "\n")
	issues := make(map[int][]int) // box number -> line lengths

	boxNum := 0
	var currentBoxLengths []int
	inBox := false

	for _, line := range lines {
		if strings.Contains(line, "┌") {
			boxNum++
			inBox = true
			currentBoxLengths = []int{runeLen(line)}
		} else if strings.Contains(line, "└") {
			currentBoxLengths = append(currentBoxLengths, runeLen(line))
			inBox = false

			// Check if all lengths are the same
			if len(currentBoxLengths) > 0 {
				firstLen := currentBoxLengths[0]
				allSame := true
				for _, l := range currentBoxLengths {
					if l != firstLen {
						allSame = false
						break
					}
				}
				if !allSame {
					issues[boxNum] = currentBoxLengths
				}
			}
			currentBoxLengths = []int{}
		} else if inBox {
			currentBoxLengths = append(currentBoxLengths, runeLen(line))
		}
	}

	return len(issues) == 0, issues
}

// PrintDiagramAnalysis displays detailed line-by-line analysis
func PrintDiagramAnalysis(diagram string, title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 60))

	lines := strings.Split(diagram, "\n")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("Line %2d: %2d chars (visual) | %s\n", i+1, runeLen(line), line)
		}
	}
}
