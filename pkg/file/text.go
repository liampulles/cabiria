package file

import (
	"bufio"
	"os"
)

// ReadLinesFromTextFile is a convenience method for extracting lines of text from
// a text file. The end of line character is omitted.
func ReadLinesFromTextFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
