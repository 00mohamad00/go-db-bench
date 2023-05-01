package storage

import "context"

type Storage interface {
	GenerateRandomData(ctx context.Context, count int) error
	ImportRecords(ctx context.Context, records ...*Record) error
	CountRecordsWithNames(ctx context.Context, names ...string) (int64, error)
	Flush(ctx context.Context) error
}
