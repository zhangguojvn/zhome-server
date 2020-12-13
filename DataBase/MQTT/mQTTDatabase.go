package MQTT

import "database/sql"

type mQTTDatabase struct {
	db *sql.DB
	tableName string
}

