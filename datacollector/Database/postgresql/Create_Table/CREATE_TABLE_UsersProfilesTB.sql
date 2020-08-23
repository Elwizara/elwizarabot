DROP TABLE IF EXISTS public."UsersProfilesTB";

CREATE TABLE public."UsersProfilesTB"
(
    "UserId" bigint NOT NULL,
    "UserName" text COLLATE pg_catalog."default",
    "ViewName" text COLLATE pg_catalog."default",
    "Bio" text COLLATE pg_catalog."default",
    "Location" text COLLATE pg_catalog."default",
    "Link" text COLLATE pg_catalog."default",
    "JoinDate" bigint,
    "TweetsCount" integer,
    "LikesCount" integer,
    "FollowingCount" integer,
    "FollowersCount" integer,
    "BirthDate" text COLLATE pg_catalog."default",
    "UpdatedAt" bigint,
    "LastTweetDate" bigint,
    "LastTweetId" bigint,
    "TwitterState" integer,
    "LastSource" integer,
    CONSTRAINT "UsersProfilesTB_pkey" PRIMARY KEY ("UserId")
);


