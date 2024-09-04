package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"martinjonson.com/file"
)

func main() {
	log.SetPrefix("ccwc: ")
	log.SetFlags(0)

	bytesPtr := flag.Bool("c", false, "outputs number of bytes if chosen")
	linesPtr := flag.Bool("l", false, "outputs number of lines if chosen")
	wordsPtr := flag.Bool("w", false, "outputs number of white space seperated words if chosen")
	charsPtr := flag.Bool("m", false, "outputs number of characters if chosen")

	flag.Parse()

	fileName := ""
	width := 6
	var f *os.File

	rest := flag.Args()
	if len(rest) < 1 {
		f = os.Stdin
		width = 7
	} else {
		fileName = rest[0]
		var err error
		f, err = os.Open(fileName)
		if err != nil {
			log.Fatalf("Could not find file with name %s", fileName)
		}
		defer f.Close()
	}

	info, err := file.NewFileInfo(f)
	if err != nil {
		log.Fatal("Could not get file info")
	}

	noFlags := !(*bytesPtr || *linesPtr || *wordsPtr)
	var fields []int
	if noFlags {
		fields = []int{info.Lines(), info.Words(), info.Bytes()}
	} else {
		fields = make([]int, 0)
		if *bytesPtr {
			fields = append(fields, info.Bytes())
		}
		if *linesPtr {
			fields = append(fields, info.Lines())
		}
		if *wordsPtr {
			fields = append(fields, info.Words())
		}
		if *charsPtr {
			fields = append(fields, info.Chars())
		}
	}

	if len(fields) == 1 {
		fmt.Println(fields[0], fileName)
	} else {
		fmt.Println(FormatOutput(fileName, width, fields...))
	}
}

func FormatOutput(path string, width int, args ...int) string {
	if len(args) == 0 {
		return path
	}
	return fmt.Sprintf("%*d ", width, args[0]) + FormatOutput(path, width, args[1:]...)
}
