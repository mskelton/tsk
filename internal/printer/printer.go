package printer

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mskelton/tsk/internal/utils"
)

func Error(error utils.CLIError) {
	// If no message is provided, use the error message
	message := error.Message
	if message == "" && error.Err != nil {
		message = error.Err.Error()
	}

	fmt.Fprintf(os.Stderr, "%s\n", error.Message)

	// If there is a detail message, print it in faint color
	if error.Detail != "" {
		color.New().Add(color.Faint).Println(error.Detail)
	}

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
