package ascii

import (
	"fmt"
	"os"
	"strings"
)

// PrintBanner generates a string representation of the given word using the banner from the specified file.
// It returns the generated banner as a string.
func PrintBanner(word, filename string) string {
	bannerOutPut := make([][]string, 8)
	FileCheck(filename)
	banner := LoadBanner(filename)
	for _, char := range word {
		if char < 32 || char > 126 {
			fmt.Printf("Character out of range:%q\n", char)
			os.Exit(1)
		}
		if ascii, Ok := banner[char]; Ok {
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				bannerOutPut[i] = append(bannerOutPut[i], asciiLines[i])
			}
		} else {
			fmt.Printf("Character not found: %q\n", char)
			continue
		}
	}
	// Create a StringBuilder to efficiently build the output string.
	var builder strings.Builder
	for i, line := range bannerOutPut {
		builder.WriteString(strings.Join(line, ""))
		if i != len(bannerOutPut) {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}
