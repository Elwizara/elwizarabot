CREATE TABLE public."GoldenTweetsTB"
(
    "TweetId" BIGINT NOT NULL,
    "UserId" BIGINT NOT NULL,
	"UserName" TEXT COLLATE pg_catalog."default" NOT NULL,
    "Favorite" INTEGER NOT NULL,
    "Reply" INTEGER NOT NULL,
    "Retweet" INTEGER NOT NULL,
    "TweetLanguage" TEXT COLLATE pg_catalog."default" NOT NULL, 
	"UserPrimaryLanguage" TEXT COLLATE pg_catalog."default" NOT NULL,     
    "TweetSilverFactor" INTEGER NOT NULL, 
	"UserGoldenFactor" INTEGER NOT NULL, 
    "MeasureRate" INTEGER NOT NULL,  
	"CreatedAt" BIGINT NOT NULL,
    "UpdatedAt" BIGINT NOT NULL,
	"Expired" BOOLEAN NOT NULL,
	"Type" JSON NOT NULL,
    CONSTRAINT "GoldenTweetsTB_pkey" PRIMARY KEY ("TweetId")
);