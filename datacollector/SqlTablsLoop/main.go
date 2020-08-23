package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var configuration *Configuration

var l = flag.String("l", "a", "logger{f:file|t:treminal|a:all|n:None")
var ll = flag.Int64("ll", 4, "log level{0:CRITICAL|1:ERROR|2:WARNING|3:NOTICE|4:INFO|5:DEBUG")
var c = flag.String("c", "conf.json", "configration path")

func doScript(configuration *Configuration, index int64) error {
	output, err := exec.Command(configuration.ScriptPath1, configuration.ConnectionString, fmt.Sprint(index)).Output()
	if err != nil || !strings.HasPrefix(fmt.Sprintf("%s", output), "COMMIT") {
		logger.Criticalf("%v:%v:%v Script1 Command failed:'%v'  output:%s", configuration.Name, configuration.Table, index, err, output)
		return err
	}
	logger.Infof("%v:%v:%v Script1 Successfull", configuration.Name, configuration.Table, index)
	if configuration.ScriptPath2 != "" {
		duration := time.Duration(configuration.Duration)
		<-time.After(duration * time.Second)
		output, err = exec.Command(configuration.ScriptPath2, configuration.ConnectionString, fmt.Sprint(index)).Output()
		if err != nil || !strings.HasPrefix(fmt.Sprintf("%s", output), "COMMIT") {
			logger.Criticalf("%v:%v:%v Script2 Command failed:'%v'  output:%s", configuration.Name, configuration.Table, index, err, output)
			return err
		}
		logger.Infof("%v:%v:%v Script2 Successfull", configuration.Name, configuration.Table, index)
	}
	return nil
}

func main() {
	flag.Parse()
	loggerType := *l
	loggerLevel := *ll
	configuration = parseConfiguration(*c)
	setLogger(loggerType, loggerLevel)

	duration := time.Duration(configuration.Duration)
	for i := configuration.StartLoop; i < configuration.EndLoop; i++ {
		if err := doScript(configuration, i); err != nil {
			if err != nil {
				logger.Criticalf("doScript Error %v", err)
				continue
			}
		}
		<-time.After(duration * time.Second)
	}

}
