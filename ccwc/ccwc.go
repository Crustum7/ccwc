package main

import (
	"fmt"
	"log"

	"martinjonson.com/file"
)

func main() {
	log.SetPrefix("ccwc: ")
	log.SetFlags(0)

	message, err := file.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
