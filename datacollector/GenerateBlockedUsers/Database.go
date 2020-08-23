package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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

//GetUsersIDs : get users ids from database which have tweet more then max count
func GetUsersIDs(db *sql.DB, TableName string, maxcount int) ([]int64, error) {
	q := fmt.Sprintf(`
		SELECT "UserId" FROM (
			SELECT "UserId",COUNT("UserId") AS "mycount"
			FROM "%v"
			GROUP by "UserId"
		) AS t
		WHERE "mycount" > %d
	`, TableName, maxcount)
	rows, err := db.Query(q)
	if err != nil {
		logger.Critical(err)
		return nil, err
	}
	defer rows.Close()
	var UsersID []int64
	for rows.Next() {
		var u int64
		if err := rows.Scan(&u); err != nil {
			logger.Critical(err)
			return nil, err
		}
		if u != 0 {
			UsersID = append(UsersID, u)
		}
	}
	return UsersID, nil
}
