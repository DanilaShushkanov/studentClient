package app

import (
	"context"
	"fmt"
	"github.com/danilashushkanov/studentClient/internal/closer"
	"github.com/danilashushkanov/studentClient/internal/config"
	"github.com/danilashushkanov/studentClient/internal/http"
	"github.com/danilashushkanov/studentClient/internal/student"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c, err := student.NewClient(cfg)
	if err != nil {
		log.Printf("failed to create client")
	}

	httpServer := http.New(ctx, cfg, c)
	httpServer.Start()
	closer.Add(
		func() error {
			return httpServer.Stop()
		},
	)

	gracefulShutdown(cancel)
	return nil
}

func gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	fmt.Println(errorMessage)
	cancel()
	closer.CloseAll()
}
