package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"martinjonson.com/file"
)

func main() {
	log.SetPrefix("ccwc: ")
	log.SetFlags(0)
	log.SetOutput(io.Discard)

	bytesPtr := flag.Bool("c", false, "outputs number of bytes if chosen")
	linesPtr := flag.Bool("l", false, "outputs number of lines if chosen")
	wordsPtr := flag.Bool("w", false, "outputs number of white space seperated words if chosen")

	flag.Parse()

	rest := flag.Args()
	log.Println("Trailing arguments:", rest)
	if len(rest) < 1 {
		log.Fatal("No file provided")
	}
	fileName := rest[0]
	f, err := file.OpenFile(fileName)
	if err != nil {
		log.Fatal("Could not find file")
	}
	info, err := file.FileParseInfo(f)
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
	}

	fmt.Println(FormatOutput(fileName, fields...))
}

func FormatOutput(path string, args ...int) string {
	if len(args) == 0 {
		return path
	}
	return fmt.Sprintf("%6d ", args[0]) + FormatOutput(path, args[1:]...)
}
