package main

import (
	"bufio"
	"fmt"
	"os"
)

// String matching with `*` wildcard support.
// Returns a slice with positions of all matched wildcardStr substrings in sourceStr.
func wildcardMatch(sourceStr, wildcardStr string) []int {
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Print("Interactive Demo of Wildcard Matching\n")
	for {
		fmt.Print("> ")
		scanner.Scan()
		text = scanner.Text()
		// exit command
		if text == "q" || text == "quit" {
			break
		}
		// other commands
		switch text {
		case "m", "match":
			var sourceStr, wildcardStr string
			fmt.Print("- Source string > ")
			scanner.Scan()
			sourceStr = scanner.Text()
			fmt.Print("- Wildcard string > ")
			scanner.Scan()
			wildcardStr = scanner.Text()
			result := wildcardMatch(sourceStr, wildcardStr)
			switch len(result) {
			case 0:
				fmt.Print("Substring not found.\n")
			case 1:
				fmt.Printf("Found at %v .\n", result)
			default:
				fmt.Printf("Found at: %v .\n", result)
			}
		case "h", "help":
			fallthrough
		default:
			fmt.Print("Available commands: \n")
			fmt.Print("m, match:	execute patten matching\n")
			fmt.Print("h, help: 	print this text\n")
			fmt.Print("q, quit: 	quit application\n")
		}
	}
}
