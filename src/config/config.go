package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	Environment string       `mapstructure:"environment"`
	LoggerEnv   LoggerConfig `mapstructure:"logger"`
	MongoDBEnv  MongoConfig  `mapstructure:"mongodb"`
	CasbinEnv   CasbinConfig `mapstructure:"casbin"`
	FiberEnv    FiberConfig  `mapstructure:"server"`
	JWTEnv      JWTConfig    `mapstructure:"jwt"`
}

type LoggerConfig struct {
	LogLevel       string `mapstructure:"logLevel"`
	Console        bool   `mapstructure:"console"`
	ElkEnabled     bool   `mapstructure:"elkEnabled"`
	ElkHost        string `mapstructure:"elkHost"`
	ElkPort        string `mapstructure:"elkPort"`
	ElkSearchIndex string `mapstructure:"elkSearchIndex"`
}

type CasbinConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	TLSEnabled bool   `mapstructure:"tlsEnabled"`
	DBName     string `mapstructure:"database"`
	DBUser     string `mapstructure:"username"`
	DBPassword string `mapstructure:"password"`
}

type MongoConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	TLSEnabled bool   `mapstructure:"tlsEnabled"`
	DBName     string `mapstructure:"database"`
	DBUser     string `mapstructure:"username"`
	DBPassword string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime string `mapstructure:"expiration"`
}

type FiberConfig struct {
	Port string `mapstructure:"port"`
}

func EnvGet() (EnvConfig, error) {
	viper.SetConfigFile("src/config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return EnvConfig{}, errors.New("error reading config file: " + err.Error())
	}

	var envConfig EnvConfig
	if err := viper.Unmarshal(&envConfig); err != nil {
		return EnvConfig{}, errors.New("unable to decode into struct: " + err.Error())
	}

	return envConfig, nil
}

func LoggerEnvGet() (LoggerConfig, error) {
	envConfig, err := EnvGet()
	if err != nil {
		msg := fmt.Sprintf("Unable to get Logger Env Config Error : %v", err)
		return LoggerConfig{}, errors.New(msg)
	}

	if !envConfig.LoggerEnv.Console && !envConfig.LoggerEnv.ElkEnabled {
		return LoggerConfig{}, errors.New("at least one logging option (Console or ELK) must be enabled")
	}

	return envConfig.LoggerEnv, nil
}

func MongoEnvGet() (MongoConfig, error) {
	envConfig, err := EnvGet()
	if err != nil {
		msg := fmt.Sprintf("Unable to get MongoDB Env Config Error : %v", err)
		return MongoConfig{}, errors.New(msg)
	}

	return envConfig.MongoDBEnv, nil
}

func CasbinEnvGet() (CasbinConfig, error) {
	envConfig, err := EnvGet()
	if err != nil {
		msg := fmt.Sprintf("Unable to get Casbin Env Config Error : %v", err)
		return CasbinConfig{}, errors.New(msg)
	}

	return envConfig.CasbinEnv, nil
}

func FiberEnvGet() (FiberConfig, error) {
	envConfig, err := EnvGet()
	if err != nil {
		msg := fmt.Sprintf("Unable to get Fiber Env Config Error : %v", err)
		return FiberConfig{}, errors.New(msg)
	}

	return envConfig.FiberEnv, nil
}

func JWTEnvGet() (JWTConfig, error) {
	envConfig, err := EnvGet()
	if err != nil {
		msg := fmt.Sprintf("Unable to get JWT Env Config Error : %v", err)
		return JWTConfig{}, errors.New(msg)
	}

	return envConfig.JWTEnv, nil
}
