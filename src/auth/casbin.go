package auth

import (
	"fieldweb/src/config"
	"fmt"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var enforcer *casbin.Enforcer

func InitCasbin(cfg config.CasbinConfig) error {
	//Create MongoDB client Option for casbin

	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)).
		SetAuth(options.Credential{
			Username:      cfg.DBUser,
			Password:      cfg.DBPassword,
			AuthSource:    cfg.DBName,
			AuthMechanism: "SCRAM-SHA-256",
		})

		// Initialize MongoDB adapter for Casbin with the client options
	adapter, err := mongodbadapter.NewAdapterWithClientOption(clientOptions, cfg.DBName)
	if err != nil {
		return fmt.Errorf("failed to create MongoDB adapter: %v", err)
	}

	// Initialize Casbin enforcer with the adapter and model.conf
	enforcer, err = casbin.NewEnforcer("src/config/model.conf", adapter)
	if err != nil {
		return fmt.Errorf("failed to create Casbin enforcer: %v", err)
	}

	return nil

}

// GetEnforcer returns the Casbin enforcer instance
func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
