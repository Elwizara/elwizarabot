package main

import (
	"encoding/json"
	"os"
)

//Configuration :
type Configuration struct {
	DBName             string
	DBConnectionString string
	LogPath            string
	ResultPath         string
}

func parseConfiguration(config string) *Configuration {
	var confi *Configuration

	file, err := os.Open(config)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&confi)
	file.Close()
	if err != nil {
		file.Close()
		panic(err)
	}
	return confi
}
