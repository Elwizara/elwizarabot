package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	iconv "github.com/djimenez/iconv-go"
	"github.com/tarekbadrshalaan/anaconda"
)

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

// UserProfileTB :
type UserProfileTB struct {
	UserID         int64
	UserName       string
	ViewName       string
	Bio            sql.NullString
	Location       sql.NullString
	Link           sql.NullString
	JoinDate       int64
	TweetsCount    int64
	LikesCount     int
	FollowingCount int
	FollowersCount int
	BirthDate      string
	UpdatedAt      int64
	LastTweetDate  *int64
	LastTweetID    *int64
	TwitterState   int
	LastSource     int
}

// TweetsTB :
type TweetsTB struct {
	TweetID          int64
	UserID           int64
	Favorite         int
	Reply            int
	Retweet          int
	SilverFactor     int
	UserName         string
	CreatedAt        int64
	UpdatedAt        int64
	Lang             string
	CountryCode      *string
	RetweetedTweetID *int64
	RetweetedUserID  *int64
	QuotedTweetID    *int64
	QuotedUserID     *int64
	ReplyTweetID     *int64
	ReplyUserID      *int64
	Source           int
	Text             string
	User             *UserProfileTB
	TweetType        TweetType
}

func (tweet *TweetsTB) toGoldenTweetsTB(userPrimaryLanguage string, userGoldenFactor int) *GoldenTweetsTB {
	tweetSilverFactor := (tweet.Favorite + tweet.Reply + (tweet.Retweet * 2) + 1)
	measureRate := tweetSilverFactor / userGoldenFactor
	return &GoldenTweetsTB{
		TweetID:             tweet.TweetID,
		UserID:              tweet.UserID,
		UserName:            tweet.UserName,
		Favorite:            tweet.Favorite,
		Reply:               tweet.Reply,
		Retweet:             tweet.Retweet,
		TweetLanguage:       tweet.Lang,
		UserPrimaryLanguage: userPrimaryLanguage,
		UserGoldenFactor:    userGoldenFactor,
		MeasureRate:         measureRate,
		CreatedAt:           tweet.CreatedAt,
		UpdatedAt:           time.Now().Unix(),
		Expired:             false,
		TweetType:           tweet.TweetType,
		TweetSilverFactor:   tweetSilverFactor,
	}
}

func getSilverFactor(t *anaconda.Tweet) int {
	return ((t.FavoriteCount + t.ReplyCount) * t.RetweetCount)
}

func stringtoUnixTime(date string) int64 {
	if res, err := time.Parse(time.RubyDate, date); err == nil {
		return res.Unix()
	}
	return 0
}

func encodeWindows1250(inp string) string {
	output, _ := iconv.ConvertString(inp, "utf-8", "utf-8")
	return output
}

func convertAnacondaTweetToTweetTB(tweetDTO *anaconda.Tweet, isLastUserTweet bool) (*TweetsTB, error) {
	defer recovery()
	tweetTB := &TweetsTB{}
	tweetTB.TweetType = TweetType{
		PossiblySensitive: tweetDTO.PossiblySensitive,
	}
	//is video or images
	if &tweetDTO.ExtendedEntities != nil && len(tweetDTO.ExtendedEntities.Media) > 0 {
		if tweetDTO.ExtendedEntities.Media[0].Type == "photo" {
			tweetTB.TweetType.Photo = true
		}
		if tweetDTO.ExtendedEntities.Media[0].Type == "video" {
			tweetTB.TweetType.Video = true
		}
	}

	tweetTB.TweetID = tweetDTO.Id
	tweetTB.UserID = tweetDTO.User.Id
	tweetTB.Favorite = tweetDTO.FavoriteCount
	tweetTB.Reply = tweetDTO.ReplyCount
	tweetTB.Retweet = tweetDTO.RetweetCount
	tweetTB.Retweet += tweetDTO.QuoteCount
	tweetTB.SilverFactor = getSilverFactor(tweetDTO)
	tweetTB.UserName = tweetDTO.User.ScreenName
	//tweet time
	tweetTime, err := tweetDTO.CreatedAtTime()
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	tweetTB.CreatedAt = tweetTime.Unix()
	tweetTB.UpdatedAt = time.Now().Unix() //updated at time now
	tweetTB.Lang = tweetDTO.Lang
	//country code
	if &tweetDTO.Place != nil && len(tweetDTO.Place.CountryCode) > 1 {
		tweetTB.CountryCode = &tweetDTO.Place.CountryCode
		tweetTB.TweetType.HasLocation = true
	}

	if tweetDTO.RetweetedStatus != nil {
		tweetTB.RetweetedTweetID = &tweetDTO.RetweetedStatus.Id
		tweetTB.RetweetedUserID = &tweetDTO.RetweetedStatus.User.Id
		tweetTB.TweetType.Retweete = true
	}

	if tweetDTO.QuotedStatus != nil {
		tweetTB.QuotedTweetID = &tweetDTO.QuotedStatus.Id
		tweetTB.QuotedUserID = &tweetDTO.QuotedStatus.User.Id
		tweetTB.TweetType.Quote = true
	}

	if tweetDTO.InReplyToStatusID != 0 {
		tweetTB.ReplyTweetID = &tweetDTO.InReplyToStatusID
		tweetTB.ReplyUserID = &tweetDTO.InReplyToUserID
		tweetTB.TweetType.Reply = true
	}
	tweetTB.Source = 0
	tweetTB.Text = encodeWindows1250(tweetDTO.FullText)

	tweetTB.User = &UserProfileTB{
		UserID:         tweetDTO.User.Id,
		UserName:       tweetDTO.User.ScreenName,
		ViewName:       tweetDTO.User.Name,
		JoinDate:       stringtoUnixTime(tweetDTO.User.CreatedAt),
		TweetsCount:    tweetDTO.User.StatusesCount,
		LikesCount:     tweetDTO.User.FavouritesCount,
		FollowingCount: tweetDTO.User.FriendsCount,
		FollowersCount: tweetDTO.User.FollowersCount,
		UpdatedAt:      time.Now().Unix(),
	}

	if len(tweetDTO.User.Description) > 2 {
		tweetTB.User.Bio = sql.NullString{String: tweetDTO.User.Description, Valid: true}
	}

	if len(tweetDTO.User.Location) > 2 {
		tweetTB.User.Location = sql.NullString{String: tweetDTO.User.Location, Valid: true}
	}

	if len(tweetDTO.User.URL) > 2 {
		tweetTB.User.Link = sql.NullString{String: tweetDTO.User.URL, Valid: true}
	}

	if isLastUserTweet {
		tweetTB.User.LastTweetID = &tweetTB.TweetID
		tweetTB.User.LastTweetDate = &tweetTB.CreatedAt
	}

	tweetTB.User.TwitterState = 0 //Available
	if tweetDTO.User.Protected {
		tweetTB.User.TwitterState = 2 //protected
	}

	tweetTB.User.LastSource = 0

	return tweetTB, nil
}

// UserRate hold data calculated about user.
type UserRate struct {
	UserID               int64
	UserName             string
	CollectedTweetsCount int
	LikesTweetsCount     int
	ReplyTweetsCount     int
	RetweetedTweetsCount int
	GoldenFactor         int
	Languages            map[string]int
	LanguagesStr         string
	primaryLanguage      string
	primaryLanguageCount int
	UpdatedAt            int64
}

// GoldenTweetsTB summary tweet hold data we need to analysis.
type GoldenTweetsTB struct {
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
