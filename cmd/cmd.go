package cmd

import (
	"context"

	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/server"
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "token-athena",
	Short: "Token Service",
}

func Execute(ctx context.Context, appConfig *config.AppConfig) error {
	appCmd.AddCommand(startServer(ctx, appConfig))
	err := appCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func startServer(ctx context.Context, appConfig *config.AppConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the Token Service Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Start(ctx, appConfig)
		},
	}
}
