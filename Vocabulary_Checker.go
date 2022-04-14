package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var file_name string
	fmt.Scanln(&file_name)

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	taboo_words := make(map[string]bool)

	for scanner.Scan() {
		taboo_words[strings.ToLower(scanner.Text())] = true
	}

	input_words := make([]string, 0)
	word_scanner := bufio.NewScanner(os.Stdin)

	for true {
		word_scanner.Scan()
		line := word_scanner.Text()
		for _, words := range strings.Fields(line) {
			input_words = append(input_words, words)
		}
		for i := 0; i < len(input_words); i++ {
			if _, ok := taboo_words[strings.ToLower(input_words[i])]; ok {
				input_words[i] = strings.Repeat("*", len(input_words[i]))
			}
		}
		if line == "exit" {
			fmt.Println("Bye!")
			break
		}
		for _, value := range input_words {
			fmt.Printf("%s ", value)

		}
		fmt.Println()
		input_words = nil
	}
}
