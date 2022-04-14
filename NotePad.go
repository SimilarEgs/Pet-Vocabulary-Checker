package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var notes_number int
	chekcer := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	notes := make([]string, 0, notes_number)

	for {
		fmt.Print("Enter the maximum number of notes: ")
		chekcer.Scan()
		if n, err := strconv.Atoi(chekcer.Text()); err == nil {
			notes_number = n
			break
		} else {
			err = errors.New("[Error] The amounts of notes must be a positive number, try again.\n")
			fmt.Println(err)
		}
	}
	fmt.Println()
	for {
		fmt.Printf("Enter command and data: ")
		scanner.Scan()

		var note string
		var note_index int
		var note_index_str string
		var note_index_checker int

		line := strings.SplitAfterN(scanner.Text(), " ", 2)
		command := line[0]
		command = strings.ReplaceAll(line[0], " ", "")
		command = strings.ToLower(command)

		if command == "create" && len(line) > 1 {
			note = strings.TrimSpace(line[1])
		}
		if len(line) > 1 {
			line[1] = strings.TrimSpace(line[1])
		}

		switch command {
		case "create":
			if len(notes) < notes_number && len(note) >= 1 {
				notes = append(notes, note)
				fmt.Println("[OK] The note was successfully created")
			} else if len(notes) == notes_number && len(note) >= 1 {
				fmt.Println("[Error] Notepad is full")
			} else if len(note) < 1 {
				fmt.Println("[Error] Missing note argument")
			}
		case "list":
			if len(notes) >= 1 {
				for i := 0; i < len(notes); i++ {
					fmt.Printf("[Info] %d: %s\n", i+1, notes[i])
				}
			} else {
				fmt.Println("[Info] Notepad is empty")
			}
		case "clear":
			notes = nil
			fmt.Println("[OK] All notes were successfully deleted")
		case "update":
			if command == "update" && len(line) > 1 {
				line = strings.SplitAfterN(scanner.Text(), " ", 3)
				if len(line) > 1 {
					note_index_str = strings.TrimSpace(line[1])
					if note_index_str != "" {
						note_index_checker, _ = strconv.Atoi(note_index_str)
						if note_index_checker > 0 {
							note_index = note_index_checker
							if len(line) > 2 {
								note = strings.TrimSpace(line[2])
							}
							if len(line) == 3 && note != "" && len(notes) >= note_index && note_index >= 0 {
								notes[note_index-1] = note
								fmt.Printf("[OK] The note at position %d was successfully updated\n", note_index)
							} else if note_index-1 > len(notes) {
								fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", note_index, notes_number)
							} else if note_index-1 > len(notes)-1 && note != "" {
								fmt.Println("[Error] There is nothing to update")
							} else if len(notes) >= note_index-1 && note == "" {
								fmt.Println("[Error] Missing note argument")
							}
						} else {
							fmt.Printf("[Error] Invalid position: %s\n", note_index_str)
						}
					} else if note_index == 0 {
						fmt.Println("[Error] Missing position argument")
						fmt.Println()
						continue
					}
				}
			} else if len(line) == 1 {
				fmt.Println("[Error] Missing position argument")
			}
		case "delete":
			if command == "delete" && len(line) > 1 && line[1] != "" {
				if len(line) > 1 {
					note_index_str = strings.TrimSpace(line[1])
					if note_index_str != "" {
						note_index_checker, _ = strconv.Atoi(note_index_str)
						if note_index_checker > 0 {
							note_index = note_index_checker
							if len(notes) > note_index-1 {
								note_index--
								copy(notes[note_index:], notes[note_index+1:])
								notes[len(notes)-1] = ""
								notes = notes[:len(notes)-1]
								fmt.Printf("[OK] The note at position %d was successfully updated\n", note_index)
							} else if note_index-1 > len(notes)-1 {
								fmt.Println("[Error] There is nothing to delete")
							}
						} else {
							fmt.Printf("[Error] Invalid position: %s\n", note_index_str)
						}
					}
				} else if note_index == 0 {
					fmt.Println("[Error] Missing position argument")
				}
			} else if len(line) == 1 {
				fmt.Println("[Error] Missing position argument")
			} else {
				fmt.Println("[Error] Missing position argument")
			}
		case "exit":
			fmt.Print("[Info] Bye!")
			os.Exit(0)
		default:
			fmt.Println("[Error] Unknown command")
		}
		fmt.Println()
	}
}
