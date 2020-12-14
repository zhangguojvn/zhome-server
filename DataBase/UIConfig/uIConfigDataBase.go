package UIConfig

import (
	"database/sql"
	"errors"
	"zhome-server/DataBase"
)

type UIConfigDatabase struct {
	db        *sql.DB
	tableName string
}

func (d *UIConfigDatabase) Init(db *sql.DB) error {
	d.tableName = "UICONFIG"
	d.db = db
	_ = db.QueryRow("CREATE TABLE IF NOT EXISTS UICONFIG " +
		"(UIType CHAR PRIMARY KEY," +
		"UIConfig JSON, " +
		"CHECK (JSON_VALID(UIConfig)))")
	return nil
}
func (d *UIConfigDatabase) Get(i interface{}) (interface{}, error) {
	require, ok := i.(UIConfig)
	if !ok {
		return nil, errors.New("Need UIConfig")
	}
	var (
		res      *sql.Rows
		err      error
		result   []UIConfig
		UIConfig UIConfig
	)
	res, err = d.db.Query(
		"SELECT UIType,UIConfig FROM UICONFIG WHERE UITYPE=?",
		require.UIType)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err := res.Scan(&UIConfig)
		if err != nil {
			return nil, err
		}
		result = append(result, UIConfig)
	}
	return result, nil
}
func (d *UIConfigDatabase) Put(i interface{}) error {
	data, ok := i.(UIConfig)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("INSERT INTO UICONFIG(UIType,UIConfig) values (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.UIType, data.UIConfig)
	if err != nil {
		return err
	}
	return nil
}
func (d *UIConfigDatabase) Del(i interface{}) error {
	data, ok := i.(UIConfig)
	if !ok {
		return errors.New("Need Device")
	}
	stmt, err := d.db.Prepare("DELETE FROM UICONFIG WHERE UIType=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.UIType)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	DataBase.GetRegisterDatabase().RegisterDatabase(new(UIConfigDatabase), UIConfig{})
}
