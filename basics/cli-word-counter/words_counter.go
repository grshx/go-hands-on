package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	fileFlag := flag.String("file", "", "path of the file to count words in")
	flag.Parse()

	var text string
	var scanner *bufio.Scanner

	if *fileFlag == "" {
		fmt.Println("Enter text to count words (press Ctrl+D to finish):")
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Printf("something went wrong while opening the file \n%v", err.Error())
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		text += strings.TrimSpace(scanner.Text()) + " "
	}

	err := scanner.Err()

	if err != nil {
		fmt.Printf("something went wrong while scanning text \n %v", err.Error())
		return
	}

	var countByWords = make(map[string]int)

	for _, word := range strings.Split(text, " ") {
		if strings.TrimSpace(word) == "" {
			continue
		}

		if _, ok := countByWords[word]; ok {
			countByWords[word]++
		} else {
			countByWords[word] = 1
		}

	}

	totalWordCount := 0
	for word, count := range countByWords {
		totalWordCount += count
		fmt.Printf("%s : %d\n", word, count)
	}

	fmt.Println("total words count : ", totalWordCount)
}
