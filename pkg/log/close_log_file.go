package log

import (
	"fmt"
	"os"
)

func Close(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Printf("Error closing log file: %v", err)
	}
}
