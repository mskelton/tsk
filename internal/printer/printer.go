package printer

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Error(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}

func Message(message string) {
	color.Blue(message)
}

func Confirm(message string) bool {
	color.New().Add(color.Bold).Print(message)
	color.New().Print(" (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Failed to read input")
		return true
	}

	switch char {
	case 'y':
		return true
	case 'n':
		return false
	default:
		fmt.Println("Invalid input")
		return true
	}
}
