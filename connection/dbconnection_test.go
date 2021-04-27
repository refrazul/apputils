package connection

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties"
)

func TestOracleConn(t *testing.T) {
	conn := &Oracle{}
	p := properties.MustLoadFile("oracle.properties", properties.UTF8)
	fmt.Println(p)
	db, err := conn.Connect(DBParams{Host: p.MustGetString("host"), Port: int16(p.MustGetInt("port")), User: p.MustGetString("user"), Password: p.MustGetString("password"), Db: p.MustGetString("db")})

	if err != nil {
		t.Error("Error generando conexión", err)
	}

	fmt.Println(db)
	conn.Close()
}

func TestPostgresConn(t *testing.T) {
	conn := &Postgres{}

	p := properties.MustLoadFile("postgres.properties", properties.UTF8)
	fmt.Println(p)
	db, err := conn.Connect(DBParams{Host: p.MustGetString("host"), Port: int16(p.MustGetInt("port")), User: p.MustGetString("user"), Password: p.MustGetString("password"), Db: p.MustGetString("db")})
	if err != nil {
		t.Error("no se pudo crear la conexión ", err)
	}

	fmt.Println(db)
	conn.Close()
}
