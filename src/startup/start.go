package startup

import (
	"fieldweb/src/auth"
	"fieldweb/src/config"
	"fieldweb/src/logger"
	"fieldweb/src/mongodb"
	"fmt"
)

func FieldWebStart() error {

	cfg, err := config.EnvGet()
	if err != nil {
		return err
	}

	// Set up logger
	if err := logger.SetupLogger(cfg.LoggerEnv); err != nil {
		return err
	}

	logger.Info("Logger initilized", "version", "1.0.0")

	// Initialize MongoDB

	err = mongodb.InitMongo(cfg.MongoDBEnv)
	if err != nil {

		msg := fmt.Sprintf("Failed to initialize MongoDB: %v", err)
		logger.Fatal(msg)
	}

	msg := fmt.Sprintf("Successfully connected to MongoDB on host: %s  Port: %d", cfg.MongoDBEnv.Host, cfg.MongoDBEnv.Port)
	logger.Info(msg)
	defer mongodb.CloseClient()

	err = auth.InitCasbin(cfg.CasbinEnv)

	if err != nil {
		msg := fmt.Sprintf("Failed to initialize Casbin: %v", err)
		logger.Fatal(msg)
	}

	msg = fmt.Sprintf("Successfully connected to Casbin MongoDB on host: %s Port: %d", cfg.CasbinEnv.Host, cfg.CasbinEnv.Port)
	logger.Info(msg)

	return nil

}
