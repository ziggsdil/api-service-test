package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	_ "github.com/lib/pq"
	"github.com/ziggsdil/api-service-test/pkg/config"
	"github.com/ziggsdil/api-service-test/pkg/db"
	"github.com/ziggsdil/api-service-test/pkg/handler"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	var cfg config.Config

	err := confita.NewLoader(
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		// todo: log and info
		fmt.Printf("Failed to load config: %v", err.Error())
		return
	}

	postgres, err := db.NewDatabase(cfg.Postgres)
	fmt.Println(cfg)
	if err != nil {
		// todo: log and info
		fmt.Printf("Failed to connect to postgres: %v", err.Error())
		return
	}

	err = postgres.Init(ctx)
	if err != nil {
		// todo: log and info
		fmt.Printf("Failed to use migrate and init postgres: %v", err.Error())
		return
	}

	handlers := handler.NewHandler(postgres, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handlers.Router(),
	}

	go func() {
		// todo: log and info
		fmt.Printf("Starting server at %s\n", srv.Addr)
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
