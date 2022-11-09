package logger

import (
	"os"

	"github.com/fatih/color"
)

func Error(message string) {
	errColor := color.New(color.FgRed).PrintfFunc()
	errColor("[ERROR]: %s\n", message)
	os.Exit(1)
}

func Warn(message string) {
	yellow := color.New(color.FgYellow).PrintfFunc()
	yellow("[WARNING]: %s\n", message)
}

func Success(message string) {
	green := color.New(color.FgGreen).PrintfFunc()
	green("[SUCCESS]: %s\n", message)
}
