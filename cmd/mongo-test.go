package main

import (
	"context"
	"fmt"
	"time"

	"github.com/00mohamad00/go-db-bench/internal/mongo"
	"github.com/00mohamad00/go-db-bench/pkg/storage"
	"github.com/00mohamad00/go-db-bench/utils"
	"github.com/spf13/cobra"
)

var mongoTestCmd = &cobra.Command{
	Use:   "mongo-test",
	Short: "start mongo-test",
	Run:   mongoTest,
}

func init() {
	rootCmd.AddCommand(mongoTestCmd)
}

func mongoTest(cmd *cobra.Command, _ []string) {
	conf := loadConfigOrPanic(cmd)
	mongoDatabase := mongo.NewMongoStorageOrPanic(conf.Mongo)

	testRecords := []*storage.Record{
		{Name: utils.GenerateRandomString(6), PaymentID: utils.GenerateRandomString(6), CreateAt: time.Now()},
		{Name: utils.GenerateRandomString(6), PaymentID: utils.GenerateRandomString(6), CreateAt: time.Now()},
	}

	err := mongoDatabase.ImportRecords(context.Background(), testRecords...)
	if err != nil {
		panic(err)
	}

	var completeTime time.Duration
	numberOfRun := 100
	for i := 0; i < numberOfRun; i++ {
		start := time.Now()
		var count int64
		count, err = mongoDatabase.CountRecordsWithNames(context.Background(), testRecords[0].Name, testRecords[1].Name)
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
