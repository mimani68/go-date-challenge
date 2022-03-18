package graceful_shutdown

import (
	"os"

	"dev.io/v1/pkg/logger"
)

func StopApplication(exitCode int, message string) {
	switch exitCode {
	case 0:
		logger.Log(message)
	case 1:
		logger.LogError(message)
	}
	os.Exit(exitCode)
}
