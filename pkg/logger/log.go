package logger

import (
	"fmt"
	"time"

	"dev.io/v1/pkg/console_color"
)

func Log(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Println("[" + string(console_color.ColorGreen) + now + string(console_color.ColorReset) + "] " + message)
}

func LogError(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Println("[" + string(console_color.ColorRed) + now + string(console_color.ColorReset) + "] " + message)
}
