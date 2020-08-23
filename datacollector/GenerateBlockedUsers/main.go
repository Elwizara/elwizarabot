package main

import (
	"flag"
	"fmt"
	"time"
)

var configuration *Configuration

var l = flag.String("l", "a", "logger{f:file|t:treminal|a:all|n:None")
var ll = flag.Int64("ll", 4, "log level{0:CRITICAL|1:ERROR|2:WARNING|3:NOTICE|4:INFO|5:DEBUG")
var c = flag.String("c", "conf.json", "configration path")

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond) //loader
	flag.Parse()
	loggerType := *l
	loggerLevel := *ll
	configuration = parseConfiguration(*c)
	setLogger(loggerType, loggerLevel)

	db := initDatabase(configuration.DBConnectionString, configuration.DBName)

	GenerateBlockedUsers(db)
}
