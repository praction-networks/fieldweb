package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

// CustomFormatter defines the format for log messages with colors
type CustomFormatter struct{}

// Define colors for different log levels
const (
	infoColor  = "\033[34m" // Blue
	warnColor  = "\033[33m" // Yellow
	errorColor = "\033[31m" // Red
	debugColor = "\033[36m" // Cyan
	resetColor = "\033[0m"  // Reset
)

// Format formats the log entry into a string with custom colors and details
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Get the process ID
	pid := os.Getpid()

	// Determine color based on log level
	var levelColor string
	switch entry.Level {
	case logrus.InfoLevel:
		levelColor = infoColor
	case logrus.WarnLevel:
		levelColor = warnColor
	case logrus.ErrorLevel:
		levelColor = errorColor
	case logrus.DebugLevel:
		levelColor = debugColor
	default:
		levelColor = resetColor
	}

	// Create the log message with color and additional details
	logMessage := fmt.Sprintf(
		"%s%s [%s] [PID:%d] %s%s %s%s\n",
		levelColor,
		entry.Time.Format(time.RFC3339),
		entry.Level.String(),
		pid,
		resetColor,
		entry.Message,
		formatFields(entry.Data),
		resetColor,
	)

	return []byte(logMessage), nil
}

// Helper function to format log fields
func formatFields(fields logrus.Fields) string {
	if len(fields) == 0 {
		return ""
	}

	var formattedFields string
	for key, value := range fields {
		formattedFields += fmt.Sprintf(" [%s:%v]", key, value)
	}
	return formattedFields
}

// Helper function to get stack trace
func getStackTrace() string {
	stackBuf := make([]byte, 1024)
	n := runtime.Stack(stackBuf, false)
	return string(stackBuf[:n])
}
