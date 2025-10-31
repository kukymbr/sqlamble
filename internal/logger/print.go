package logger

import (
	"fmt"
	"os"
)

var silentMode bool

func SetSilentMode(silent bool) {
	silentMode = silent
}

func Hellof(format string, args ...any) {
	if silentMode {
		return
	}

	fmt.Printf("👋 "+format+"\n", args...)
}

func Debugf(format string, args ...any) {
	if silentMode {
		return
	}

	fmt.Printf("⚙️ "+format+"\n", args...)
}

func Warningf(format string, args ...any) {
	if silentMode {
		return
	}

	message := "⚠️ WARNING: " + fmt.Sprintf(format, args...) + "\n"

	if _, err := fmt.Fprint(os.Stderr, message); err != nil {
		fmt.Print(message)
	}
}

func Errorf(format string, args ...any) {
	message := "🚫 ERROR: " + fmt.Sprintf(format, args...) + "\n"

	if _, err := fmt.Fprint(os.Stderr, message); err != nil {
		fmt.Print(message)
	}
}

func Successf(format string, args ...any) {
	if silentMode {
		return
	}

	fmt.Printf("👍 "+format, args...)
}
