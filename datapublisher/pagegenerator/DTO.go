package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tarekbadrshalaan/goStuff/numbercompression"
)

// covertTweetURL :
func covertTweetURL(UserName string, TweetID int64) string {
	return fmt.Sprintf("https://twitter.com/%v/status/%d", UserName, TweetID)
}

// TweetType :
type TweetType struct {
	Retweete          bool
	Quote             bool
	Reply             bool
	HasLocation       bool
	Photo             bool
	Video             bool
	PossiblySensitive bool
}

func (t *TweetType) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		logs.Criticalf("JSON marshaling failed: %s", err)
	}
	return fmt.Sprintf("%s", data)
}

// SoicalPublish : all fileds related to soical media
type SoicalPublish struct {
	PublishStatus      string `json:"PublishStatus"`
	TwitterPublishID   string `json:"TwitterPublishId"`
	TumblrPublishID    string `json:"TumblrPublishId"`
	PinterestPublishID string `json:"PinterestPublishId"`
	RedditPublishID    string `json:"RedditPublishId"`
	FacebookPublishID  string `json:"FacebookPublishId"`
}

func (t SoicalPublish) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		logs.Criticalf("JSON marshaling failed: %s", err)
	}
	return fmt.Sprintf("%s", data)
}

//GoldenTweet : collection of data within one tweet page
type GoldenTweet struct {
	TweetID             int64
	UserID              int64
	UserName            string
	Favorite            int
	Reply               int
	Retweet             int
	TweetLanguage       string
	UserPrimaryLanguage string
	TweetSilverFactor   int
	UserGoldenFactor    int
	MeasureRate         int
	CreatedAt           int64
	UpdatedAt           int64
	Expired             bool
	TweetType           TweetType
}

// RelatedTweet :
type RelatedTweet struct {
	TweetID  string `json:"Id"`
	UserID   string `json:"UserId"`
	UserName string `json:"UserName"`
}

// PublishedTweetsTB : collection of field from PublishedTweetsTB table
type PublishedTweetsTB struct {
	TweetID         int64
	UserID          int64
	UserName        string
	TweetCreatedAt  int64
	Lang            string
	TweetType       string
	PagePublishPath string
	BaseShortURL    string
	TemplateVersion string
	IsJSONCreated   bool
	IsImageCreated  bool
	IsGitPushed     bool
	CreatedAt       int64
	UpdatedAt       int64
	IsDeleted       bool
	RelatedTweets   string
	SoicalPublish   string
}

// ConvToPublishedTweetsDTO :
func (t *PublishedTweetsTB) ConvToPublishedTweetsDTO() (*PublishedTweetsDTO, error) {
	tweetMonth := time.Unix(t.CreatedAt, 0).Month()
	tweetYear := time.Unix(t.CreatedAt, 0).Year()
	BaseDirectory := fmt.Sprintf("%v/%d/%d", t.Lang, tweetYear, tweetMonth)
	BasePath := t.PagePublishPath
	pub := &PublishedTweetsDTO{
		TweetID:                    t.TweetID,
		UserID:                     t.UserID,
		UserName:                   t.UserName,
		Lang:                       t.Lang,
		TweetCreatedAt:             t.TweetCreatedAt,
		CreatedAt:                  t.CreatedAt,
		UpdatedAt:                  t.UpdatedAt,
		TweetURL:                   covertTweetURL(t.UserName, t.TweetID),
		BaseDirectory:              BaseDirectory,
		BasePath:                   BasePath,
		LanguageDirectory:          fmt.Sprintf("%v/%v/%v", config.GeneratedDataPath, config.JSONDirectory, t.Lang),
		LanguageRelativeDirectory:  fmt.Sprintf("%v/%v", config.JSONDirectory, t.Lang),
		LanguageCollectionFilePath: fmt.Sprintf("%v/%v/%v/%v", config.GeneratedDataPath, config.JSONDirectory, t.Lang, config.CollectionFileName),
		DumpDirectory:              fmt.Sprintf("%v/%v/%v", config.GeneratedDataPath, config.DumpDirectory, BaseDirectory),
		DumpPath:                   fmt.Sprintf("%v/%v/%v.html", config.GeneratedDataPath, config.DumpDirectory, BasePath),
		DumpRelativePath:           fmt.Sprintf("%v/%v.html", config.DumpDirectory, BasePath),
		JSONDirectory:              fmt.Sprintf("%v/%v/%v", config.GeneratedDataPath, config.JSONDirectory, BaseDirectory),
		JSONPath:                   fmt.Sprintf("%v/%v/%v.json", config.GeneratedDataPath, config.JSONDirectory, BasePath),
		ImagePath:                  fmt.Sprintf("%v/%v/%v.png", config.GeneratedDataPath, config.JSONDirectory, BasePath),
		PublishURL:                 fmt.Sprintf(config.ElwizaraPostURL, t.Lang, tweetYear, tweetMonth, t.TweetID),
		ImageURL:                   fmt.Sprintf(config.ElwizaraImageURL, t.Lang, tweetYear, tweetMonth, t.TweetID),
		PublishShortURL:            fmt.Sprintf(config.ElwizaraShortURL, t.BaseShortURL),
	}
	if t.TweetType != "" {
		err := json.Unmarshal([]byte(t.TweetType), &pub.TweetType)
		if err != nil {
			logs.Criticalf("JSON marshaling failed: %s", err)
			return nil, err
		}
	}
	if t.RelatedTweets != "" {
		err := json.Unmarshal([]byte(t.RelatedTweets), &pub.RelatedTweets)
		if err != nil {
			logs.Criticalf("JSON marshaling failed: %s", err)
			return nil, err
		}
	}
	if t.SoicalPublish != "" {
		err := json.Unmarshal([]byte(t.SoicalPublish), &pub.SoicalPublish)
		if err != nil {
			logs.Criticalf("JSON marshaling failed: %s", err)
			return nil, err
		}
	}
	if len(t.BaseShortURL) < 2 {
		pub.addShortURL()
	}
	return pub, nil
}

// RelatedTweets : array of related tweets
type RelatedTweets []RelatedTweet

func (t RelatedTweets) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		logs.Criticalf("JSON marshaling failed: %s", err)
	}
	return fmt.Sprintf("%s", data)
}

// PublishedTweetsDTO : collection of field need to publish
type PublishedTweetsDTO struct {
	TweetID                    int64
	UserID                     int64
	UserName                   string
	Lang                       string
	TweetCreatedAt             int64
	CreatedAt                  int64
	UpdatedAt                  int64
	TweetURL                   string
	LanguageDirectory          string
	LanguageRelativeDirectory  string
	LanguageCollectionFilePath string
	BaseDirectory              string
	BasePath                   string
	DumpDirectory              string
	DumpPath                   string
	DumpRelativePath           string
	JSONDirectory              string
	JSONPath                   string
	ImagePath                  string
	PublishURL                 string
	ImageURL                   string
	BaseShortURL               string
	PublishShortURL            string

	TweetType     TweetType
	RelatedTweets RelatedTweets
	SoicalPublish SoicalPublish
}

// ConvToPublishedTweetsDTO :
func (t *GoldenTweet) ConvToPublishedTweetsDTO(RelatedTweets RelatedTweets) *PublishedTweetsDTO {
	timeNow := time.Now()

	BaseDirectory := fmt.Sprintf("%v/%d/%d", t.UserPrimaryLanguage, timeNow.Year(), timeNow.Month())
	BasePath := fmt.Sprintf("%v/%d", BaseDirectory, t.TweetID)
	pub := &PublishedTweetsDTO{
		TweetID:                    t.TweetID,
		UserID:                     t.UserID,
		UserName:                   t.UserName,
		Lang:                       t.UserPrimaryLanguage,
		TweetCreatedAt:             t.CreatedAt,
		CreatedAt:                  timeNow.Unix(),
		UpdatedAt:                  timeNow.Unix(),
		TweetURL:                   covertTweetURL(t.UserName, t.TweetID),
		BaseDirectory:              BaseDirectory,
		BasePath:                   BasePath,
		LanguageDirectory:          fmt.Sprintf("%v/%v/%v", config.GeneratedDataPath, config.JSONDirectory, t.UserPrimaryLanguage),
		LanguageRelativeDirectory:  fmt.Sprintf("%v/%v", config.JSONDirectory, t.UserPrimaryLanguage),
		LanguageCollectionFilePath: fmt.Sprintf("%v/%v/%v/%v", config.GeneratedDataPath, config.JSONDirectory, t.UserPrimaryLanguage, config.CollectionFileName),
		DumpDirectory:              fmt.Sprintf("%v/%v/%v", config.GeneratedDataPath, config.DumpDirectory, BaseDirectory),
		DumpPath:                   fmt.Sprintf("%v/%v/%v.html", config.GeneratedDataPath, config.DumpDirectory, BasePath),
		DumpRelativePath:           fmt.Sprintf("%v/%v.html", config.DumpDirectory, BasePath),
		JSONDirectory:              fmt.Sprintf("%v/%v/%v", config.GeneratedDataPath, config.JSONDirectory, BaseDirectory),
		JSONPath:                   fmt.Sprintf("%v/%v/%v.json", config.GeneratedDataPath, config.JSONDirectory, BasePath),
		ImagePath:                  fmt.Sprintf("%v/%v/%v.png", config.GeneratedDataPath, config.JSONDirectory, BasePath),
		PublishURL:                 fmt.Sprintf(config.ElwizaraPostURL, t.UserPrimaryLanguage, timeNow.Year(), timeNow.Month(), t.TweetID),
		ImageURL:                   fmt.Sprintf(config.ElwizaraImageURL, t.UserPrimaryLanguage, timeNow.Year(), timeNow.Month(), t.TweetID),
		TweetType:                  t.TweetType,
	}
	if RelatedTweets != nil {
		pub.RelatedTweets = RelatedTweets
	}
	pub.addShortURL()

	return pub
}

func (t *PublishedTweetsDTO) addShortURL() {
	str := t.Lang
	tweetMonth := time.Unix(t.CreatedAt, 0).Month()
	tweetYear := time.Unix(t.CreatedAt, 0).Year()
	date := (((int64(tweetYear) % 2000) * 100) + int64(tweetMonth))
	str += numbercompression.CompresNumberDefault(date)
	str += numbercompression.CompresNumberDefault(t.TweetID)
	t.BaseShortURL = str
	t.PublishShortURL = fmt.Sprintf(config.ElwizaraShortURL, str)
}

type tweetFileSimpleDTO struct {
	Nextpageid string           `json:"Nextpageid"`
	TweetsList []tweetSimpleDTO `json:"TweetsList"`
}

func (t tweetFileSimpleDTO) JSON() ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		logs.Criticalf("JSON marshaling failed: %s", err)
		return nil, err
	}
	return data, nil
}

type tweetSimpleDTO struct {
	ID       string `json:"Id"`
	UserID   string `json:"UserId"`
	UserName string `json:"UserName"`
}

func (t PublishedTweetsDTO) convTotweetSimpleDTO() *tweetSimpleDTO {
	tw := &tweetSimpleDTO{
		ID:       fmt.Sprint(t.TweetID),
		UserID:   fmt.Sprint(t.UserID),
		UserName: t.UserName,
	}
	return tw
}

// JSON : get simple json (ID ,UserID ,UserName ,Recommended)
func (t PublishedTweetsDTO) JSON() ([]byte, error) {
	type pubtw struct {
		ID          string        `json:"Id"`
		UserID      string        `json:"UserId"`
		UserName    string        `json:"UserName"`
		Recommended RelatedTweets `json:"Recommended"`
	}

	tw := pubtw{
		ID:          fmt.Sprint(t.TweetID),
		UserID:      fmt.Sprint(t.UserID),
		UserName:    t.UserName,
		Recommended: t.RelatedTweets,
	}

	data, err := json.Marshal(tw)
	if err != nil {
		logs.Criticalf("JSON marshaling failed: %s", err)
		return nil, err
	}
	return data, nil
}

// LanguageRate : collection of fields tell how to deal with language
type LanguageRate struct {
	Lang                 string
	Duration             int
	MinimumRate          int
	TwitterPublishKeys   *TwitterAppKeys
	TumblrPublishKeys    *TumblrAppKeys
	PinterestPublishKeys *PinterestAppKeys
	RedditPublishKeys    *RedditAppKeys
	FacebookPublishKeys  *FacebookAppKeys
}
