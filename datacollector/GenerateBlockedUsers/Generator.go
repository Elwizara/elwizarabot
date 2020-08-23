package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//GenerateBlockedUsers :
func GenerateBlockedUsers(db *sql.DB) error {
	resultdir := configuration.ResultPath
	createDir(resultdir)
	resultdir = fmt.Sprintf("%v/BlockedUsers.json", resultdir)
	createFile(resultdir)

	var UsersID []int64
	for i := 0; i < 1000; i++ {
		tableName := fmt.Sprintf("TweetsTB_%d", i)
		res, err := GetUsersIDs(db, tableName, 1000)
		if err != nil {
			logger.Critical(err)
			return err
		}
		UsersID = append(UsersID, res...)
		logger.Infof("%v Done", tableName)
	}

	data, err := json.Marshal(UsersID)
	if err != nil {
		logger.Critical(err)
		return err
	}

	if err := writeFile(resultdir, data); err != nil {
		logger.Critical(err)
		return err
	}
	return nil
}
