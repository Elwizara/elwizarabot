CREATE TABLE public."UsersRateTB"
(
    "UserId" 	BIGINT NOT NULL,
	"UserName" 	TEXT COLLATE pg_catalog."default",
    "CollectedTweetsCount" INTEGER,
    "LikesTweetsCount" INTEGER,
    "ReplyTweetsCount" INTEGER,
    "RetweetedTweetsCount" INTEGER,
	"Languages" TEXT COLLATE pg_catalog."default",
	"primaryLanguage" TEXT COLLATE pg_catalog."default",
	"primaryLanguageCount" INTEGER,
    "GoldenFactor" BIGINT, 
    "UpdatedAt" BIGINT,
    CONSTRAINT "UsersRateTB_pkey" PRIMARY KEY ("UserId")
);