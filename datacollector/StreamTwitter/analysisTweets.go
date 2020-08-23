package main

import (
	"time"
)

func analysisTweet(tweet *TweetsTB) {

	dateNow := time.Now().Unix()
	userJoinDateFromNow := dateNow - tweet.User.JoinDate
	tweetDataFromNow := dateNow - tweet.CreatedAt

	/*
		Tweet becomes good when :-
		- RetweetCount > MinimumRetweetCount.
		- Tweet created recently with in two days | Three days.	// in unix  2 day = 172800 & 3 day = 259200
		- User joins date older then year.			// 1 year in unix 	= 31556926.
		- User tweets less then 200000
	*/

	if tweet.Retweet >= config.MinimumRetweetCount &&
		tweetDataFromNow <= 259200 &&
		userJoinDateFromNow >= 31556926 &&
		tweet.User.TweetsCount <= 200000 {
		chAnalysisTweets <- tweet
		/*
			chGoodTweets <- &GoldenTweetsTB{
				TweetID:             tweet.TweetID,
				UserID:              tweet.UserID,
				UserName:            tweet.UserName,
				Favorite:            tweet.Favorite,
				Reply:               tweet.Reply,
				Retweet:             tweet.Retweet,
				TweetLanguage:       tweet.Lang,
				UserPrimaryLanguage: "",
				UserGoldenFactor:    0,
				MeasureRate:         0,
				CreatedAt:           tweet.CreatedAt,
				UpdatedAt:           dateNow,
				Expired:             false,
				TweetType:           tweet.TweetType,
				TweetSilverFactor:   (tweet.Favorite + tweet.Reply + (tweet.Retweet * 2) + 1),
			}
		*/
	}
}
