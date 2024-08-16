package logger

import (
	"fieldweb/src/config"
	"fmt"
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

var log *logrus.Logger

func SetupLogger(cfg config.LoggerConfig) error {

	log = logrus.New()

	//Set log Level
	level, err := logrus.ParseLevel(cfg.LogLevel)

	if err != nil {
		return fmt.Errorf("invalid log level: %v", err)
	}

	log.SetLevel(level)

	// Set the custom formatter
	log.SetFormatter(&CustomFormatter{})

	// Enable console logging if specified
	if cfg.Console {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(os.Stderr)
	}

	// Enable ELK logging if specified
	if cfg.ElkEnabled {
		elkURL := "http://" + cfg.ElkHost + ":" + cfg.ElkPort

		client, err := elastic.NewClient(
			elastic.SetURL(elkURL),
			elastic.SetSniff(false),
		)

		if err != nil {
			return fmt.Errorf("error creating Elasticsearch client: %v", err)
		}

		hook, err := elogrus.NewElasticHook(client, "localhost", level, cfg.ElkSearchIndex)

		if err != nil {
			return fmt.Errorf("error creating Elasticsearch hook: %v", err)
		}

		log.Hooks.Add(hook)
	}

	return nil
}

// Helper function to log with fields
// Helper function to log with key-value pairs
func withFields(args []interface{}) *logrus.Entry {
	fields := logrus.Fields{}

	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				continue // Skip invalid keys that aren't strings
			}
			fields[key] = args[i+1]
		}
	}

	return log.WithFields(fields)
}

// Log methods with additional key-value pairs// Log methods with additional key-value pairs
func Info(msg string, args ...interface{}) {
	withFields(args).Info(msg)
}

func Warn(msg string, args ...interface{}) {
	withFields(args).Warn(msg)
}

func Error(msg string, args ...interface{}) {
	withFields(args).Error(msg)
}

func Debug(msg string, args ...interface{}) {
	entry := withFields(args)
	entry.Debug(msg)
	entry.Debug("Stack trace:\n", getStackTrace())
}

func Fatal(msg string, args ...interface{}) {
	entry := withFields(args)
	entry.Fatal(msg)
	entry.Fatal("Stack trace:\n", getStackTrace())
}

func Panic(msg string, args ...interface{}) {
	entry := withFields(args)
	entry.Panic(msg)
	entry.Panic("Stack trace:\n", getStackTrace())
}
