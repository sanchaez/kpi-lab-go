package main

import (
	"bufio"
	"fmt"
	"os"
)

//WildcardMatch does string matching with `*` wildcard support.
//Returns a slice with positions of all matched wildcardStr substrings in sourceStr.
func WildcardMatch(sourceStr, wildcardStr string) []int {
	return nil
}

/* Auxiliary functions */

//prompts user using given scanner and message
func promptLine(scanner *bufio.Scanner, message string) string {
	fmt.Print(message)
	scanner.Scan()
	return scanner.Text()
}

//prints help
func helpOption() {
	fmt.Print("Available commands: \n")
	fmt.Print("m, match:	execute patten matching\n")
	fmt.Print("h, help: 	print this text\n")
	fmt.Print("q, quit: 	quit application\n")
}

//matches strings provided by user
func matchOption(scanner *bufio.Scanner) {
	sourceStr := promptLine(scanner, "- Source string > ")
	wildcardStr := promptLine(scanner, "- Wildcard string > ")
	result := WildcardMatch(sourceStr, wildcardStr)
	switch len(result) {
	case 0:
		fmt.Print("Substring not found.\n")
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
