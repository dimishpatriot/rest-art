package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/dimishpatriot/rest-art-of-development/internal/client/mongodb"
	"github.com/dimishpatriot/rest-art-of-development/internal/config"
	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
	"github.com/dimishpatriot/rest-art-of-development/internal/user"
	"github.com/dimishpatriot/rest-art-of-development/internal/user/db"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	ctx := context.Background()
	client := mongoConnect(ctx, cfg, logger)
	collection := db.NewCollection(client, cfg.Storage.Collection)

	// test
	id, err := collection.Create(ctx, &user.User{
		Username:     "Pop",
		PasswordHash: "1234",
		Email:        "example@example.com",
	})
	if err != nil {
		logger.Error(err)
	}
	logger.Debugf("new user ID: %s", id)
	// ---

	router := httprouter.New()
	logger.Infof("router created: %+v", router)
	user.NewHandler(logger).Register(router)
	logger.Info("user handler registered")

	start(router, cfg, logger)
}

func mongoConnect(ctx context.Context, cfg *config.Config, logger *logging.Logger) *mongo.Database {
	dbClient, err := mongodb.NewClient(ctx, &mongodb.MongoParams{
		Host:     cfg.Storage.Host,
		Port:     cfg.Storage.Port,
		Database: cfg.Storage.Database,
		Username: cfg.Storage.Username,
		Password: cfg.Storage.Password,
	})
	if err != nil {
		logger.Error(err)
	}
	return dbClient
}

func start(router *httprouter.Router, cfg *config.Config, logger *logging.Logger) {
	logger.Info("start app")

	var err error
	var listener net.Listener
	var address, network string

	if cfg.Listen.Type == "sock" {
		network = "unix"

		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		address = path.Join(appDir, "app.sock")
	} else {
		network = "tcp"

		address = fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	listener, err = net.Listen(network, address)
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  1 * time.Second,
	}

	logger.Infof(fmt.Sprintf("server listen at %s: %s", network, address))
	logger.Fatal(server.Serve(listener))
}
