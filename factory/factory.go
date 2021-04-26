package factory

import "github.com/refrazul/apputils/connection"

func FactoryConn(connType string) connection.DBConnction {

	switch connType {
	case "Oracle":
		return &connection.Oracle{}
	case "Postgres":
		return &connection.Postgres{}
	default:
		return nil
	}

}
