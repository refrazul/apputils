package connection

import "time"

type DBParams struct {
	Host     string
	User     string
	Password string
	Port     int16
	Db       string
}

type DBConnction interface {
	Connect(params DBParams) error
	GetNow() (*time.Time, error)
	Close() error
}
