package main

import (
	"context"
	"fmt"
	"time"

	"github.com/00mohamad00/go-db-bench/internal/mongo"
	"github.com/00mohamad00/go-db-bench/internal/postgres"
	"github.com/00mohamad00/go-db-bench/pkg/storage"
	"github.com/00mohamad00/go-db-bench/utils"
)

func main() {
	testRecords := []*storage.Record{
		{Name: utils.GenerateRandomString(6), PaymentID: utils.GenerateRandomString(6), CreateAt: time.Now()},
		{Name: utils.GenerateRandomString(6), PaymentID: utils.GenerateRandomString(6), CreateAt: time.Now()},
	}
	testMongo(testRecords)
}

func testMongo(testRecords []*storage.Record) {
	mongoConfig := mongo.Config{
		URL:      "mongodb://localhost:27017",
		Database: "testing",
		Timeout:  time.Second,
	}

	mongoDatabase, err := mongo.NewMongoStorage(mongoConfig)
	if err != nil {
		panic(err)
	}

	//err = mongoDatabase.Flush(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = mongoDatabase.GenerateRandomData(context.Background(), 1000000)
	//if err != nil {
	//	panic(err)
	//}

	err = mongoDatabase.ImportRecords(context.Background(), testRecords...)
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

func testPostgres(testRecords []*storage.Record) {
	postgresConfig := postgres.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "test",
		Password: "test",
		DBName:   "test",
	}

	postgresDatabase, err := postgres.NewPostgresStorage(postgresConfig)
	if err != nil {
		panic(err)
	}

	//
	//err = postgresDatabase.Flush(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = postgresDatabase.GenerateRandomData(context.Background(), 1000000)
	//if err != nil {
	//	panic(err)
	//}

	err = postgresDatabase.ImportRecords(context.Background(), testRecords...)
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
