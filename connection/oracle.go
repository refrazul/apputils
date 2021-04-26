package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "gopkg.in/goracle.v2"
)

type Oracle struct {
	db *sql.DB
}

func (o *Oracle) Connect(params DBParams) error {
	dns := fmt.Sprintf("%s/%s@%s:%d/%s",
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Db)
	db, err := sql.Open("goracle", dns)

	if err != nil {
		return err
	}

	o.db = db
	return nil
}

func (o *Oracle) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := o.db.QueryRow("select sysdate from dual").Scan(t)

	if err != nil {
		fmt.Printf("Error al leer la fecha del servidor %v", err)
		return nil, err
	}

	return t, nil
}

func (o *Oracle) Close() error {
	return o.db.Close()
}
