package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gookit/slog"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	_ "github.com/lib/pq"

	"github.com/ziggsdil/api-service-test/pkg/config"
	"github.com/ziggsdil/api-service-test/pkg/db"
	"github.com/ziggsdil/api-service-test/pkg/handler"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.Parse()
}

func main() {
	slog.SetFormatter(slog.NewJSONFormatter())
	slog.Info("Starting application")

	ctx := context.Background()

	var cfg config.Config

	err := confita.NewLoader(
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		slog.Errorf("Failed to load config: %v", err.Error())
		return
	}
	slog.Info("Config loaded")

	postgres, err := db.NewDatabase(cfg.Postgres)
	if err != nil {
		slog.Errorf("Failed to connect to postgres: %v", err.Error())
		return
	}
	slog.Info("Connected to postgres")

	err = postgres.Init(ctx)
	if err != nil {
		slog.Errorf("Failed to use migrate and init postgres: %v", err.Error())
		return
	}
	slog.Info("Successfully use migrate and init postgres")

	handlers := handler.NewHandler(postgres, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      handlers.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		slog.Infof("Starting server at %s", srv.Addr)
		_ = srv.ListenAndServe()
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
