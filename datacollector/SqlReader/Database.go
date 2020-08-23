package main

import (
	"database/sql"
	"fmt"
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

//MAXUpdatedAtLimit :
func MAXUpdatedAtLimit(db *sql.DB, TableName string, limit int64) (int64, error) {
	q := fmt.Sprintf(`SELECT MAX("UpdatedAt") AS "MAXUpdatedAt" FROM (SELECT "UpdatedAt" FROM "%v" ORDER BY "UpdatedAt" ASC LIMIT %v) as t;`, TableName, limit)
	rows, err := db.Query(q)
	if err != nil {
		logger.Critical(err)
		return 0, err
	}
	defer rows.Close()
	var MAXUpdatedAt sql.NullInt64
	for rows.Next() {
		if err := rows.Scan(&MAXUpdatedAt); err != nil {
			logger.Critical(err)
			return 0, err
		}
		if MAXUpdatedAt.Valid {
			return MAXUpdatedAt.Int64, nil
		}
	}
	return 0, nil
}

//GetUsersNeedToCrawlCount :
func GetUsersNeedToCrawlCount(db *sql.DB) (int64, error) {
	q := fmt.Sprintf(`SELECT COUNT(*) FROM "UsersRateTB" WHERE "NeedToCrawl" IS TRUE;`)
	rows, err := db.Query(q)
	if err != nil {
		logger.Critical(err)
		return 0, err
	}
	defer rows.Close()
	var NeedToCrawlCount sql.NullInt64
	for rows.Next() {
		if err := rows.Scan(&NeedToCrawlCount); err != nil {
			logger.Critical(err)
			return 0, err
		}
		if NeedToCrawlCount.Valid {
			return NeedToCrawlCount.Int64, nil
		}
	}
	return 0, nil
}
