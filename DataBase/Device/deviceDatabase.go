package Device

import (
	"database/sql"
	"fmt"
)

type deviceDatabse struct {
	db *sql.DB
	tableName string
}

func (d *deviceDatabse)Init(db *sql.DB) error{
	d.tableName= "Device"
	d.db = db
	_=db.QueryRow(fmt.Sprint("CREATE TABLE IF NOT EXISTS ",
		d.tableName,
		" （DUID CHAR,Name CHAR,Info CHAR)"),
	)
	return nil
}
func (d *deviceDatabse)Get(interface{})(interface{},error) {

	return nil,nil
}
func (d *deviceDatabse)Put(interface{})error {

	return nil
}