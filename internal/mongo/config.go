package mongo

import "time"

type Config struct {
	URL      string
	Database string
	Timeout  time.Duration
}
