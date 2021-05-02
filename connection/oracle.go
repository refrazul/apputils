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

func (o *Oracle) TableInfo(table string) ([]Column, error) {
	res, err := o.db.Query(`select column_name, data_type, 
	data_length, data_precision, nullable, 
	data_default, column_id 
		from ALL_TAB_COLUMNS 
		where
		table_name = :name order by column_id asc`, table)

	if err != nil {
		return []Column{}, err
	}

	var columns []Column

	for res.Next() {
		var column Column
		var default1 sql.NullString
		var precision sql.NullInt32

		err := res.Scan(&column.ColumnName, &column.ColumnType,
			&column.ColumnLegth, &precision, &column.ColumnNullable,
			&default1, &column.ColumnOrder)

		if err != nil {
			return []Column{}, err
		}

		if default1.Valid {
			column.ColumnDefault = default1.String
		}

		if precision.Valid {
			column.ColumnPrecision = int(precision.Int32)
		}

		columns = append(columns, column)
	}

	return columns, nil
}
