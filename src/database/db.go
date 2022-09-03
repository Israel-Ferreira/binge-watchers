package database

import (
	"database/sql"
	"fmt"

	"github.com/Israel-Ferreira/binge-watchers/src/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	confg := config.DbConfigVars

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		confg.Host,
		confg.Port,
		confg.Username,
		confg.Password,
		confg.DbName,
	)

	connection, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}
