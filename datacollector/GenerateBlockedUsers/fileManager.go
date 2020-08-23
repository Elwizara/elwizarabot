package main

import (
	"os"
)

func createFile(path string) error {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			logger.Critical(err)
			return err
		}
		defer file.Close()
		logger.Infof("createFile(path %v)", path)
	}
	return nil
}

func createDir(path string) {
	if _, errIsNotExist := os.Stat(path); os.IsNotExist(errIsNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Critical(err)
			return
		}
		logger.Infof("createDir(path %v)", path)
	}
}

func writeFile(path string, data []byte) error {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		logger.Critical(err)
		return err
	}
	defer file.Close()

	// write data file
	_, err = file.Write(data)
	if err != nil {
		logger.Critical(err)
		return err
	}

	// save changes
	err = file.Sync()
	if err != nil {
		logger.Critical(err)
		return err
	}

	logger.Infof("writeFile(path %v)", path)
	return nil
}
