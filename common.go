package common

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadTestFile(fileName string) *os.File {
	var filePath, _ = filepath.Abs(filepath.Join("./test_data", fileName))
	inputFile, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("bronk CANNOT READ TEST FILE AAAA\n%s", err)
	}

	return inputFile
}
