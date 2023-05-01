package storage

import "time"

type Record struct {
	Name      string
	PaymentID string
	CreateAt  time.Time
}
