package Device

import (
	"database/sql"
	"errors"
	"zhome-server/DataBase"
)

type deviceDatabse struct {
	db        *sql.DB
	tableName string
}

func (d *deviceDatabse) Init(db *sql.DB) error {
	d.tableName = "DEVICE"
	d.db = db
	_ = db.QueryRow(
		"CREATE TABLE IF NOT EXISTS DEVICE " +
			"(DUID CHAR PRIMARY KEY," +
			"Name CHAR," +
			"UIType CHAT," +
			"FOREIGN KEY (UIType) REFERENCES UIConfig(UIType)," +
			"Info CHAR)",
	)
	return nil
}
func (d *deviceDatabse) Get(i interface{}) (interface{}, error) {
	require, ok := i.(Device)
	if !ok {
		return nil, errors.New("Need Device")
	}
	res, err := d.db.Query(
		"SELECT DUID,Name,UIType,Info FROM DEVICE WHERE DUID=?",
		require.DUID)
	if err != nil {
		return nil, err
	}
	var (
		result []Device
		device Device
	)
	for res.Next() {
		err := res.Scan(&device)
		if err != nil {
			return nil, err
		}
		result = append(result, device)
	}
	return result, nil
}
func (d *deviceDatabse) Put(i interface{}) error {
	data, ok := i.(Device)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("INSERT INTO DEVICE(DUID, Name, UIType,Info) values (?, ?, ?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DUID, data.Name, data.UI.UIType, data.Info)
	if err != nil {
		return err
	}
	return nil
}
func (d *deviceDatabse) Del(i interface{}) error {
	data, ok := i.(Device)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("DELETE FROM DEVICE WHERE DUID = ?")
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
	DataBase.GetRegisterDatabase().RegisterDatabase(new(deviceDatabse), Device{})
}
