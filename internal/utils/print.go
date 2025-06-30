package utils

import (
	"fmt"
	"os"
)

func PrintHellof(format string, args ...any) {
	fmt.Printf("👋 "+format+"\n", args...)
}

func PrintDebugf(format string, args ...any) {
	fmt.Printf("⚙️ "+format+"\n", args...)
}

func PrintWarningf(format string, args ...any) {
	message := "⚠️ WARNING: " + fmt.Sprintf(format, args...) + "\n"

	if _, err := fmt.Fprint(os.Stderr, message); err != nil {
		fmt.Print(message)
	}
}

func PrintErrorf(format string, args ...any) {
	message := "🚫 ERROR: " + fmt.Sprintf(format, args...) + "\n"

	if _, err := fmt.Fprint(os.Stderr, message); err != nil {
		fmt.Print(message)
	}
}

func PrintSuccessf(format string, args ...any) {
	fmt.Printf("👍 "+format, args...)
}
