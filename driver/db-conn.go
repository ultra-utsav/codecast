package driver

import (
	"database/sql"
	"fmt"
)

type MySQLConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DB       string
}

// GetConnection returns database connection
func GetConnection(cfg MySQLConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	db, er := sql.Open("mysql", connString)
	if er != nil {
		return nil, er
	}

	er = db.Ping()
	if er != nil {
		return nil, er
	}

	return db, nil
}
