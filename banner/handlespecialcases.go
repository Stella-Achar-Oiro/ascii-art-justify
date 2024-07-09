package ascii

import (
	"fmt"
	"os"
	"strings"
)

// HandleSpecialCases checks if the input string contains any unsupported special escape sequences.
func HandleSpecialCases(s string) {
	cases := []string{"\\a", "\\t", "\\b", "\\v", "\\r", "\\f"}
	for _, char := range cases {
		if strings.Contains(s, char) {
			fmt.Printf("Special case %q is not supported\n", char)
			os.Exit(1)
		}
	}
}
