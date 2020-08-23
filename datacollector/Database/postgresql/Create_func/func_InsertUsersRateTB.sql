CREATE OR REPLACE FUNCTION public."InsertUsersRateTB"
    (
        UserId  BIGINT,
        UserName    TEXT,
        CollectedTweetsCount    INTEGER,
        LikesTweetsCount    INTEGER,
        ReplyTweetsCount    INTEGER,
        RetweetedTweetsCount    INTEGER,
        Languages   TEXT,
        primaryLanguage TEXT,
        primaryLanguageCount    INTEGER,
        GoldenFactor    BIGINT,
        UpdatedAt   BIGINT
    )
RETURNS VOID AS
$$
BEGIN
		INSERT INTO "UsersRateTB"(
			"UserId",
            "UserName",
            "CollectedTweetsCount",
            "LikesTweetsCount",
            "ReplyTweetsCount",
            "RetweetedTweetsCount",
            "Languages",
            "primaryLanguage",
            "primaryLanguageCount",
            "GoldenFactor",
            "UpdatedAt")
		VALUES (
			UserId,
            UserName,
            CollectedTweetsCount,
            LikesTweetsCount,
            ReplyTweetsCount,
            RetweetedTweetsCount,
            Languages,
            primaryLanguage,
            primaryLanguageCount,
            GoldenFactor,
            UpdatedAt)
		ON CONFLICT ("UserId") DO UPDATE 
			SET  
                "UserName" = excluded."UserName",
                "CollectedTweetsCount" = excluded."CollectedTweetsCount",
                "LikesTweetsCount" = excluded."LikesTweetsCount",
                "ReplyTweetsCount" = excluded."ReplyTweetsCount",
                "RetweetedTweetsCount" = excluded."RetweetedTweetsCount",
                "Languages" = excluded."Languages",
                "primaryLanguage" = excluded."primaryLanguage",
                "primaryLanguageCount" = excluded."primaryLanguageCount",
                "GoldenFactor" = excluded."GoldenFactor",
                "UpdatedAt" = excluded."UpdatedAt";
END
$$
LANGUAGE 'plpgsql';

