package mongo

import (
	"context"
	"time"

	"github.com/00mohamad00/go-db-bench/pkg/storage"
	"github.com/00mohamad00/go-db-bench/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Impl struct {
	collection *mongo.Collection
}

func newMongoClient(ctx context.Context, conf Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, conf.Timeout)
	defer cancel()

	opts := options.Client().ApplyURI(conf.URL)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoStorage(conf Config) (storage.Storage, error) {
	client, err := newMongoClient(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	impl := &Impl{
		collection: client.Database(conf.Database).Collection(CollectionName),
	}
	if err := impl.ensureIndexes(conf); err != nil {
		return nil, err
	}
	return impl, nil
}

func (i *Impl) ensureIndexes(conf Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout)
	defer cancel()
	_, err := i.collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}},
	})
	return err
}

func (i *Impl) GenerateRandomData(ctx context.Context, count int) error {
	documents := make([]interface{}, count)
	for i := 0; i < count; i++ {
		documents[i] = storage.Record{
			Name:      utils.GenerateRandomString(6),
			PaymentID: utils.GenerateRandomString(6),
			CreateAt:  time.Now(),
		}
	}
	_, err := i.collection.InsertMany(ctx, documents)
	return err
}

func (i *Impl) CountRecordsWithNames(ctx context.Context, names ...string) (int64, error) {
	filter := bson.D{{"name", bson.D{{"$in", names}}}}
	count, err := i.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (i *Impl) ImportRecords(ctx context.Context, records ...*storage.Record) error {
	documents := make([]interface{}, len(records))
	for i, record := range records {
		documents[i] = record
	}
	_, err := i.collection.InsertMany(ctx, documents)
	return err
}

func (i *Impl) Flush(ctx context.Context) error {
	_, err := i.collection.DeleteMany(ctx, bson.D{})
	return err
}
