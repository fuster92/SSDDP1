package utils

import (
	"fmt"
	"os"
)

// Checks if there's an error and exits.
func CheckError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
