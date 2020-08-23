package main

import (
	"bytes"
	"time"

	"github.com/tarekbadrshalaan/goStuff/jsonparser"
	"github.com/tarekbadrshalaan/graw/reddit"
)

// RedditAppKeys :
type RedditAppKeys struct {
	Name      string
	Agent     string
	AppID     string
	AppSecret string
	Username  string
	Password  string
	Subreddit string
}

func redditLanguagePublisher(apiKeys *RedditAppKeys, tweet *PublishedTweetsDTO) (string, error) {
	config := reddit.BotConfig{
		Agent: apiKeys.Agent,
		App: reddit.App{
			ID:       apiKeys.AppID,
			Secret:   apiKeys.AppSecret,
			Username: apiKeys.Username,
			Password: apiKeys.Password,
		},
		Rate: 5 * time.Second,
	}

	bot, err := reddit.NewBot(config)
	if err != nil {
		logs.Critical(err)
		return "", err
	}
	//res, err := bot.PostLink(apiKeys.Subreddit, "Awesome Tweet", tweet.ImageURL)
	res, err := bot.PostLink(apiKeys.Subreddit, tweet.SoicalPublish.PublishStatus, tweet.ImageURL)
	if err != nil {
		logs.Critical(err)
		return "", err
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		logs.Critical(err)
		return "", err
	}
	body := buf.Bytes()

	//response :
	//{"json": {"errors": [], "data": {"url": "https://www.reddit.com/r/u_elwizaracom/comments/9ewfg2/awesome_tweet/", "drafts_count": 0, "id": "9ewfg2", "name": "t3_9ewfg2"}}}
	result, err := jsonparser.JSONParser(body, "json", "data")
	if err != nil {
		logs.Criticalf("body : %s , err %v", body, err)
		return "", err
	}
	postID := result["name"].(string)

	_, err = bot.Reply(postID, tweet.SoicalPublish.PublishStatus)
	if err != nil {
		logs.Critical(err)
		return postID, err
	}
	return postID, nil
}

func redditPublisher(lang *LanguageRate, tweet *PublishedTweetsDTO) error {

	if lang.RedditPublishKeys != nil {
		postID, err := redditLanguagePublisher(lang.RedditPublishKeys, tweet)
		if err != nil {
			logs.Critical(err)
			return err
		}
		tweet.SoicalPublish.RedditPublishID = postID
		logs.Infof("Publish in reddit %v", tweet.BasePath)
	}
	return nil
}
