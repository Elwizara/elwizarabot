package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/tarekbadrshalaan/anaconda"
)

func runTwitterAppsOnGoodTweets(db *sql.DB) {
	for _, apiKeys := range config.TwitterAppsKeys {
		go func(apiKeys TwitterAppKeys) {
			logs.Infof("=================== %v ================= start collect Good Tweet", apiKeys.Name)
			defer func() {
				logs.Criticalf("=================== %v ================= Stop collect Good Tweet", apiKeys.Name)
			}()
			api := anaconda.NewTwitterApiWithCredentials(apiKeys.AccessToken, apiKeys.AccessTokenSecret, apiKeys.ConsumerKey, apiKeys.ConsumerSecret)
			api.SetLogger(anaconda.BasicLogger)
			for tweet := range chAnalysisTweets {
				if err := analysisUserTweet(db, api, tweet); err != nil {
					logs.Critical(err)
				}
			}
		}(apiKeys)
	}
}

func analysisUserTweet(db *sql.DB, api *anaconda.TwitterApi, tweet *TweetsTB) error {
	defer recovery()

	// check user in database.
	userRate, err := GetUserRateDB(db, tweet.UserID)
	if err != nil {
		logs.Critical(err)
		return err
	}
	// calculate User Rate from Twitter
	if userRate == nil {
		userRate = &UserRate{
			UserID:   tweet.UserID,
			UserName: tweet.UserName,
		}
		if userRate, err = GetUserRateTwitter(db, api, userRate); err != nil {
			logs.Critical(err)
			return err
		}
	}
	if userRate == nil {
		err := fmt.Errorf("No User Found:%v:%d", tweet.UserName, tweet.UserID)
		logs.Critical(err)
		return err
	}

	if userRate.CollectedTweetsCount > 100 {
		measureRate := (tweet.Favorite + tweet.Reply + (tweet.Retweet * 2) + 1) / userRate.GoldenFactor
		if measureRate >= config.MinimumGoldenTweetRate { // golden tweet
			//review tweet finally
			gTweetDTO, err := api.GetTweet(tweet.TweetID, nil)
			if err != nil {
				logs.Critical(err)
				return err
			}
			gTweetTB, err := convertAnacondaTweetToTweetTB(&gTweetDTO, false)
			if err != nil {
				logs.Critical(err)
				return err
			}
			finalyGoldentweet := gTweetTB.toGoldenTweetsTB(userRate.primaryLanguage, userRate.GoldenFactor)
			//review tweet finally

			logs.Infof("Golden Tweet : https://twitter.com/%v/status/%d %v %v", finalyGoldentweet.UserName, finalyGoldentweet.TweetID, finalyGoldentweet.MeasureRate, finalyGoldentweet.TweetType.Quote)
			if err := saveGoldenTweetsToDB(finalyGoldentweet, db); err != nil {
				logs.Critical(err)
				return err
			}
		}
	}
	return nil
}

// GetUserRateTwitter : use twitter api to get user timeline and calculate rate and save to db
func GetUserRateTwitter(db *sql.DB, api *anaconda.TwitterApi, userRate *UserRate) (*UserRate, error) {
	defer recovery()

	// GetUserTimeline
	userTimeline, err := GetUserTimeline(*api, userRate.UserID, 500)

	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	userRate.CollectedTweetsCount = len(userTimeline)

	// calculateUserRate
	userRate.Languages = make(map[string]int)
	for _, tw := range userTimeline {
		userRate.LikesTweetsCount += tw.FavoriteCount
		userRate.ReplyTweetsCount += tw.ReplyCount
		userRate.RetweetedTweetsCount += tw.RetweetCount
		userRate.Languages[tw.Lang]++
		// get primary languages.
		if userRate.primaryLanguageCount < userRate.Languages[tw.Lang] {
			userRate.primaryLanguageCount = userRate.Languages[tw.Lang]
			userRate.primaryLanguage = tw.Lang
		}
	}

	userRate.GoldenFactor = ((userRate.LikesTweetsCount+userRate.ReplyTweetsCount+(userRate.RetweetedTweetsCount*2))/userRate.CollectedTweetsCount + 1) + 1

	// measurementUserRateWithTweet
	languagesString, err := json.Marshal(userRate.Languages)
	if err != nil {
		logs.Critical(err)
	}
	userRate.LanguagesStr = fmt.Sprintf("%s", languagesString)
	if err := saveUserRateToDB(userRate, db); err != nil {
		logs.Critical(err)
		return nil, err
	}
	return userRate, nil
}

//GetUserTimeline :
func GetUserTimeline(api anaconda.TwitterApi, userID int64, TweetsCount int) ([]anaconda.Tweet, error) {
	uid := fmt.Sprint(userID)
	u := url.Values{
		"user_id":         []string{uid},
		"count":           []string{"200"},
		"exclude_replies": []string{"true"},
		"include_rts":     []string{"false"},
	}
	userTimeline, err := api.GetUserTimeline(u)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}

	userTimelineCount := len(userTimeline)
	if userTimelineCount < 1 {
		addblockedUser(userID)
		err := fmt.Errorf("userTimelineCount < 1 :%d", userID)
		logs.Critical(err)
		return nil, err
	}
	for i := 0; i < 10; i++ {
		maxID := userTimeline[userTimelineCount-1].IdStr
		u = url.Values{
			"user_id":         []string{uid},
			"count":           []string{"200"},
			"exclude_replies": []string{"true"},
			"include_rts":     []string{"false"},
			"max_id":          []string{maxID},
		}
		res, err := api.GetUserTimeline(u)
		if err != nil {
			logs.Critical(err)
			return userTimeline, err
		}
		userTimeline = append(userTimeline, res...)
		if len(userTimeline) >= TweetsCount || userTimelineCount == len(userTimeline) {
			return userTimeline, nil
		}
		userTimelineCount = len(userTimeline)
	}
	return userTimeline, nil
}

func addblockedUser(UserID int64) {
	BlockedUsers[UserID] = true
}

func checkblockedUser(UserID int64) bool {
	return BlockedUsers[UserID]
}
