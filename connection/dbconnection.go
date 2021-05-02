package connection

import (
	"database/sql"
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

type Column struct {
	ColumnLegth     int
	ColumnName      string
	ColumnNullable  bool
	ColumnType      string
	ColumnOrder     int
	ColumnPrecision int
	ColumnDefault   string
}

type DBConnction interface {
	Connect(params DBParams) (*sql.DB, error)
	GetNow() (*time.Time, error)
	Close() error
	GetConn() *sql.DB
	TableInfo(table string) ([]Column, error)
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
