package TSDB

import (
	"database/sql"
	"errors"
	"zhome-server/DataBase"
)

type tSDBDatabase struct {
	db        *sql.DB
	tableName string
}

func (d *tSDBDatabase) Init(db *sql.DB) error {
	d.tableName = "TSDB"
	d.db = db
	_ = db.QueryRow("CREATE TABLE IF NOT EXISTS TSDB (DUID CHAR PRIMARY KEY,TSDB CHAR)")
	return nil
}
func (d *tSDBDatabase) Get(i interface{}) (interface{}, error) {
	require, ok := i.(TSDB)
	if !ok {
		return nil, errors.New("Need TSDB")
	}
	var (
		res    *sql.Rows
		err    error
		result []TSDB
		TSDB   TSDB
	)
	res, err = d.db.Query(
		"SELECT DUID,TSDB FROM TSDB WHERE DUID=?",
		require.DUID)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err := res.Scan(&TSDB)
		if err != nil {
			return nil, err
		}
		result = append(result, TSDB)
	}
	return result, nil
}
func (d *tSDBDatabase) Put(i interface{}) error {
	data, ok := i.(TSDB)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("INSERT INTO TSDB(DUID,TSDB) values ( ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DUID, data.TSDB)
	if err != nil {
		return err
	}
	return nil
}
func (d *tSDBDatabase) Del(i interface{}) error {
	data, ok := i.(TSDB)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("DELETE FROM TSDB WHERE  DUID =?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DUID)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	DataBase.GetRegisterDatabase().RegisterDatabase(new(tSDBDatabase), TSDB{})
}
