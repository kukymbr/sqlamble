package utils

import (
	"fmt"
	"os"
)

var silentMode bool

func SetSilentMode(silent bool) {
	silentMode = silent
}

func PrintHellof(format string, args ...any) {
	if silentMode {
		return
	}

	fmt.Printf("👋 "+format+"\n", args...)
}

func PrintDebugf(format string, args ...any) {
	if silentMode {
		return
	}

	fmt.Printf("⚙️ "+format+"\n", args...)
}

func PrintWarningf(format string, args ...any) {
	if silentMode {
		return
	}

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
	if silentMode {
		return
	}

	fmt.Printf("👍 "+format, args...)
}
