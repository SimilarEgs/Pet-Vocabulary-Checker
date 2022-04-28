package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//program reads name of the txt_file from input
//inside this file, every word marks like a «bad word»
//then when you type anything to CLI it would output same sentence
//with censored words.

//program can't read punctuation

func main() {
	fmt.Print("Enter name of the file without filename extension\n- ")
	var file_name string
	fmt.Scanln(&file_name)

	file, err := os.Open(file_name + ".txt")
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

	for {
		fmt.Print("\nCheck sentence for bad words: ")
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
		var s string
		for _, value := range input_words {
			s += value + " "

		}

		fmt.Printf("\nSentence after censorship: %s\n", s)

		input_words = nil
	}
}
