package main

import (
	"fmt"
	"os"
	"time"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("logger")

var logFileformat = logging.MustStringFormatter(
	`%{level:.4s},%{time:15:04:05.0},%{shortfile},%{message}`,
)

var logTerminalformat = logging.MustStringFormatter(
	`%{color}%{level:.4s},%{time:15:04:05.0},%{shortfile},%{message}%{color:reset}`,
)

func setTerminalLogger(level logging.Level) *logging.LeveledBackend {
	terminalLogger := logging.NewLogBackend(os.Stderr, "", 0)
	terminalformatted := logging.NewBackendFormatter(terminalLogger, logTerminalformat)
	terminallevelled := logging.AddModuleLevel(terminalformatted)
	terminallevelled.SetLevel(level, "")
	return &terminallevelled
}

func setFileLogger(level logging.Level, filePath string) *logging.LeveledBackend {
	fileName := fmt.Sprintf(filePath)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fileLogger := logging.NewLogBackend(f, "", 0)
	fileformatted := logging.NewBackendFormatter(fileLogger, logFileformat)
	filelevelled := logging.AddModuleLevel(fileformatted)
	filelevelled.SetLevel(level, "")
	return &filelevelled
}

func setLogger(loggerType string, loggerLevel int64) {
	switch loggerType {
	case "f":
		go func() {
			logdir := config.LogPath
			createDir(&logdir)
			for {
				now := time.Now()
				year := now.Year()
				month := now.Month()
				day := now.Day()
				fileName := fmt.Sprintf("%v/%d-%d-%d.csv", logdir, year, month, day)
				fileLogger := setFileLogger(logging.Level(loggerLevel), fileName)
				logging.SetBackend(*fileLogger)
				waitFor := 24 - now.Hour()
				<-time.After(time.Duration(waitFor) * time.Hour)
			}
		}()
	case "t":
		terminalLogger := setTerminalLogger(logging.Level(loggerLevel))
		logging.SetBackend(*terminalLogger)
	case "a":
		terminalLogger := setTerminalLogger(logging.Level(loggerLevel))
		go func(terminalLogger *logging.LeveledBackend) {
			logdir := config.LogPath
			createDir(&logdir)
			for {
				now := time.Now()
				now.Add(1 * time.Hour)
				year := now.Year()
				month := now.Month()
				day := now.Day()
				fileName := fmt.Sprintf("%v/%d-%d-%d.csv", logdir, year, month, day)
				fileLogger := setFileLogger(logging.Level(loggerLevel), fileName)
				logging.SetBackend(*fileLogger, *terminalLogger)
				waitFor := 24 - now.Hour()
				<-time.After(time.Duration(waitFor) * time.Hour)
			}
		}(terminalLogger)
	case "n":
		logging.SetBackend()
	}
}
