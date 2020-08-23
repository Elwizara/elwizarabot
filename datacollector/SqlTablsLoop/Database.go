package main

import (
	"database/sql"
)

func initDatabase(connectionString string, database string) *sql.DB {
	if connectionString == "" {
		logger.Error("no ConnectionString, please add db ConnectionString")
		panic("no ConnectionString, please add db ConnectionString")
	}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Critical(err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Critical(err)
		panic(err)
	}
	logger.Infof("Opened database '%v' successfully", database)
	return db
}
