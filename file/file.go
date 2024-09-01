package file

import (
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

func PrintFileStats(file *os.File) (string, error) {
	stats, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf(`could not load stats from file %v`, file.Name())
	}
	return fmt.Sprintf(`File %v, is of size %v and was last modified at %v`, file.Name(), stats.Size(), stats.ModTime()), nil
}
