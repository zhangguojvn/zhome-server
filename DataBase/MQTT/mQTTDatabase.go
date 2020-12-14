package MQTT

import (
	"database/sql"
	"errors"
	"zhome-server/DataBase"
)

type mQTTDatabse struct {
	db        *sql.DB
	tableName string
}

func (d *mQTTDatabse) Init(db *sql.DB) error {
	d.tableName = "MQTT"
	d.db = db
	_ = db.QueryRow("CREATE TABLE IF NOT EXISTS MQTT" +
		"(Path CHAR PRIMARY KEY," +
		"DUID CHAR," +
		"FOREIGN KEY (DUID) REFERENCES DEVICE(DUID)," +
		"MQTT CHAR)")
	return nil
}
func (d *mQTTDatabse) Get(i interface{}) (interface{}, error) {
	require, ok := i.(MQTT)
	if !ok {
		return nil, errors.New("Need MQTT")
	}
	var (
		res    *sql.Rows
		err    error
		result []MQTT
		mqtt   MQTT
	)
	if require.MQTT == "" {
		res, err = d.db.Query(
			"SELECT PATH,DUID,MQTT FROM MQTT WHERE DUID=?",
			require.DUID)
		if err != nil {
			return nil, err
		}
	} else {
		res, err = d.db.Query(
			"SELECT PATH,DUID,MQTT FROM MQTT WHERE MQTT=?",
			require.MQTT)
		if err != nil {
			return nil, err
		}
	}
	for res.Next() {
		err := res.Scan(&mqtt)
		if err != nil {
			return nil, err
		}
		result = append(result, mqtt)
	}
	return result, nil
	return nil, nil
}
func (d *mQTTDatabse) Put(i interface{}) error {
	data, ok := i.(MQTT)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("INSERT INTO MQTT(PATH,DUID,MQTT) values (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Path, data.DUID, data.MQTT)
	if err != nil {
		return err
	}
	return nil
}
func (d *mQTTDatabse) Del(i interface{}) error {
	data, ok := i.(MQTT)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("DELETE FROM MQTT WHERE DUID = ?")
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
	DataBase.GetRegisterDatabase().RegisterDatabase(new(mQTTDatabse), MQTT{})
}
