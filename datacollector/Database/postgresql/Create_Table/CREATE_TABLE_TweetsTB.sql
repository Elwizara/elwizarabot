DROP TABLE IF EXISTS public."TweetsTB";

CREATE TABLE public."TweetsTB"
(
    "TweetId" BIGINT NOT NULL,
    "UserId" BIGINT,
    "Favorite" INTEGER,
    "Reply" INTEGER,
    "Retweet" INTEGER,
    "SilverFactor" BIGINT,
    "UserName" TEXT COLLATE pg_catalog."default",
    "CreatedAt" BIGINT,
    "UpdatedAt" BIGINT,
    "Lang" TEXT COLLATE pg_catalog."default",
    "CountryCode" TEXT COLLATE pg_catalog."default",
    "RetweetedTweetId" BIGINT,
    "RetweetedUserId" BIGINT,
    "QuotedTweetId" BIGINT,
    "QuotedUserId" BIGINT,
    "ReplyTweetId" BIGINT,
    "ReplyUserId" BIGINT,
    "Source" SMALLINT,
    "Text" TEXT COLLATE pg_catalog."default",
    "Type" JSON,
    CONSTRAINT "TweetsTB_pkey" PRIMARY KEY ("TweetId")
);


