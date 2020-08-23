package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tarekbadrshalaan/goStuff/configuration"
)

// StartGeneratePages :
func StartGeneratePages(db *sql.DB) {
	languages := &[]LanguageRate{}
	if err := configuration.ParseJSONConfiguration(config.LanguageRatePath, languages); err != nil {
		logs.Panic(err)
	}
	for _, lang := range *languages {
		go languageGeneratePages(db, lang)
		go languageSocialPublisher(db, lang)
		<-time.After(time.Minute * time.Duration(config.DurationBetweenLangs))
	}
}

func waitAppropriateTime(operation, lang string, lastTime int64, duration time.Duration) {
	timenowunix := time.Now().Unix()
	//if last time + lang duration bigger then now , need to wait until appropriate time
	subtractionTime := (lastTime + int64(duration*60)) - timenowunix
	if subtractionTime > 0 {
		logs.Warningf("%v-%v  lastcreated:%v|timenow:%v|willwaitfor:%v", lang, operation, lastTime, timenowunix, subtractionTime)
		<-time.After(time.Second * time.Duration(subtractionTime))
	}
}

func languageGeneratePages(db *sql.DB, lang LanguageRate) {
	lastCreatedTime, err := GetTimeLastLanguagePublishedTweetsDB(db, lang.Lang)
	if err != nil {
		logs.Critical(err)
	}
	duration := time.Duration(lang.Duration)
	waitAppropriateTime("GeneratePages", lang.Lang, lastCreatedTime, duration)
	for {
		githubPushWG.Wait()
		generatePageWG.Add(1)
		if err := publishGoldenTweet(db, &lang); err != nil {
			logs.Critical(err)
		}
		generatePageWG.Done()
		<-time.After(time.Minute * duration)
	}
}

func publishGoldenTweet(db *sql.DB, lang *LanguageRate) error {
	defer recovery()
	tweet, err := GetLangGoldenTweetsDB(db, lang)
	if err != nil {
		logs.Critical(err)
		return err
	}
	if tweet == nil {
		return nil
	}

	relatedTweets, err := GetLatestLangPublishedTweetsDB(db, tweet.UserPrimaryLanguage, 10)
	if err != nil {
		logs.Critical(err)
		return err
	}
	tw := tweet.ConvToPublishedTweetsDTO(relatedTweets)

	if err = createDir(tw.DumpDirectory); err != nil {
		logs.Critical(err)
		return err
	}

	if err = createDir(tw.JSONDirectory); err != nil {
		logs.Critical(err)
		return err
	}

	if err = generateTweetpages(db, tw); err != nil {
		logs.Critical(err)
		return err
	}
	return nil
}

func generateTweetpages(db *sql.DB, tweet *PublishedTweetsDTO) error {
	defer recovery()

	if err := logPublishedTweetsDB(db, tweet); err != nil {
		logs.Critical(err)
		return err
	}

	// dump generator
	if err := generatepage(tweet, tweet.DumpPath, config.Template.DumpPath); err != nil {
		logs.Critical(err)
		return err
	}
	//create image
	if err := generateTweetImage(db, tweet); err != nil {
		logs.Critical(err)
		return err
	}

	if err := logImageGeneratedDB(db, tweet); err != nil {
		logs.Critical(err)
		return err
	}
	//create image

	//
	//generate json
	pubJSON, err := tweet.JSON()
	if err != nil {
		logs.Critical(err)
		return err
	}
	if err = saveJSON(tweet.JSONPath, pubJSON); err != nil {
		logs.Critical(err)
		return err
	}
	//save general file
	if err = addTweetToLanugageFile(tweet); err != nil {
		logs.Critical(err)
		return err
	}
	err = logJSONGeneratedDB(db, tweet)
	if err != nil {
		logs.Critical(err)
		return err
	}
	//generate json

	logs.Infof("Generate Page %v", tweet.BasePath)
	return nil
}

func languageSocialPublisher(db *sql.DB, lang LanguageRate) {
	/*
		lastCreatedTime, err := GetTimeLastSocialPublisheDB(db, lang.Lang)
		if err != nil {
			logs.Critical(err)
		}
		duration := time.Duration(lang.Duration)
		waitAppropriateTime("SocialPublisher", lang.Lang, lastCreatedTime, duration)
	*/
	duration := time.Duration(config.CommitDuration)
	for {
		githubPushWG.Wait()
		if err := publishInSocial(db, &lang); err != nil {
			logs.Critical(err)
		}
		<-time.After(time.Minute * duration)
	}
}

func publishInSocial(db *sql.DB, lang *LanguageRate) error {
	defer recovery()
	// get tweets published in website and github.
	tweet, err := GetLastLanguagePublishedTweetsDB(db, lang.Lang)
	if err != nil {
		logs.Critical(err)
		return nil
	}
	if tweet == nil {
		return nil
	}
	// get tweets published in website and github.
	url := tweet.PublishURL
	if config.AllowShortURL {
		url = tweet.PublishShortURL
	}
	if tweet.TweetType.Video {
		tweet.SoicalPublish.PublishStatus = fmt.Sprintf("you can watch the video here :-\n%v", url)
	} else {
		tweet.SoicalPublish.PublishStatus = fmt.Sprintf("you can find more awesome tweets here:-\n%v", url)
	}

	//publish in reddit
	if err = redditPublisher(lang, tweet); err != nil {
		logs.Critical(err)
	}
	//publish in reddit

	//publish in pinterest
	if err = pinterestPublisher(lang, tweet); err != nil {
		logs.Critical(err)
	}
	//publish in pinterest

	//publish in tumblr
	if err = tumblrPublisher(lang, tweet); err != nil {
		logs.Critical(err)
	}
	//publish in tumblr

	//publish in twitter
	if err = twitterPublisher(lang, tweet); err != nil {
		logs.Critical(err)
	}
	//publish in twitter

	//publish in facebook
	if err = facebookPublisher(lang, tweet); err != nil {
		logs.Critical(err)
	}
	//publish in facebook

	if err = logSoicalPublishDB(db, tweet); err != nil {
		logs.Critical(err)
		return err
	}

	return nil
}
