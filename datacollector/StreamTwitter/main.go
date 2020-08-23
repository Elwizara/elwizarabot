package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tarekbadrshalaan/goStuff/configuration"
	"github.com/tarekbadrshalaan/goStuff/logger"
)

//Source represent type of data source
//stream twitter 			0
//crawler 					1
//APIUserTimeline 			2
//APIUserTimelineRetweet 	21
//APIUserTimelineQuote 		22

/*Twitter State
0 Available
1 Unavailable
2 protected
3 notexist
4 hasn'tTweeted
*/

// log : singleton var used in loggin
var logs = logger.Logger

// config : singleton var hold configuration data
var config *Configuration

//BlockedUsers to save more data to them
var BlockedUsers map[int64]bool

func after(stop *bool, duration time.Duration) {
	<-time.After(duration * time.Second)
	logs.Infof("Stop After: %d sec", duration)
	*stop = true
	os.Exit(0)
}

func recovery() {
	if err := recover(); err != nil {
		logs.Critical("recovered:", err)
	}
}

var l = flag.String("l", "f", "logger{f:file|t:treminal|a:all|n:None")
var ll = flag.Int64("ll", 2, "log level{0:CRITICAL|1:ERROR|2:WARNING|3:NOTICE|4:INFO|5:DEBUG")
var c = flag.String("c", "conf.json", "configration path")

var chAnalysisTweets chan *TweetsTB

func main() {
	flag.Parse()
	config = &Configuration{}
	if err := configuration.ParseJSONConfiguration(*c, config); err != nil {
		logs.Panic(err)
	}
	if err := logger.SetLogger(*l, *ll, config.LogPath); err != nil {
		logs.Panic(err)
	}

	stop := false
	if config.StopAfter > 0 {
		go after(&stop, time.Duration(config.StopAfter))
	}

	BlockedUsers = readjsonIntMap(config.BlockedUsersPath)

	db := initDatabase(config.DBConnectionString, config.DBName)

	chAnalysisTweets = make(chan *TweetsTB, 1000)

	go stream(db, &stop)

	go runTwitterAppsOnGoodTweets(db)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C).
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	logs.Warning("Stop Streaming...")
}
