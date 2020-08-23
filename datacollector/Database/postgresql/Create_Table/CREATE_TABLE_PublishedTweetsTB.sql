DROP TABLE IF EXISTS public."PublishedTweetsTB";


CREATE TABLE public."PublishedTweetsTB"
(
    "TweetId" BIGINT NOT NULL,
    "UserId" BIGINT NOT NULL,
    "UserName" TEXT COLLATE pg_catalog."default" NOT NULL,
    "TweetCreatedAt" BIGINT NOT NULL,
    "Lang" TEXT COLLATE pg_catalog."default" NOT NULL,
    "Type" JSON NOT NULL,
    "PagePublishPath" TEXT COLLATE pg_catalog."default" NOT NULL,
    "TemplateVersion" TEXT COLLATE pg_catalog."default" NOT NULL,
    "IsJSONCreated" BOOLEAN NOT NULL,
    "IsImageCreated" BOOLEAN NOT NULL,
    "IsGitPushed" BOOLEAN NOT NULL,
    "CreatedAt" BIGINT NOT NULL,
    "UpdatedAt" BIGINT NOT NULL,
    "IsDeleted" BOOLEAN NOT NULL,
    "RelatedTweets" JSON, 
    "SoicalPublish" JSON,
    "ShortURL" TEXT COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "PublishedTweetsTB_pkey" PRIMARY KEY ("TweetId")
); 
