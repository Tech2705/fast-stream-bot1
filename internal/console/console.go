package console

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Success(format string, a ...any) {
	color.Green(format, a...)
}

func Error(format string, a ...any) {
	color.Red(format, a...)
}

func Warning(format string, a ...any) {
	color.Yellow(format, a...)
}

func Info(format string, a ...any) {
	color.Blue(format, a...)
}

func Debug(format string, a ...any) {
	color.Cyan(format, a...)
}

func Fatal(format string, a ...any) {
	color.Red(format, a...)
	os.Exit(1)
}

func Print(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}
