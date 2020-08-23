package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func initDatabase(connectionString string, database string) *sql.DB {
	if connectionString == "" {
		logs.Panic("no ConnectionString, please add db ConnectionString")
	}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logs.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		logs.Panic(err)
	}
	logs.Infof("Opened database '%v' successfully", database)
	return db
}

// GetLatestLangPublishedTweetsDB :
func GetLatestLangPublishedTweetsDB(db *sql.DB, lang string, limit int) (RelatedTweets, error) {
	defer recovery()

	q := fmt.Sprintf(`	SELECT "TweetId","UserId","UserName"
						FROM "PublishedTweetsTB"
						WHERE "Lang"= '%v' AND "IsDeleted" = FALSE
						ORDER BY "CreatedAt" DESC
						LIMIT %d`, lang, limit)
	rows, err := db.Query(q)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	defer rows.Close()

	relatedTweets := RelatedTweets{}
	for rows.Next() {
		relatedTweet := RelatedTweet{}
		if err := rows.Scan(
			&relatedTweet.TweetID,
			&relatedTweet.UserID,
			&relatedTweet.UserName); err != nil {
			logs.Critical(err)
			return nil, err
		}
		relatedTweets = append(relatedTweets, relatedTweet)
	}
	return relatedTweets, nil

}

// GetLangGoldenTweetsDB :
func GetLangGoldenTweetsDB(db *sql.DB, lang *LanguageRate) (*GoldenTweet, error) {
	defer recovery()

	mintime := time.Now().Unix() - 259200 //3 day = 259200

	q := fmt.Sprintf(`SELECT * FROM "GoldenTweetsTB"
						WHERE 
							"UserPrimaryLanguage" = "TweetLanguage" 
							AND "UserPrimaryLanguage" = '%v' 
							AND "MeasureRate" >= %d
							AND "CreatedAt" >= %d
							AND "Expired" = FALSE
							AND "Type"->>'Reply' = 'false'
							AND "Type"->>'Quote' = 'false'
							AND "Type"->>'PossiblySensitive' = 'false'
							ORDER BY "MeasureRate" DESC
						LIMIT 1`, lang.Lang, lang.MinimumRate, mintime)
	rows, err := db.Query(q)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		goldenTweet := &GoldenTweet{}
		var tweetType []byte
		if err := rows.Scan(
			&goldenTweet.TweetID,
			&goldenTweet.UserID,
			&goldenTweet.UserName,
			&goldenTweet.Favorite,
			&goldenTweet.Reply,
			&goldenTweet.Retweet,
			&goldenTweet.TweetLanguage,
			&goldenTweet.UserPrimaryLanguage,
			&goldenTweet.TweetSilverFactor,
			&goldenTweet.UserGoldenFactor,
			&goldenTweet.MeasureRate,
			&goldenTweet.CreatedAt,
			&goldenTweet.UpdatedAt,
			&goldenTweet.Expired,
			&tweetType); err != nil {
			logs.Critical(err)
			return nil, err
		}

		if err = json.Unmarshal(tweetType, &goldenTweet.TweetType); err != nil {
			logs.Critical(err)
			return nil, err
		}
		return goldenTweet, nil
	}
	return nil, nil
}

// logPublishedTweetsDB : insert into PublishedTweetsTB and update GoldenTweetsTB
func logPublishedTweetsDB(db *sql.DB, tweet *PublishedTweetsDTO) error {
	defer recovery()

	query := fmt.Sprintf(`SELECT public."InsertPublishedTweetsTB"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16);`)
	rows, err := db.Query(query,
		tweet.TweetID,
		tweet.UserID,
		tweet.UserName,
		tweet.TweetCreatedAt,
		tweet.Lang,
		tweet.TweetType.String(),
		tweet.BasePath,
		config.Template.Version,
		false,           //jsoncreated
		false,           //imagecreated
		false,           //gitpushed
		tweet.CreatedAt, //created at
		tweet.UpdatedAt, //updated at
		false,           //IsDeleted
		tweet.RelatedTweets.String(),
		tweet.BaseShortURL)
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

// GetTimeLastLanguagePublishedTweetsDB : Get time for Last tweets publishe for Language from "PublishedTweetsTB"
func GetTimeLastLanguagePublishedTweetsDB(db *sql.DB, lang string) (int64, error) {
	defer recovery()
	query := fmt.Sprintf(`	SELECT "CreatedAt" 
							FROM "PublishedTweetsTB"
							WHERE "Lang"= '%v'
							ORDER BY "CreatedAt" DESC
							LIMIT 1`, lang)

	rows, err := db.Query(query)
	if err != nil {
		logs.Critical(err)
		return 0, err
	}
	defer rows.Close()

	var lastCreatedTime int64
	for rows.Next() {
		if err := rows.Scan(&lastCreatedTime); err != nil {
			logs.Critical(err)
			return 0, err
		}
	}
	return lastCreatedTime, nil
}

// GetTimeLastSocialPublisheDB : Get time for Last tweets publishe for Language from "PublishedTweetsTB"
func GetTimeLastSocialPublisheDB(db *sql.DB, lang string) (int64, error) {
	defer recovery()
	query := fmt.Sprintf(`	SELECT "UpdatedAt" 
							FROM "PublishedTweetsTB"
							WHERE "Lang"= '%v'
							AND "SoicalPublish" IS NOT NULL 
							ORDER BY "UpdatedAt" DESC
							LIMIT 1`, lang)

	rows, err := db.Query(query)
	if err != nil {
		logs.Critical(err)
		return 0, err
	}
	defer rows.Close()

	var lastCreatedTime int64
	for rows.Next() {
		if err := rows.Scan(&lastCreatedTime); err != nil {
			logs.Critical(err)
			return 0, err
		}
	}
	return lastCreatedTime, nil
}

// logJSONGeneratedDB : update PublishedTweetsTB set IsJSONCreated = true
func logJSONGeneratedDB(db *sql.DB, tweet *PublishedTweetsDTO) error {
	defer recovery()
	timenow := time.Now().Unix()
	query := fmt.Sprintf(`Update "PublishedTweetsTB" SET "IsJSONCreated" = TRUE , "UpdatedAt"=%d WHERE "TweetId" = %d`, timenow, tweet.TweetID)
	rows, err := db.Query(query)
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

// logImageGeneratedDB : update PublishedTweetsTB set IsImageCreated = true
func logImageGeneratedDB(db *sql.DB, tweet *PublishedTweetsDTO) error {
	defer recovery()
	timenow := time.Now().Unix()
	query := fmt.Sprintf(`Update "PublishedTweetsTB" SET "IsImageCreated" = TRUE , "UpdatedAt"=%d WHERE "TweetId" = %d`, timenow, tweet.TweetID)
	rows, err := db.Query(query)
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

// logGitPushedDB : update PublishedTweetsTB set IsGitPushed = true
func logGitPushedDB(db *sql.DB) error {
	defer recovery()
	timenow := time.Now().Unix()
	query := fmt.Sprintf(`	Update "PublishedTweetsTB" SET "IsGitPushed" = TRUE , "UpdatedAt"=%d 
							WHERE "IsJSONCreated" = TRUE AND "IsImageCreated" = TRUE AND "IsGitPushed" = FALSE;`, timenow)
	rows, err := db.Query(query)
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

// logPublishTwitterDB : update PublishedTweetsTB set TwitterPublishId = newtweetId
func logSoicalPublishDB(db *sql.DB, tweet *PublishedTweetsDTO) error {
	defer recovery()
	timenow := time.Now().Unix()
	query := fmt.Sprintf(`Update "PublishedTweetsTB" SET "SoicalPublish" ='%v', "UpdatedAt"=%d WHERE "TweetId" = %d`, tweet.SoicalPublish.String(), timenow, tweet.TweetID)
	rows, err := db.Query(query)
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

// GetLastLanguagePublishedTweetsDB :get tweets published in website and github and not publish in social.
func GetLastLanguagePublishedTweetsDB(db *sql.DB, lang string) (*PublishedTweetsDTO, error) {
	defer recovery()

	q := fmt.Sprintf(`SELECT * FROM "PublishedTweetsTB"
						WHERE 
							"Lang" = '%v'
							AND "IsJSONCreated" = TRUE
							AND "IsImageCreated" = TRUE
							AND "IsGitPushed" = TRUE
							AND "SoicalPublish" IS NULL 
							ORDER BY "CreatedAt" DESC
						LIMIT 1`, lang)
	rows, err := db.Query(q)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		publishTweet := &PublishedTweetsTB{}
		var SoicalPublish sql.NullString
		var BaseShortURL sql.NullString
		if err := rows.Scan(
			&publishTweet.TweetID,
			&publishTweet.UserID,
			&publishTweet.UserName,
			&publishTweet.TweetCreatedAt,
			&publishTweet.Lang,
			&publishTweet.TweetType,
			&publishTweet.PagePublishPath,
			&publishTweet.TemplateVersion,
			&publishTweet.IsJSONCreated,
			&publishTweet.IsImageCreated,
			&publishTweet.IsGitPushed,
			&publishTweet.CreatedAt,
			&publishTweet.UpdatedAt,
			&publishTweet.IsDeleted,
			&publishTweet.RelatedTweets,
			&SoicalPublish,
			&BaseShortURL); err != nil {
			logs.Critical(err)
			return nil, err
		}
		if SoicalPublish.Valid {
			publishTweet.SoicalPublish = SoicalPublish.String
		}
		if BaseShortURL.Valid {
			publishTweet.BaseShortURL = BaseShortURL.String
		}
		res, err := publishTweet.ConvToPublishedTweetsDTO()
		if err != nil {
			logs.Critical(err)
			return nil, err
		}
		return res, nil
	}
	return nil, nil
}
