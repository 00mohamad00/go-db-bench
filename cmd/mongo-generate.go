package main

import (
	"context"

	"github.com/00mohamad00/go-db-bench/internal/mongo"
	"github.com/spf13/cobra"
)

var mongoGenerateCmd = &cobra.Command{
	Use:   "mongo-generate",
	Short: "start mongo-generate",
	Run:   mongoGenerate,
}

func init() {
	rootCmd.AddCommand(mongoGenerateCmd)
}

func mongoGenerate(cmd *cobra.Command, _ []string) {
	conf := loadConfigOrPanic(cmd)
	mongoDatabase := mongo.NewMongoStorageOrPanic(conf.Mongo)

	err := mongoDatabase.Flush(context.Background())
	if err != nil {
		panic(err)
	}

	err = mongoDatabase.GenerateRandomData(context.Background(), 1000000)
	if err != nil {
		panic(err)
	}
}
