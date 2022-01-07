package files

import (
	"os"
	"strings"
)

// ReadFileByLines reads the lines of a file into argument "lines"
func ReadFileByLines(filePath string, lines []string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	lines = strings.Split(string(file), "\n")
}
