package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/tarekbadrshalaan/anaconda"
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

func savetweetToDB(tweetDTO *anaconda.Tweet, db *sql.DB) error {
	defer recovery()

	if err := insertTweetsTB(tweetDTO, true, db); err != nil {
		logs.Critical(err)
		return err
	}

	if tweetDTO.RetweetedStatus != nil {
		tw := tweetDTO.RetweetedStatus
		if err := insertTweetsTB(tw, false, db); err != nil {
			logs.Critical(err)
			return err
		}
	}

	if tweetDTO.QuotedStatus != nil {
		tw := tweetDTO.QuotedStatus
		if err := insertTweetsTB(tw, false, db); err != nil {
			logs.Critical(err)
			return err
		}
	}

	return nil
}

func insertTweetsTB(tweetDTO *anaconda.Tweet, isLastUserTweet bool, db *sql.DB) error {
	defer recovery()
	if checkblockedUser(tweetDTO.User.Id) {
		return nil
	}
	tweetTB, err := convertAnacondaTweetToTweetTB(tweetDTO, isLastUserTweet)
	if err != nil {
		logs.Critical(err)
		return err
	}
	analysisTweet(tweetTB)
	query := fmt.Sprintf(`SELECT public."InsertTweetsTB"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20)`)
	rows, err := db.Query(query, tweetTB.TweetID,
		tweetTB.UserID,
		tweetTB.Favorite,
		tweetTB.Reply,
		tweetTB.Retweet,
		tweetTB.SilverFactor,
		tweetTB.UserName,
		tweetTB.CreatedAt,
		tweetTB.UpdatedAt, //updated at time now
		tweetTB.Lang,
		tweetTB.CountryCode,
		tweetTB.RetweetedTweetID,
		tweetTB.RetweetedUserID,
		tweetTB.QuotedTweetID,
		tweetTB.QuotedUserID,
		tweetTB.ReplyTweetID,
		tweetTB.ReplyUserID,
		tweetTB.Source,
		tweetTB.Text,
		tweetTB.TweetType.String())
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	if err = saveUserToDB(db, tweetTB.User); err != nil {
		logs.Critical(err)
		return err
	}
	return nil
}

func saveUserToDB(db *sql.DB, user *UserProfileTB) error {
	defer recovery()
	query := fmt.Sprintf(`SELECT public."InsertUsersProfileTB"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17);`)
	rows, err := db.Query(query, user.UserID,
		user.UserName,
		user.ViewName,
		user.Bio,
		user.Location,
		user.Link,
		user.JoinDate,
		user.TweetsCount,
		user.LikesCount,
		user.FollowingCount,
		user.FollowersCount,
		nil, //BirthDate
		user.UpdatedAt,
		user.LastTweetDate,
		user.LastTweetID,
		user.TwitterState,
		user.LastSource)
	if err != nil {
		logs.Critical(err)
		return nil
	}
	rows.Close()
	return nil
}

//GetUserRateDB :
func GetUserRateDB(db *sql.DB, UserID int64) (*UserRate, error) {
	defer recovery()
	q := fmt.Sprintf(`
		SELECT 
			"UserId", "UserName", "CollectedTweetsCount", "LikesTweetsCount", 
			"ReplyTweetsCount", "RetweetedTweetsCount", "Languages",
			"primaryLanguage", "primaryLanguageCount", "GoldenFactor", "UpdatedAt"
		FROM public."UsersRateTB"  
		WHERE "UserId" = %d;`, UserID)
	rows, err := db.Query(q)
	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	defer rows.Close()
	userRate := &UserRate{}

	for rows.Next() {
		if err := rows.Scan(
			&userRate.UserID,
			&userRate.UserName,
			&userRate.CollectedTweetsCount,
			&userRate.LikesTweetsCount,
			&userRate.ReplyTweetsCount,
			&userRate.RetweetedTweetsCount,
			&userRate.LanguagesStr,
			&userRate.primaryLanguage,
			&userRate.primaryLanguageCount,
			&userRate.GoldenFactor,
			&userRate.UpdatedAt); err != nil {
			if err != nil {
				logs.Critical(err)
				return nil, err
			}
			err := json.Unmarshal([]byte(userRate.LanguagesStr), &userRate.Languages)
			if err != nil {
				logs.Critical(err)
				return userRate, err
			}
		}
		return userRate, nil
	}

	return nil, nil
}

func saveUserRateToDB(userRate *UserRate, db *sql.DB) error {
	defer recovery()
	query := fmt.Sprintf(`SELECT public."InsertUsersRateTB"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`)
	rows, err := db.Query(query,
		userRate.UserID,
		userRate.UserName,
		userRate.CollectedTweetsCount,
		userRate.LikesTweetsCount,
		userRate.ReplyTweetsCount,
		userRate.RetweetedTweetsCount,
		userRate.LanguagesStr,
		userRate.primaryLanguage,
		userRate.primaryLanguageCount,
		userRate.GoldenFactor,
		time.Now().Unix())
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

func saveGoldenTweetsToDB(goldenTweetsTB *GoldenTweetsTB, db *sql.DB) error {
	query := fmt.Sprintf(`SELECT public."InsertGoldenTweetsTB"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);`)
	rows, err := db.Query(query,
		goldenTweetsTB.TweetID,
		goldenTweetsTB.UserID,
		goldenTweetsTB.UserName,
		goldenTweetsTB.Favorite,
		goldenTweetsTB.Reply,
		goldenTweetsTB.Retweet,
		goldenTweetsTB.TweetLanguage,
		goldenTweetsTB.UserPrimaryLanguage,
		goldenTweetsTB.TweetSilverFactor,
		goldenTweetsTB.UserGoldenFactor,
		goldenTweetsTB.MeasureRate,
		goldenTweetsTB.CreatedAt,
		goldenTweetsTB.UpdatedAt,
		goldenTweetsTB.Expired,
		goldenTweetsTB.TweetType.String())
	if err != nil {
		logs.Critical(err)
		return err
	}
	rows.Close()
	return nil
}

func updateUserState(db *sql.DB, userID int64, status int) {
	q := fmt.Sprintf(`UPDATE "UsersProfilesTB" 
		SET "TwitterState"= %v,"UpdatedAt"=%v
		 WHERE "UserId" = %v;`, status, time.Now().Unix(), userID)
	rows, err := db.Query(q)
	if err != nil {
		logs.Critical(err)
	}
	rows.Close()
}

//GetUsersPageToCrawl :
func GetUsersPageToCrawl(db *sql.DB, limit int) []int64 {
	q := fmt.Sprintf(`SELECT "userid" FROM "GetUsersPageToCrawl"(%d);`, limit)
	rows, err := db.Query(q)
	if err != nil {
		logs.Critical(err)
	}
	defer rows.Close()

	pages := make([]int64, limit)

	i := 0
	for rows.Next() {
		var userid int64
		if err := rows.Scan(&userid); err != nil {
			logs.Critical(err)
		} else {
			pages[i] = userid
			i++
		}
	}
	logs.Infof("GetUsersPageToCrawl(from %v)\n", limit)
	return pages
}

/*
func savetweetsListToDB(tweets *[]anaconda.Tweet, Source int, db *sql.DB, api *anaconda.TwitterApi, chGoodTweets chan *anaconda.Tweet) {
	defer recovery()
	valueStrings := make([]string, 0, len(*tweets))
	valueArgs := make([]interface{}, 0, len(*tweets)*19)
	for i, tweet := range *tweets {
		var RetweetedTweetID sql.NullInt64
		var RetweetedUserID sql.NullInt64

		var QuotedTweetID sql.NullInt64
		var QuotedUserID sql.NullInt64

		var ReplyTweetID sql.NullInt64
		var ReplyUserID sql.NullInt64

		tweetTime, err := tweet.CreatedAtTime()
		if err != nil {
			logs.Critical(err)
		}

		if tweet.RetweetedStatus != nil {
			//savetweetToDB(tweet.RetweetedStatus, 21, false, db, api, chGoodTweets) //APIUserTimelineRetweet
			RetweetedTweetID = sql.NullInt64{Int64: tweet.RetweetedStatus.Id, Valid: true}
			RetweetedUserID = sql.NullInt64{Int64: tweet.RetweetedStatus.User.Id, Valid: true}
			//owner user only should hold retweet value
			tweet.FavoriteCount = 0
			tweet.ReplyCount = 0
			tweet.RetweetCount = 0
		}

		if tweet.QuotedStatus != nil {
			//savetweetToDB(tweet.QuotedStatus, 22, false, db, api, chGoodTweets) //APIUserTimelineQuote
			QuotedTweetID = sql.NullInt64{Int64: tweet.QuotedStatus.Id, Valid: true}
			QuotedUserID = sql.NullInt64{Int64: tweet.QuotedStatus.User.Id, Valid: true}
		}

		if tweet.InReplyToStatusID != 0 {
			ReplyTweetID = sql.NullInt64{Int64: tweet.InReplyToStatusID, Valid: true}
			ReplyUserID = sql.NullInt64{Int64: tweet.InReplyToUserID, Valid: true}
		}
		var CountryCode sql.NullString
		if &tweet.Place != nil && len(tweet.Place.CountryCode) > 1 {
			CountryCode = sql.NullString{String: tweet.Place.CountryCode, Valid: true}
		}

		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			i*19+1, i*19+2, i*19+3, i*19+4, i*19+5, i*19+6, i*19+7, i*19+8, i*19+9, i*19+10, i*19+11, i*19+12, i*19+13, i*19+14, i*19+15, i*19+16, i*19+17, i*19+18, i*19+19))
		valueArgs = append(valueArgs, tweet.Id)
		valueArgs = append(valueArgs, tweet.User.Id)
		valueArgs = append(valueArgs, tweet.FavoriteCount)
		valueArgs = append(valueArgs, tweet.ReplyCount)
		valueArgs = append(valueArgs, tweet.RetweetCount)
		valueArgs = append(valueArgs, getSilverFactor(&tweet))
		valueArgs = append(valueArgs, tweet.User.ScreenName)
		valueArgs = append(valueArgs, tweetTime.Unix())
		valueArgs = append(valueArgs, time.Now().Unix()) //updated at time now)
		valueArgs = append(valueArgs, tweet.Lang)
		valueArgs = append(valueArgs, CountryCode)
		valueArgs = append(valueArgs, RetweetedTweetID)
		valueArgs = append(valueArgs, RetweetedUserID)
		valueArgs = append(valueArgs, QuotedTweetID)
		valueArgs = append(valueArgs, QuotedUserID)
		valueArgs = append(valueArgs, ReplyTweetID)
		valueArgs = append(valueArgs, ReplyUserID)
		valueArgs = append(valueArgs, Source)
		valueArgs = append(valueArgs, encodeWindows1250(tweet.FullText))
	}

	//should not update Reply count
	//"Reply" = COALESCE(excluded."Reply","TweetsTB"."Reply"),

	query := fmt.Sprintf(`
		INSERT INTO public."TweetsTB"(
			"TweetId", "UserId", "Favorite", "Reply", "Retweet", "SilverFactor", "UserName", "CreatedAt", "UpdatedAt", "Lang", "CountryCode", "RetweetedTweetId", "RetweetedUserId", "QuotedTweetId", "QuotedUserId", "ReplyTweetId", "ReplyUserId", "Source", "Text")
			VALUES %v
			ON CONFLICT ("TweetId") DO UPDATE
			SET
			  "UserId" = COALESCE(excluded."UserId","TweetsTB"."UserId"),
			  "Favorite" = COALESCE(excluded."Favorite","TweetsTB"."Favorite"),
			  "Retweet" = COALESCE(excluded."Retweet","TweetsTB"."Retweet"),
			  "SilverFactor" = COALESCE(excluded."SilverFactor","TweetsTB"."SilverFactor"),
			  "UserName" = COALESCE(excluded."UserName","TweetsTB"."UserName"),
			  "CreatedAt" = COALESCE(excluded."CreatedAt","TweetsTB"."CreatedAt"),
			  "UpdatedAt" = COALESCE(excluded."UpdatedAt","TweetsTB"."UpdatedAt"),
			  "Lang" = COALESCE(excluded."Lang","TweetsTB"."Lang"),
			  "CountryCode" = COALESCE(excluded."CountryCode","TweetsTB"."CountryCode"),
			  "RetweetedTweetId" = COALESCE(excluded."RetweetedTweetId","TweetsTB"."RetweetedTweetId"),
			  "RetweetedUserId" = COALESCE(excluded."RetweetedUserId","TweetsTB"."RetweetedUserId"),
			  "QuotedTweetId" = COALESCE(excluded."QuotedTweetId","TweetsTB"."QuotedTweetId"),
			  "QuotedUserId" = COALESCE(excluded."QuotedUserId","TweetsTB"."QuotedUserId"),
			  "ReplyTweetId" = COALESCE(excluded."ReplyTweetId","TweetsTB"."ReplyTweetId"),
			  "ReplyUserId" = COALESCE(excluded."ReplyUserId","TweetsTB"."ReplyUserId"),
			  "Source" = COALESCE(excluded."Source","TweetsTB"."Source"),
			  "Text" = COALESCE(excluded."Text","TweetsTB"."Text")
			  ;
		`, strings.Join(valueStrings, ","))
	rows, err := db.Query(query, valueArgs...)
	if err != nil {
		logs.Critical(err)
	}
	rows.Close()
	for _, tweet := range *tweets {
		saveUserToDB(&tweet, Source, true, db)
		break
	}
}
*/
