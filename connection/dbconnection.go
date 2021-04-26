package connection

import (
	"errors"
	"time"
)

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

func IsParams(params *DBParams) error {
	if params == nil {
		return errors.New("No connection parameters")
	}
	if params.Host == "" {
		return errors.New("Invalid host")
	}
	if params.Port == 0 {
		return errors.New("Invalid port")
	}
	if params.Db == "" {
		return errors.New("Invalid db")
	}
	if params.User == "" {
		return errors.New("Invalid user")
	}
	if params.Password == "" {
		return errors.New("Invalid password")
	}

	return nil
}
