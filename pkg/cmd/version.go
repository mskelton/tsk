package cmd

import (
	"fmt"
	"os"
)

func Version() {
	fmt.Println(os.Getenv("VERSION"))
}
