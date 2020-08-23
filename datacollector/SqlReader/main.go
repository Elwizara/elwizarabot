package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/tarekbadrshalaan/goStuff/configuration"
)

var config *Configuration

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
	//go spinner(100 * time.Millisecond) //loader

	flag.Parse()
	loggerType := *l
	loggerLevel := *ll
	config = &Configuration{}
	err := configuration.ParseJSONConfiguration(*c, config)
	if err != nil {
		logger.Criticalf("error Parse configuration : %v", err)
	}
	setLogger(loggerType, loggerLevel)

	for _, source := range config.DBSources {
		go StartDownloadSource(config, source)
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stop Downloading...")
}
