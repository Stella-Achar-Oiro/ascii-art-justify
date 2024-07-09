package ascii

import (
	"fmt"
	"os"
)

func PrintError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
	os.Exit(1)
}
