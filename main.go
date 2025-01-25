package main

import (
	"context"
	"fmt"
	"log"

	"os/signal"
	"syscall"

	"github.com/piyushverma013/token-athena/cmd"
	"github.com/piyushverma013/token-athena/config"
)

type peripherals struct {
	appConfig *config.AppConfig
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	peripherals, err := initPeripherals(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize peripherals: %v", err)
	}

	err = cmd.Execute(ctx, peripherals.appConfig)
	if err != nil {
		log.Fatalf("Failed to execute the command: %v", err)
	}

}

func initPeripherals(ctx context.Context) (peripherals peripherals, err error) {
	appConfig, err := config.InitConfig(ctx)
	if err != nil {
		return peripherals, fmt.Errorf("cannot load appConfig, error: %w", err)
	}
	peripherals.appConfig = appConfig

	return peripherals, nil
}
