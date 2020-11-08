package utils

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// ParseLogline extract talon fields fomr logline
func ParseLogline(logLine string) []string {

	const regex = "(^.*[PM|AM])[[:space:]]*Alt=[[:space:]]*(.*)[[:space:]].*Stat=[[:space:]]*([[:digit:]])[[:space:]].*L.act=[[:space:]]*([[:digit:]])[[:space:]]*.*SW=[[:space:]]([[:xdigit:]]*)[[:space:]].*Cond=[[:space:]]*([[:digit:]])"
	re := regexp.MustCompilePOSIX(regex)
	splitline := re.FindAllStringSubmatch(logLine, -1)
	/*
		fmt.Printf("date: %s\n", splitline[0][1])
		fmt.Printf("alt: %s\n", splitline[0][2])
		fmt.Printf("status: %s\n", splitline[0][3])
		fmt.Printf("Last action: %s\n", splitline[0][4])
		fmt.Printf("Switch: %s\n", splitline[0][5])
		fmt.Printf("condition: %s\n", splitline[0][6])
	*/
	return splitline[0]
}

// GetLastLineWithSeek get the loast line of the log
func GetLastLineWithSeek(filepath string) string {

	fileHandle, err := os.Open(filepath)
	if err != nil {
		panic("Cannot open talon log file")
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()

	for {
		cursor--
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 2)
		fileHandle.Read(char)

		// detect windows newline but not the first one
		if cursor != -2 && (char[1] == 10 && char[0] == 13) { // stop on newline
			break
		}

		line = fmt.Sprintf("%s%s", string(char[0]), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining of the file
			break
		}
	}

	return strings.TrimSpace(line)
}
