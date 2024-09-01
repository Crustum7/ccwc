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

	flag.Parse()

	rest := flag.Args()
	log.Println("Trailing arguments:", rest)
	if len(rest) < 1 {
		log.Fatal("No file provided")
	}
	fileName := rest[0]

	if *bytesPtr {
		log.Println("Bytes was chosen")
		HandleBytes(fileName)
	} else if *linesPtr {
		log.Println("Lines was chosen")
		HandleLines(fileName)
	}
}

func HandleBytes(fileName string) {
	f, err := file.OpenFile(fileName)
	if err != nil {
		log.Fatal("Could not find file")
	}
	bytes, _ := file.Bytes(f)
	fmt.Printf("%7d %s\n", bytes, fileName)
}

func HandleLines(fileName string) {
	f, err := file.OpenFile(fileName)
	if err != nil {
		log.Fatal("Could not find file")
	}
	lines := file.Lines(f)
	fmt.Printf("%7d %s\n", lines, fileName)
}
