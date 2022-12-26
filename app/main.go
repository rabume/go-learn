package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fileToRead := os.Args[1]
	fileToRead := "myFile.txt"

	content, err := os.ReadFile(fileToRead)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
