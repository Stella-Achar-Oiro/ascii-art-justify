package ascii

import (
	"fmt"
	"os"
)

// FileCheck checks the size of a specific banner file and exits if the file has been altered.
func FileCheck(fileName string) {
	var fileSize int64
	file_info, err := os.Stat("./banners/" + fileName + ".txt")
	if err != nil {
		fmt.Println(err)
		PrintError()
	}
	fileSize = file_info.Size()
	validSizes := map[string]int64{
		"standard":   6623,
		"thinkertoy": 5558,
		"shadow":     7463,
	}
	// Check if the file size matches the expected size for the given banner type.
	if validSize, ok := validSizes[fileName]; ok && fileSize != validSize {
		fmt.Printf("%s Banner file has been altered\n", fileName)
		os.Exit(1)
	}
}
