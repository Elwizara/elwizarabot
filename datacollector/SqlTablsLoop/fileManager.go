package main

import (
	"os"
)

func createDir(path *string) {
	if _, errIsNotExist := os.Stat(*path); os.IsNotExist(errIsNotExist) {
		err := os.MkdirAll(*path, os.ModePerm)
		if err != nil {
			logger.Critical(err)
			return
		}
		logger.Infof("createDir(path %v)", *path)
	}
}
