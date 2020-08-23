package main

import (
	"encoding/json"
	"os"
)

//TwitterAppKeys :
type TwitterAppKeys struct {
	Name              string
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	StreamType        string
	StreamArgs        []string
}

//Configuration :
type Configuration struct {
	TwitterAppsKeys        []TwitterAppKeys
	DBConnectionString     string
	DBName                 string
	BatchSize              int
	CrawlPageSize          int
	StopAfter              int
	IdleConns              int
	BlockedUsersPath       string
	MinimumRetweetCount    int
	MinimumGoldenTweetRate int
	LogPath                string
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
