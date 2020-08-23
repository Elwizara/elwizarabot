CREATE OR REPLACE FUNCTION public."InsertGoldenTweetsTB"
    ( 
        TweetId BIGINT,
        UserId BIGINT,
        UserName TEXT,
        Favorite INTEGER,
        Reply INTEGER,
        Retweet INTEGER,
        TweetLanguage TEXT,
        UserPrimaryLanguage TEXT,
        TweetSilverFactor INTEGER,
        UserGoldenFactor INTEGER,
        MeasureRate INTEGER,
        CreatedAt BIGINT,
        UpdatedAt BIGINT,
        Expired BOOLEAN,
        TweetType JSON
    )
RETURNS VOID AS
$$
BEGIN
		INSERT INTO "GoldenTweetsTB"(
            "TweetId",
            "UserId",
            "UserName",
            "Favorite",
            "Reply",
            "Retweet",
            "TweetLanguage",
            "UserPrimaryLanguage",
            "TweetSilverFactor",
            "UserGoldenFactor",
            "MeasureRate",
            "CreatedAt",
            "UpdatedAt",
            "Expired",
            "Type")
        VALUES (
			TweetId,
            UserId,
            UserName,
            Favorite,
            Reply,
            Retweet,
            TweetLanguage,
            UserPrimaryLanguage,
            TweetSilverFactor,
            UserGoldenFactor,
            MeasureRate,
            CreatedAt,
            UpdatedAt,
            Expired,
            TweetType)
		ON CONFLICT ("TweetId") DO UPDATE 
			SET    
                "UserId" = excluded."UserId",
                "UserName" = excluded."UserName",
                "Favorite" = excluded."Favorite",
                "Reply" = excluded."Reply",
                "Retweet" = excluded."Retweet",
                "TweetLanguage" = excluded."TweetLanguage",
                "UserPrimaryLanguage" = excluded."UserPrimaryLanguage",
                "TweetSilverFactor" = excluded."TweetSilverFactor",
                "UserGoldenFactor" = excluded."UserGoldenFactor",
                "MeasureRate" = excluded."MeasureRate",
                "CreatedAt" = excluded."CreatedAt",
                "UpdatedAt" = excluded."UpdatedAt",
                "Expired" = excluded."Expired",
                "Type"= excluded."Type";
                
END
$$
LANGUAGE 'plpgsql';

