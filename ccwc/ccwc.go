package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"martinjonson.com/file"
)

type Flags struct {
	bytes bool
	lines bool
	words bool
	chars bool
}

func main() {
	log.SetPrefix("ccwc: ")
	log.SetFlags(0)

	bytesPtr := flag.Bool("c", false, "outputs number of bytes if chosen")
	linesPtr := flag.Bool("l", false, "outputs number of lines if chosen")
	wordsPtr := flag.Bool("w", false, "outputs number of white space separated words if chosen")
	charsPtr := flag.Bool("m", false, "outputs number of characters if chosen")

	flag.Parse()

	var flags Flags
	flags.bytes = *bytesPtr
	flags.lines = *linesPtr
	flags.words = *wordsPtr
	flags.chars = *charsPtr

	rest := flag.Args()
	if len(rest) < 1 {
		fields := handleFile(os.Stdin, flags)
		if len(fields) == 1 {
			fmt.Println(fields[0])
		} else {
			fmt.Println(FormatOutput("", 7, fields...))
		}
	} else {
		fileResults := make([][]int, 0)
		result := make([]int, 0)
		for _, fileName := range rest {
			f, err := os.Open(fileName)
			if err != nil {
				log.Fatalf("Could not find file with name %s", fileName)
			}
			defer f.Close()
			fields := handleFile(f, flags)
			fileResults = append(fileResults, make([]int, 0))
			fileResults[len(fileResults)-1] = fields
			for i, field := range fields {
				if i == len(result) {
					result = append(result, 0)
				}
				result[i] += field
			}
		}
		width := numberOfDigits(result[len(result)-1])
		for i, file := range fileResults {
			fmt.Println(FormatOutput(rest[i], width, file...))
		}
		fmt.Println(FormatOutput("total", width, result...))
	}
}

func numberOfDigits(num int) int {
	if num == 0 {
		return 1
	}
	return numberOfDigitsHelper(num)
}

func numberOfDigitsHelper(num int) int {
	if num == 0 {
		return 0
	}
	return numberOfDigitsHelper(num/10) + 1
}

func handleFile(f *os.File, flags Flags) []int {
	info, err := file.NewFileInfo(f)
	if err != nil {
		log.Fatal("Could not get file info")
	}

	return Fields(info, flags)
}

func Fields(info *file.FileInfo, flags Flags) []int {
	noFlags := !(flags.bytes || flags.lines || flags.words)
	var fields []int
	if noFlags {
		fields = []int{info.Lines(), info.Words(), info.Bytes()}
	} else {
		fields = make([]int, 0)
		if flags.bytes {
			fields = append(fields, info.Bytes())
		}
		if flags.lines {
			fields = append(fields, info.Lines())
		}
		if flags.words {
			fields = append(fields, info.Words())
		}
		if flags.chars {
			fields = append(fields, info.Chars())
		}
	}
	return fields
}

func FormatOutput(path string, width int, args ...int) string {
	if len(args) == 0 {
		return path
	}
	return fmt.Sprintf("%*d ", width, args[0]) + FormatOutput(path, width, args[1:]...)
}
