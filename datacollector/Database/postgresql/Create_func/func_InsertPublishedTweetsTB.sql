CREATE OR REPLACE FUNCTION public."InsertPublishedTweetsTB"
    (
        TweetId BIGINT,
        UserId  BIGINT,
        UserName    TEXT,
        TweetCreatedAt  BIGINT,
        Lang    TEXT,
        TweetType JSON,
        PagePublishPath TEXT,
        TemplateVersion TEXT,
        IsJSONCreated   BOOLEAN,
        IsImageCreated BOOLEAN,
        IsGitPushed BOOLEAN,
        CreatedAt   BIGINT,
        UpdatedAt   BIGINT,
        IsDeleted   BOOLEAN,
        RelatedTweets JSON,
        ShortURL      TEXT
    )
RETURNS VOID AS
$$
BEGIN
        UPDATE "GoldenTweetsTB"
            SET
            "Expired" = TRUE
            WHERE
            "TweetId" = TweetId;

		INSERT INTO "PublishedTweetsTB"(
			"TweetId",
            "UserId",
            "UserName",
            "TweetCreatedAt",
            "Lang",
            "Type",
            "PagePublishPath",
            "TemplateVersion",
            "IsJSONCreated",
            "IsImageCreated",
            "IsGitPushed",
            "CreatedAt",
            "UpdatedAt",
            "IsDeleted",
            "RelatedTweets",
            "ShortURL")
		VALUES (
			TweetId,
            UserId,
            UserName,
            TweetCreatedAt,
            Lang,
            TweetType,
            PagePublishPath,
            TemplateVersion,
            IsJSONCreated,
            IsImageCreated,
            IsGitPushed,
            CreatedAt,
            UpdatedAt,
            IsDeleted,
            RelatedTweets,
            ShortURL)
		ON CONFLICT ("TweetId") DO UPDATE 
			SET  
                "UserId" = excluded."UserId",
                "UserName" = excluded."UserName",
                "TweetCreatedAt" = excluded."TweetCreatedAt",
                "Lang" = excluded."Lang",
                "Type"= excluded."Type",
                "PagePublishPath" = excluded."PagePublishPath",
                "TemplateVersion" = excluded."TemplateVersion",
                "IsJSONCreated" = excluded."IsJSONCreated",
                "IsImageCreated" = excluded."IsImageCreated",
                "IsGitPushed" = excluded."IsGitPushed",
                "CreatedAt" = excluded."CreatedAt",
                "UpdatedAt" = excluded."UpdatedAt",
                "IsDeleted" = excluded."IsDeleted",
                "RelatedTweets" = excluded."RelatedTweets",
                "ShortURL" = excluded."ShortURL";
END
$$
LANGUAGE 'plpgsql';

