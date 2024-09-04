package file

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

type FileInfo struct {
	bytes int
	lines int
	words int
	chars int
}

func (info *FileInfo) Bytes() int {
	return info.bytes
}

func (info *FileInfo) Lines() int {
	return info.lines
}

func (info *FileInfo) Words() int {
	return info.words
}

func (info *FileInfo) Chars() int {
	return info.chars
}

func NewFileInfo(file *os.File) (*FileInfo, error) {
	var info FileInfo
	info.lines = 0
	info.words = 0
	info.chars = 0
	info.bytes = 0

	reader := bufio.NewReader(file)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		info.lines++
		line := scan.Text()
		info.words += WordsInLine(line)

		// Account for new lines
		info.bytes += len(line) + 2
		info.chars += utf8.RuneCountInString(line) + 2
	}

	return &info, nil
}

func WordsInLine(line string) int {
	line = strings.ReplaceAll(line, "\t", " ")
	words := strings.Split(line, " ")
	length := 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) > 0 {
			length++
		}
	}
	return length
}
