package main

import (
	"encoding/json"
	"os"
)

//Configuration :
type Configuration struct {
	Name             string
	ConnectionString string
	Table            string
	StartLoop        int64
	EndLoop          int64
	Duration         int64
	ScriptPath1      string
	ScriptPath2      string
	LogPath          string
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
