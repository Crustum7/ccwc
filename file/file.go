package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

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

func Bytes(file *os.File) (int64, error) {
	stats, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf(`could not load bytes from file %v`, file.Name())
	}
	return stats.Size(), nil
}

func Lines(file *os.File) int64 {
	lines := int64(0)

	reader := bufio.NewReader(file)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		lines++
	}

	return lines
}
