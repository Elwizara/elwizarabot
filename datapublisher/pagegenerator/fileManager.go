package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

func createDir(path string) error {
	if _, errIsNotExist := os.Stat(path); os.IsNotExist(errIsNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logs.Critical(err)
			return err
		}
		logs.Infof("createDir(path %v)", path)
	}
	return nil
}

func generatepage(data interface{}, filePath, TemplatePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		logs.Critical(err)
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.ParseFiles(TemplatePath))

	if err = tmpl.Execute(f, data); err != nil {
		logs.Critical(err)
		return err
	}
	return nil
}

func readJSON(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		logs.Warningf("file not exist %v", path)
		return nil
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, v)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func saveJSON(filePath string, data []byte) error {
	//http://permissions-calculator.org/decode/0666/
	//0666 anyone can : READ,WRITE
	err := ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		logs.Critical(err)
		return err
	}
	return nil
}

func addTweetToLanugageFile(tweet *PublishedTweetsDTO) error {
	twfile := &tweetFileSimpleDTO{}
	err := readJSON(tweet.LanguageCollectionFilePath, twfile)
	if err != nil {
		logs.Critical(err)
		return err
	}

	if len(twfile.TweetsList) >= 100 {
		unixNow := time.Now().Unix()
		newpath := fmt.Sprintf("%v/%d.json", tweet.LanguageDirectory, unixNow)
		err := os.Rename(tweet.LanguageCollectionFilePath, newpath)
		if err != nil {
			logs.Critical(err)
			return err
		}
		twfile.Nextpageid = fmt.Sprintf("%v/%d.json", tweet.LanguageRelativeDirectory, unixNow)
		twfile.TweetsList = nil
	}
	// prepend new tweet at first
	twfile.TweetsList = append([]tweetSimpleDTO{*tweet.convTotweetSimpleDTO()}, twfile.TweetsList...)

	data, err := twfile.JSON()
	if err != nil {
		logs.Critical(err)
		return err
	}

	err = saveJSON(tweet.LanguageCollectionFilePath, data)
	if err != nil {
		logs.Critical(err)
		return err
	}
	return nil
}
