package main

import (
	"context"
	"fmt"
	"time"

	"github.com/00mohamad00/go-db-bench/internal/postgres"
	"github.com/00mohamad00/go-db-bench/pkg/storage"
	"github.com/00mohamad00/go-db-bench/utils"
	"github.com/spf13/cobra"
)

var postgresTestCmd = &cobra.Command{
	Use:   "postgres-test",
	Short: "start postgres-test",
	Run:   postgresTest,
}

func init() {
	rootCmd.AddCommand(postgresTestCmd)
}

func postgresTest(cmd *cobra.Command, _ []string) {
	conf := loadConfigOrPanic(cmd)
	postgresDatabase := postgres.NewPostgresStorageOrPanic(conf.Postgres)

	testRecords := []*storage.Record{
		{Name: utils.GenerateRandomString(6), PaymentID: utils.GenerateRandomString(6), CreateAt: time.Now()},
		{Name: utils.GenerateRandomString(6), PaymentID: utils.GenerateRandomString(6), CreateAt: time.Now()},
	}

	err := postgresDatabase.ImportRecords(context.Background(), testRecords...)
	if err != nil {
		panic(err)
	}

	var completeTime time.Duration
	numberOfRun := 100
	for i := 0; i < numberOfRun; i++ {
		start := time.Now()
		var count int64
		count, err = postgresDatabase.CountRecordsWithNames(context.Background(), testRecords[0].Name, testRecords[1].Name)
		if err != nil {
			panic(err)
		}
		elapsed := time.Since(start)
		completeTime += elapsed
		fmt.Println("document count ", count)
		fmt.Println("found process take ", elapsed)
	}
	fmt.Println("average process take ", completeTime/time.Duration(numberOfRun))

}
