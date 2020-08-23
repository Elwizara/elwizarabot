package main

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/tarekbadrshalaan/anaconda"
)

func stream(db *sql.DB, stop *bool) {
	for _, apiKeys := range config.TwitterAppsKeys {
		go startStream(apiKeys, db, stop)
	}
}

func startStream(apiKeys TwitterAppKeys, db *sql.DB, stop *bool) {
	for {
		//to handel start after stop
		//will wait for 1 hour and start stream again
		func(apiKeys TwitterAppKeys, db *sql.DB, stop *bool) {

			api := anaconda.NewTwitterApiWithCredentials(apiKeys.AccessToken, apiKeys.AccessTokenSecret, apiKeys.ConsumerKey, apiKeys.ConsumerSecret)
			api.SetLogger(logs)
			stream := &anaconda.Stream{}

			if apiKeys.StreamType == "sample" {
				u := url.Values{}
				stream = api.PublicStreamSample(u)
				logs.Infof("=================== %v ================= start Streaming %v", apiKeys.Name, apiKeys.StreamType)
			} else if apiKeys.StreamType == "location" {
				u := url.Values{"locations": apiKeys.StreamArgs}
				stream = api.PublicStreamFilter(u)
				logs.Infof("=================== %v ================= start Streaming %v", apiKeys.Name, apiKeys.StreamType)
			} else {
				return
			}
			defer func() {
				logs.Criticalf("=================== %v ================= Stop streaming", apiKeys.Name)
			}()
			for t := range stream.C {
				switch tweet := t.(type) {
				case anaconda.Tweet:
					if err := savetweetToDB(&tweet, db); err != nil {
						logs.Criticalf("%v:%v", apiKeys.StreamType, err)
					}
				//case anaconda.StatusDeletionNotice:
				//handle StatusDeletionNotice:
				default:
					//handle default
					//fmt.Printf("type %T - data %v\n", t, t)
				}
				if *stop {
					stream.Stop()
				}
			}
		}(apiKeys, db, stop)

		<-time.After(time.Hour)
	}
}
