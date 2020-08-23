package main

import (
	"fmt"

	"github.com/huandu/facebook"
	"github.com/tarekbadrshalaan/goStuff/jsonparser"
)

// FacebookAppKeys :
type FacebookAppKeys struct {
	Name        string
	AccessToken string
	PageID      string
}

func facebookLanguagePublisher(apiKeys *FacebookAppKeys, tweet *PublishedTweetsDTO) (string, error) {
	res, err := facebook.Batch(facebook.Params{
		"access_token": apiKeys.AccessToken,
		"file1":        facebook.File(tweet.ImagePath),
	}, facebook.Params{
		"method":         "POST",
		"relative_url":   apiKeys.PageID + "/photos",
		"body":           fmt.Sprintf("message=%v", tweet.SoicalPublish.PublishStatus),
		"attached_files": "file1",
	})
	if err != nil {
		logs.Critical(err)
		return "", err
	}
	id, err := jsonparser.Getkeystring(res[0], "body", "post_id")
	if err != nil {
		logs.Critical(err)
		return "", err
	}
	return id, nil
}

func facebookPublisher(lang *LanguageRate, tweet *PublishedTweetsDTO) error {
	if lang.FacebookPublishKeys != nil {
		postID, err := facebookLanguagePublisher(lang.FacebookPublishKeys, tweet)
		if err != nil {
			logs.Critical(err)
			return err
		}
		tweet.SoicalPublish.FacebookPublishID = postID
		logs.Infof("Publish in Facebook %v", tweet.BasePath)
	}
	return nil
}
