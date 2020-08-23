package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"os"
	"time"

	"github.com/tarekbadrshalaan/anaconda"
)

//TwitterAppKeys :
type TwitterAppKeys struct {
	Name              string
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func twitterUploadMedia(api *anaconda.TwitterApi, tweet *PublishedTweetsDTO) (string, error) {
	data, err := ioutil.ReadFile(tweet.ImagePath)
	if os.IsNotExist(err) {
		logs.Criticalf("file not exist %v", tweet.ImagePath)
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	// when we need some shit we do shit.
	api.HttpClient.Timeout = (time.Minute * 5)

	res, err := api.UploadMedia(encoded)
	if err != nil {
		logs.Critical(err)
		return "", err
	}
	return res.MediaIDString, nil
}
func twitterLanguagePublisher(apiKeys *TwitterAppKeys, tweet *PublishedTweetsDTO) (*anaconda.Tweet, error) {

	api := anaconda.NewTwitterApiWithCredentials(apiKeys.AccessToken, apiKeys.AccessTokenSecret, apiKeys.ConsumerKey, apiKeys.ConsumerSecret)
	api.SetLogger(logs)
	imageID, err := twitterUploadMedia(api, tweet)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}

	u := url.Values{"media_ids": []string{imageID}}
	res, err := api.PostTweet(tweet.SoicalPublish.PublishStatus, u)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	return &res, nil

}

func twitterPublisher(lang *LanguageRate, tweet *PublishedTweetsDTO) error {
	if lang.TwitterPublishKeys != nil {
		res, err := twitterLanguagePublisher(lang.TwitterPublishKeys, tweet)
		if err != nil {
			logs.Critical(err)
			return err
		}
		tweet.SoicalPublish.TwitterPublishID = res.IdStr
		logs.Infof("Publish in Twitter %v", tweet.BasePath)
	}
	return nil
}
