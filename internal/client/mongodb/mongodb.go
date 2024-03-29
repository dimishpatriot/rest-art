package mongodb

import (
	"context"
	"fmt"

	"github.com/dimishpatriot/rest-art-of-development/internal/config"
	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoParams struct {
	Database, Host, Port string
	Username, Password   string
}

type UserCred struct {
	Username, Password string
}

func NewClient(ctx context.Context, params *MongoParams) (*mongo.Database, error) {
	var err error

	logger := logging.GetLogger()

	uri := fmt.Sprintf("mongodb://%s:%s", params.Host, params.Port)
	clientOptions := options.Client().ApplyURI(uri)

	if params.Username != "" && params.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: params.Username,
			Password: params.Password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	logger.Infof("[OK] mongo client at URI=%s created: %+v", uri, *client)

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Infof("[OK] database <%s> available", params.Database)

	return client.Database(params.Database), nil
}

func Connect(ctx context.Context, cfg *config.Config, logger *logging.Logger) *mongo.Database {
	client, err := NewClient(ctx, &MongoParams{
		Host:     cfg.Storage.Host,
		Port:     cfg.Storage.Port,
		Database: cfg.Storage.Database,
		Username: cfg.Storage.Username,
		Password: cfg.Storage.Password,
	})
	if err != nil {
		logger.Fatalf("can't get db client: %s", err)
	}
	return client
}
