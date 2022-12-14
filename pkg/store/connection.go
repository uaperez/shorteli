package store

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type connection struct {
	Username string
	Password string
	Host     string
	Schema   string
}

func GetConnection(connectionType string) (*sql.DB, error) {
	var settings connection
	var connectionString string
	switch connectionType {
	case "LOCAL":
		settings.Username = "root"
		settings.Password = ""
		settings.Host = "localhost"
		settings.Schema = "shorteli"
		connectionString = fmt.Sprintf("%s:%s/%s", settings.Username, settings.Password, settings.Schema)
	default:
		return nil, errors.New("Configuración incorrecta para la base de datos")
	}
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
