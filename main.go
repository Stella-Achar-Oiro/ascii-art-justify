package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii-art-justify/banner"
)

func main() {
	var outputFileName string
	var input string
	var alignment string = "left" // Default alignment

	if len(os.Args) < 2 || len(os.Args) > 5 {
		ascii.PrintError()
		return
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--output=") && strings.HasSuffix(arg, ".txt") {
			outputFileName = strings.TrimPrefix(arg, "--output=")
		} else if strings.HasPrefix(arg, "--align=") {
			alignType := strings.TrimPrefix(arg, "--align=")
			if alignType != "left" && alignType != "center" && alignType != "right" && alignType != "justify" {
				ascii.PrintError()
				return
			}
			alignment = alignType
		} else {
			input = arg
		}
	}

	if input == "" {
		ascii.PrintError()
		return
	}

	filename := "standard"
	if len(os.Args) == 5 {
		filename = os.Args[4]
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	ascii.HandleSpecialCases(input)

	var formattedOutput strings.Builder
	if input == "\n" {
		formattedOutput.WriteString("\n")
	} else if input == "" {
		return
	}

	Input := strings.Split(input, "\n")
	spaceCount := 0
	for i, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				formattedOutput.WriteString("\n")
			}
		} else {
			bannerOutput := ascii.PrintBanner(word, filename)
			formattedOutput.WriteString(ascii.AlignText(bannerOutput, alignment))
			if i < len(Input)-1 {
				formattedOutput.WriteString("\n")
			}
		}
	}

	outputString := formattedOutput.String()

	// Print to console
	if strings.HasPrefix(os.Args[1], "--align=") {
		fmt.Println(outputString)
	} else if len(os.Args) == 2 {
		fmt.Println(outputString)
	}

	// Write to file if outputFileName is specified
	if outputFileName != "" {
		err := ascii.WriteToFile(outputFileName, outputString)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}
}
