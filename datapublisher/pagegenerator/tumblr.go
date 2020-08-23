package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	tumblr "github.com/tumblr/tumblr.go"
	//"github.com/tumblr/tumblrclient"
	tumblrclient "github.com/tumblr/tumblrclient.go"
)

//TumblrAppKeys :
type TumblrAppKeys struct {
	Name              string
	Endpoint          string
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func tumblrLanguagePublisher(apiKeys *TumblrAppKeys, tweet *PublishedTweetsDTO) (*tumblr.Response, error) {
	client := tumblrclient.NewClientWithToken(apiKeys.ConsumerKey, apiKeys.ConsumerSecret, apiKeys.AccessToken, apiKeys.AccessTokenSecret)

	data, err := ioutil.ReadFile(tweet.ImagePath)
	if os.IsNotExist(err) {
		logs.Criticalf("file not exist %v", tweet.ImagePath)
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(data)

	u := url.Values{
		"type":    []string{"photo"},
		"state":   []string{"published"},
		"link":    []string{tweet.PublishURL},
		"caption": []string{tweet.SoicalPublish.PublishStatus},
		"source":  []string{tweet.ImageURL}, "data64": []string{encoded},
	}

	res, err := client.PostWithParams(apiKeys.Endpoint, u)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	return &res, nil
}

func tumblrPublisher(lang *LanguageRate, tweet *PublishedTweetsDTO) error {

	if lang.TumblrPublishKeys != nil {
		res, err := tumblrLanguagePublisher(lang.TumblrPublishKeys, tweet)
		if err != nil {
			logs.Critical(err)
			return err
		}
		res.PopulateFromBody()
		if id, ok := res.Result["id"].(float64); ok {
			tumblrID := int64(id)
			tweet.SoicalPublish.TumblrPublishID = fmt.Sprintf("%v", tumblrID)
			logs.Infof("Publish in Tumblr %v", tweet.BasePath)
		} else {
			logs.Errorf("convert Tumblr id -res : %v", res)
		}

	}
	return nil
}
