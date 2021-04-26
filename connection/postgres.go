package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect(params DBParams) error {
	dns := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		params.Host,
		params.Port,
		params.User,
		params.Password,
		params.Db)
	db, err := sql.Open("postgres", dns)

	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := p.db.QueryRow("select now()").Scan(t)

	if err != nil {
		fmt.Printf("Error al leer la fecha del servidor %v", err)
		return nil, err
	}

	return t, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
