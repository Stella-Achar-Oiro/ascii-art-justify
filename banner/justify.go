package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// AlignText aligns the text based on the specified alignment type
func AlignText(text, alignment string) string {
	lines := strings.Split(text, "\n")
	width, err := getTerminalWidth()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(1)
	}

	switch alignment {
	case "left":
		return text
	case "center":
		return alignCenter(lines, width)
	case "right":
		return alignRight(lines, width)
	case "justify":
		return alignJustify(lines, width)
	default:
		return text
	}
}

func alignCenter(lines []string, width int) string {
	var result strings.Builder
	for _, line := range lines {
		padding := (width - len(line)) / 2
		if padding < 0 {
			padding = 0
		}
		result.WriteString(strings.Repeat(" ", padding) + line + "\n")
	}
	return result.String()
}

func alignRight(lines []string, width int) string {
	var result strings.Builder
	for _, line := range lines {
		padding := width - len(line)
		if padding < 0 {
			padding = 0
		}
		result.WriteString(strings.Repeat(" ", padding) + line + "\n")
	}
	return result.String()
}

// func alignJustify(lines []string, width int) string {
// 	var result strings.Builder

// 	for _, line := range lines {
// 		words := strings.Fields(line)
// 		if len(words) == 0 {
// 			result.WriteString("\n")
// 			continue
// 		}
// 		if len(words) == 1 {
// 			result.WriteString(words[0] + "\n")
// 			continue
// 		}

// 		// Calculate the number of spaces to distribute
// 		spaceCount := len(words) - 1
// 		lineLength := len(strings.Join(words, ""))
// 		totalSpaces := width - lineLength

// 		// If there are not enough spaces to distribute evenly, distribute one space between each word
// 		if totalSpaces < spaceCount {
// 			result.WriteString(strings.Join(words, " ") + "\n")
// 			continue
// 		}

// 		// Calculate even spaces and extra spaces to distribute
// 		spaces := totalSpaces / spaceCount
// 		extraSpaces := totalSpaces % spaceCount

// 		for i, word := range words {
// 			result.WriteString(word)
// 			if i < spaceCount {
// 				result.WriteString(strings.Repeat(" ", spaces))
// 				if extraSpaces > 0 {
// 					result.WriteString(" ")
// 					extraSpaces--
// 				}
// 			}
// 		}
// 		result.WriteString("\n")
// 	}

// 	return result.String()
// }

func alignJustify(text, banner string) {
	words := strings.Fields(text)
	width := getTerminalWidth()
	totalAsciLength := 0
	for _, word := range words {
		for _, letter := range word {
			totalAsciLength += len(Ascii.GetLine(1+int(letter-' ')*9, banner))
		}
	}
	totalSpaces := width - totalAsciLength
	if len(words) == 1 {
		for i := 0; i < 8; i++ {
			var lineOutput string
			for _, letter := range words[0] {
				line := Ascii.GetLine(1+int(letter-' ')*9+i, banner)
				lineOutput += line
			}
			lineOutput += strings.Repeat(" ", totalSpaces)
			fmt.Println(lineOutput)
		}
		return
	}
	spaceBetween := totalSpaces / (len(words) - 1)
	extraSpace := totalSpaces % (len(words) - 1)
	for i := 0; i < 8; i++ {
		var lineOutput string
		for j, word := range words {
			for _, letter := range word {
				line := Ascii.GetLine(1+int(letter-' ')*9+i, banner)
				lineOutput += line
			}
			if j < len(words)-1 {
				lineOutput += strings.Repeat(" ", spaceBetween)
				if j < extraSpace {
					lineOutput += " "
				}
			}
		}
		fmt.Println(lineOutput)
	}
}

func getTerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	var rows, cols int
	_, err = fmt.Sscan(string(output), &rows, &cols)
	if err != nil {
		return 0, err
	}
	return cols, nil
}
