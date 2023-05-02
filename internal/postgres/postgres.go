package postgres

import (
	"context"
	"time"

	"github.com/00mohamad00/go-db-bench/pkg/storage"
	"github.com/00mohamad00/go-db-bench/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Impl struct {
	db *gorm.DB
}

func NewPostgresStorageOrPanic(conf Config) storage.Storage {
	db, err := gorm.Open(postgres.Open(conf.GetURL()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&storage.Record{})
	if err != nil {
		panic(err)
	}
	return &Impl{
		db: db,
	}
}

func (i *Impl) GenerateRandomData(ctx context.Context, count int) error {
	records := make([]storage.Record, count)
	for i := 0; i < count; i++ {
		records[i] = storage.Record{
			Name:      utils.GenerateRandomString(6),
			PaymentID: utils.GenerateRandomString(6),
			CreateAt:  time.Now(),
		}
	}
	return i.db.WithContext(ctx).CreateInBatches(&records, 1000).Error
}

func (i *Impl) CountRecordsWithNames(ctx context.Context, names ...string) (int64, error) {
	var count int64
	err := i.db.WithContext(ctx).Model(&storage.Record{}).Where("name IN ?", names).Count(&count).Error
	return count, err
}

func (i *Impl) ImportRecords(ctx context.Context, records ...*storage.Record) error {
	return i.db.WithContext(ctx).Create(&records).Error
}

func (i *Impl) Flush(ctx context.Context) error {
	i.db.WithContext(ctx).Delete(&storage.Record{})
	return nil
}
