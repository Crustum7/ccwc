package file

import (
	"bufio"
	"errors"
	"fmt"
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

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(`could not find file with name %v`, path)
	}
	return file, nil
}

func CloseFile(file *os.File) error {
	err := file.Close()
	return err
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
	bytes, err := Bytes(file)
	if err != nil {
		return &info, fmt.Errorf("could not get file info from %s", file.Name())
	}
	info.bytes = int(bytes)

	reader := bufio.NewReader(file)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		info.lines++
		line := scan.Text()
		info.words += WordsInLine(line)
		info.chars += utf8.RuneCountInString(line) + 2 // Account for new lines
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

func Bytes(file *os.File) (int64, error) {
	stats, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf(`could not load bytes from file %v`, file.Name())
	}
	return stats.Size(), nil
}
