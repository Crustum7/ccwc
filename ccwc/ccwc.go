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

	bytesPtr := flag.Bool("c", false, "only outputs bytes if chosen")

	flag.Parse()

	rest := flag.Args()

	if *bytesPtr {
		log.Println("Trailing arguments:", rest)
		if len(rest) < 1 {
			log.Fatal("No file provided")
		}
		fileName := rest[0]
		HandleBytes(fileName)
	}
}

func HandleBytes(fileName string) {
	f, err := file.OpenFile(fileName)
	if err != nil {
		log.Fatal("Could not find file")
	}
	bytes, _ := file.Bytes(f)
	fmt.Printf("%v %v\n", bytes, fileName)
}
