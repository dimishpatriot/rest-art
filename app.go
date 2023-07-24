package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/dimishpatriot/rest-art-of-development/internal/config"
	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
	"github.com/dimishpatriot/rest-art-of-development/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.CreateLogger()
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	router := httprouter.New()
	logger.Infof("[OK] router created: %+v", *router)

	user.NewHandler(logger).Register(router)
	logger.Info("[OK] user handler registered")

	start(router, cfg, logger)
}

func start(router *httprouter.Router, cfg *config.Config, logger *logging.Logger) {
	network, address := getNetworkInfo(cfg)

	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Fatalf("can't get listener: %s", err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  1 * time.Second,
	}

	logger.Infof(fmt.Sprintf("server %+v started at <%s:%s>...", server, network, address))
	logger.Fatalf("server can't start: %s", server.Serve(listener))
}

func getNetworkInfo(cfg *config.Config) (string, string) {
	switch cfg.Listen.Type {
	case "sock":
		appDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		return "unix", path.Join(appDir, "app.sock")
	case "port":
		return "tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	default:
		panic("incorrect network")
	}
}
