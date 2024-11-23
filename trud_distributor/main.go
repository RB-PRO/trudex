package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trudex/trud_distributor/cmd"
)

const (
	configPatch     = "trud_distributor/trud_distributor.yaml"
	shutdownTimeout = 5 * time.Second
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	defer func() {
		if recoverPanic := recover(); recoverPanic != nil {
			log.Println("recover panic")
			os.Exit(1)
		}
	}()

	if closer, err := cmd.RunServer(ctx,
		cmd.WithConfigPatch(configPatch),
	); err != nil {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := closer.Close(shutdownCtx); err != nil {
			log.Printf("closer: %v\n", err)
			os.Exit(1)
		}
	}

	return
}
