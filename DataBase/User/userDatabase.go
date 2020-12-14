package User

import (
	"database/sql"
	"errors"
	"zhome-server/DataBase"
)

type userDatabase struct {
	db        *sql.DB
	tableName string
}

func (d *userDatabase) Init(db *sql.DB) error {
	d.tableName = "USER"
	d.db = db
	_ = db.QueryRow("CREATE TABLE IF NOT EXISTS USER " +
		"(UUID CHAR PRIMARY KEY ," +
		"Name CHAR," +
		"Info CHAR)")
	return nil
}
func (d *userDatabase) Get(i interface{}) (interface{}, error) {
	require, ok := i.(User)
	if !ok {
		return nil, errors.New("Need User")
	}
	var (
		res    *sql.Rows
		err    error
		result []User
		User   User
	)
	res, err = d.db.Query(
		"SELECT UUID,Name,Info FROM USER WHERE UUID=?",
		require.UUID)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err := res.Scan(&User)
		if err != nil {
			return nil, err
		}
		result = append(result, User)
	}
	return result, nil
}
func (d *userDatabase) Put(i interface{}) error {
	data, ok := i.(User)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("INSERT INTO USER(UUID,Name,Info) values (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.UUID, data.Name, data.Info)
	if err != nil {
		return err
	}
	return nil
}
func (d *userDatabase) Del(i interface{}) error {
	data, ok := i.(User)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("DELETE FROM USER WHERE UUID = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.UUID)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	DataBase.GetRegisterDatabase().RegisterDatabase(new(userDatabase), User{})
}
