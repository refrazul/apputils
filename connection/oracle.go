package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/godror/godror"
)

type Oracle struct {
	db *sql.DB
}

func (o *Oracle) Connect(params DBParams) (*sql.DB, error) {
	if err := IsParams(&params); err != nil {
		return nil, err
	}

	dns := fmt.Sprintf(`user=%s password=%s connectString="%s:%d/%s"`,
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Db)
	db, err := sql.Open("godror", dns)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	o.db = db
	return o.db, nil
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

func (o *Oracle) GetConn() *sql.DB {
	return o.db
}
