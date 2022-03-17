package logger

import (
	"fmt"
	"time"

	"dev.io/v1/pkg/color"
)

func Log(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Println("[" + string(color.ColorGreen) + now + string(color.ColorReset) + "] " + message)
}

func LogError(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Println("[" + string(color.ColorRed) + now + string(color.ColorReset) + "] " + message)
}
