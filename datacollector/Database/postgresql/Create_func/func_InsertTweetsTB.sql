CREATE OR REPLACE FUNCTION public."InsertTweetsTB"
	(
		TweetId bigint,
		UserId bigint,
		Favorite bigint,
		Reply bigint,
		Retweet bigint,
		SilverFactor bigint,
		UserName text,
		CreatedAt bigint,
		UpdatedAt bigint,
		Lang text,
		CountryCode text,
		RetweetedTweetId bigint,
		RetweetedUserId bigint,
		QuotedTweetId bigint,
		QuotedUserId bigint,
		ReplyTweetId bigint,
		ReplyUserId bigint,
		Source integer,
		Tweettext text,
		TweetType JSON) 
RETURNS VOID AS
$$
BEGIN 
	INSERT INTO "TweetsTB"(
					"TweetId",
					"UserId",
					"Favorite",
					"Reply",
					"Retweet",
					"SilverFactor",
					"UserName",
					"CreatedAt",
					"UpdatedAt",
					"Lang",
					"CountryCode",
					"RetweetedTweetId",
					"RetweetedUserId",
					"QuotedTweetId",
					"QuotedUserId",
					"ReplyTweetId",
					"ReplyUserId",
					"Source",
					"Text",
					"Type"
				)
		VALUES (
				TweetId,
				UserId,
				Favorite,
				Reply,
				Retweet,
				SilverFactor,
				UserName,
				CreatedAt,
				UpdatedAt,
				Lang,
				CountryCode,
				RetweetedTweetId,
				RetweetedUserId,
				QuotedTweetId,
				QuotedUserId,
				ReplyTweetId,
				ReplyUserId,
				Source,
				Tweettext,
				TweetType
			)
		ON CONFLICT ("TweetId") DO UPDATE 
		  SET  
			"UserId" = COALESCE(excluded."UserId","TweetsTB"."UserId"),
			"Favorite" = COALESCE(excluded."Favorite","TweetsTB"."Favorite"),
			"Reply" = COALESCE(excluded."Reply","TweetsTB"."Reply"),
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
			"Text" = COALESCE(excluded."Text","TweetsTB"."Text"),
			"Type" = COALESCE(excluded."Type","TweetsTB"."Type");
END
$$
LANGUAGE 'plpgsql';
