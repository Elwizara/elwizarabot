package main

import (
	"os"

	pinterest "github.com/carrot/go-pinterest"
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

//PinterestAppKeys :
type PinterestAppKeys struct {
	Name        string
	BoardSpec   string
	AccessToken string
}

func pinterestLanguagePublisher(apiKeys *PinterestAppKeys, tweet *PublishedTweetsDTO) (*models.Pin, error) {
	client := pinterest.NewClient().RegisterAccessToken(apiKeys.AccessToken)

	image, err := os.Open(tweet.ImagePath)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	defer image.Close()

	optionals := &controllers.PinCreateOptionals{Link: tweet.PublishURL, Image: image}
	res, err := client.Pins.Create(apiKeys.BoardSpec, tweet.SoicalPublish.PublishStatus, optionals)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}

	return res, nil
}

func pinterestPublisher(lang *LanguageRate, tweet *PublishedTweetsDTO) error {

	if lang.PinterestPublishKeys != nil {
		res, err := pinterestLanguagePublisher(lang.PinterestPublishKeys, tweet)
		if err != nil {
			logs.Critical(err)
			return err
		}
		tweet.SoicalPublish.PinterestPublishID = res.Id
		logs.Infof("Publish in pinterest %v", tweet.BasePath)
	}
	return nil
}
