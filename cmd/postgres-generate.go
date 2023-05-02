package main

import (
	"context"

	"github.com/00mohamad00/go-db-bench/internal/postgres"
	"github.com/spf13/cobra"
)

var postgresGenerateCmd = &cobra.Command{
	Use:   "postgres-generate",
	Short: "start postgres-generate",
	Run:   postgresGenerate,
}

func init() {
	rootCmd.AddCommand(postgresGenerateCmd)
}

func postgresGenerate(cmd *cobra.Command, _ []string) {
	conf := loadConfigOrPanic(cmd)
	postgresDatabase := postgres.NewPostgresStorageOrPanic(conf.Postgres)

	err := postgresDatabase.Flush(context.Background())
	if err != nil {
		panic(err)
	}

	err = postgresDatabase.GenerateRandomData(context.Background(), 1000000)
	if err != nil {
		panic(err)
	}
}
