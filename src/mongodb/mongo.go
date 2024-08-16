package mongodb

import (
	"context"
	"fieldweb/src/config"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx = context.Background()

func InitMongo(cfg config.MongoConfig) error {
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)).
		SetAuth(options.Credential{
			Username:      cfg.DBUser,
			Password:      cfg.DBPassword,
			AuthSource:    cfg.DBName,
			AuthMechanism: "SCRAM-SHA-256",
		})

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error creating MongoDB client: %v", err)
	}

	// Adding a timeout to the ping to avoid hanging
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check the connection
	err = client.Ping(ctxPing, nil)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}
	return nil
}

// GetClient returns the MongoDB client instance
func GetClient() *mongo.Client {
	return client
}

// CloseClient closes the MongoDB client connection
func CloseClient() {
	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Printf("Error disconnecting MongoDB client: %v", err)
		}
	}
}
