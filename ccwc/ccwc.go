package main

import (
	"flag"
	"fmt"
	"log"

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

	rest := flag.Args()
	if len(rest) < 1 {
		log.Fatal("No file provided")
	}

	fileName := rest[0]
	f, err := file.OpenFile(fileName)
	if err != nil {
		log.Fatalf("Could not find file with name %s", fileName)
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

	fmt.Println(FormatOutput(fileName, fields...))
}

func FormatOutput(path string, args ...int) string {
	if len(args) == 0 {
		return path
	}
	return fmt.Sprintf("%6d ", args[0]) + FormatOutput(path, args[1:]...)
}
