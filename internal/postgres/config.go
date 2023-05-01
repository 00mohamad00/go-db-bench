package postgres

import "fmt"

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (config *Config) GetURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
	)
}
