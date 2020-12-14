package Permission

import (
	"database/sql"
	"errors"
	"zhome-server/DataBase"
)

type permissionDatabase struct {
	db        *sql.DB
	tableName string
}

func (d *permissionDatabase) Init(db *sql.DB) error {
	d.tableName = "PERMISSION"
	d.db = db
	_ = db.QueryRow("CREATE TABLE IF NOT EXISTS PERMISSION (" +
		"UUID CHAR,DUID CHAR," +
		"Permission INTEGER," +
		"PRIMARY KEY (UUID,DUID)," +
		"FOREIGN KEY (DUID) REFERENCES DEVICE(DUID)," +
		"FOREIGN KEY (UUID) REFERENCES USER(UUID))")
	return nil
}
func (d *permissionDatabase) Get(i interface{}) (interface{}, error) {
	require, ok := i.(Permission)
	if !ok {
		return nil, errors.New("Need Permission")
	}
	var (
		res        *sql.Rows
		err        error
		result     []Permission
		Permission Permission
	)
	res, err = d.db.Query(
		"SELECT UUID,DUID,Permission FROM PERMISSION WHERE UUID=? AND DUID=?",
		require.UUID, require.DUID)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err := res.Scan(&Permission)
		if err != nil {
			return nil, err
		}
		result = append(result, Permission)
	}
	return result, nil
}
func (d *permissionDatabase) Put(i interface{}) error {
	data, ok := i.(Permission)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("INSERT INTO PERMISSION(UUID,DUID,Permission) values (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.UUID, data.DUID, data.Permission)
	if err != nil {
		return err
	}
	return nil
}
func (d *permissionDatabase) Del(i interface{}) error {
	data, ok := i.(Permission)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("DELETE FROM PERMISSION WHERE UUID = ? AND DUID =?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.UUID, data.DUID)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	DataBase.GetRegisterDatabase().RegisterDatabase(new(permissionDatabase), Permission{})
}
