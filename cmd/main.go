package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var numBytes = flag.Bool("c", false, "number of bytes in a file")
	var numLines = flag.Bool("l", false, "number of lines in a file")
	var numWords = flag.Bool("w", false, "number of words in a file")
	var numChars = flag.Bool("m", false, "number of characters in a file")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("missing filename, please have format: cwd flags(optional) filename")
		os.Exit(1)
	} else {
		read(*numBytes, *numLines, *numWords, *numChars, args[0])
	}

}

func read(bytes bool, lines bool, words bool, chars bool, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	byteCount := len(data)
	charCount := utf8.RuneCountInString(string(data))
	lineCount := 0
	wordCount := len(strings.Fields(string(data)))
	for _, v := range data {
		if v == '\n' {
			lineCount++
		}
	}
	if len(data) > 0 {
		lineCount++
	}

	if bytes {
		fmt.Printf("The number of bytes in the file is : %d\n", byteCount)
	}
	if lines {
		fmt.Printf("The number of lines in the file is : %d\n", lineCount)
	}
	if words {
		fmt.Printf("The number of words in the file is : %d\n", wordCount)
	}
	if chars {
		fmt.Printf("The number of characters in the file is : %d\n", charCount)
	}
}
