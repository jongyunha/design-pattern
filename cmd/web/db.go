package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	maxOpenDbConn = 25
	MaxIdleDBConn = 25
	maxDbLifeTime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxIdleTime(MaxIdleDBConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	return db, nil
}
