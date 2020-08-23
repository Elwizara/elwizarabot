package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tarekbadrshalaan/goStuff/configuration"
	"github.com/tarekbadrshalaan/goStuff/logger"
)

// log : singleton var used in loggin
var logs = logger.Logger

// config : singleton var hold configuration data
var config *Configuration

// waitgroup sync github push goroutine with page generator
var githubPushWG sync.WaitGroup

// generatePageWG sync generate pages goroutine with github push
var generatePageWG sync.WaitGroup

var l = flag.String("l", "a", "logger{f:file|t:treminal|a:all|n:None")
var ll = flag.Int64("ll", 4, "log level{0:CRITICAL|1:ERROR|2:WARNING|3:NOTICE|4:INFO|5:DEBUG")
var c = flag.String("c", "conf.json", "configration path")

func recovery() {
	if err := recover(); err != nil {
		logs.Critical("recovered:", err)
	}
}

func main() {
	flag.Parse()

	config = &Configuration{}
	if err := configuration.ParseJSONConfiguration(*c, config); err != nil {
		logs.Panic(err)
	}
	if err := logger.SetLogger(*l, *ll, config.LogPath); err != nil {
		logs.Panic(err)
	}
	db := initDatabase(config.DBConnectionString, config.DBName)

	go StartGeneratePages(db)

	go githubcommitter(db)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stop Generateing...")
}

// go build . && cp pagegenerator ../test/Generator/testpage
