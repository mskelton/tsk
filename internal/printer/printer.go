package printer

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Error(message string, err error) {
	fmt.Println("Error: ", err)
	// let code = format!("[{}]", error.code());
	// eprintln!("{}: {}", code.red().bold(), message);
	//
	// let message = match error {
	//     Error::Custom(message) => message.to_string(),
	//     Error::MigrationFailure(message) => format!("Migration failure: {}", message),
	//     Error::QueryError(message) => format!("Query error: {}", message),
	//     Error::Serialize(message) => message,
	//     Error::Deserialize(message) => message,
	//     _ => "".to_string(),
	// };

	if message != "" {
		color.New().Add(color.Faint).Println(message)
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
