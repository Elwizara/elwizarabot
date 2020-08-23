#!/bin/bash 

query=" 
    BEGIN;
    CREATE TABLE IF NOT EXISTS public.\"TweetsTB_$2\"
    (
        \"TweetId\" bigint NOT NULL,
        \"UserId\" bigint NOT NULL,
        \"Favorite\" integer NOT NULL,
        \"Reply\" integer NOT NULL,
        \"Retweet\" integer NOT NULL,
        \"SilverFactor\" bigint NOT NULL,
        \"UserName\" text COLLATE pg_catalog.\"default\",
        \"CreatedAt\" bigint,
        \"UpdatedAt\" bigint,
        \"Lang\" text COLLATE pg_catalog.\"default\",
        \"CountryCode\" text COLLATE pg_catalog.\"default\",
        \"RetweetedTweetId\" bigint,
        \"RetweetedUserId\" bigint,
        \"QuotedTweetId\" bigint,
        \"QuotedUserId\" bigint,
        \"ReplyTweetId\" bigint,
        \"ReplyUserId\" bigint,
        \"Source\" smallint,
        \"Text\" text COLLATE pg_catalog.\"default\",
        CONSTRAINT \"TweetsTB_$2_pkey\" PRIMARY KEY (\"TweetId\")
    );
    CREATE INDEX IF NOT EXISTS \"idx_TweetsTB_$2_UpdatedAt\"
        ON public.\"TweetsTB_$2\" USING btree
        (\"UpdatedAt\")
        TABLESPACE pg_default;
    COMMIT;
"
psql "$1" -c "$query"
