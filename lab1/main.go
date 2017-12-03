package main

import (
	"bufio"
	"fmt"
	"os"
)

/* Auxiliary functions */

//prompts user using given scanner and message
func promptLine(scanner *bufio.Scanner, message string) string {
	fmt.Print(message)
	scanner.Scan()
	return scanner.Text()
}

//prints help
func helpOption() {
	fmt.Println("Available commands:")
	fmt.Println("m, match:	execute pattern matching")
	fmt.Println("h, help: 	print this text")
	fmt.Println("q, quit: 	quit application")
}

//matches strings provided by user
func matchOption(scanner *bufio.Scanner) {
	sourceStr := promptLine(scanner, "- Source string > ")
	wildcardStr := promptLine(scanner, "- Wildcard string > ")
	result := Match(sourceStr, wildcardStr)
	switch len(result) {
	case 0:
		fmt.Println("Substring not found.")
	default:
		fmt.Printf("Found at: %v .\n", result)
	}
}

//checks if command is valid and does appropriate actions
//returns true if it was a quit command
func commandCheck(scanner *bufio.Scanner, command string) (quit bool) {
	switch command {
	case "q", "quit":
		quit = true
	case "m", "match":
		matchOption(scanner)
	case "h", "help":
		fallthrough
	default:
		helpOption()
	}
	return
}

//main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Print("Interactive Demo of Wildcard Matching\n")
	for {
		text = promptLine(scanner, "> ")
		if commandCheck(scanner, text) {
			break
		}
	}
}
