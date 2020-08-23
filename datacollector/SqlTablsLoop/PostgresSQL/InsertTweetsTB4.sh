#!/bin/bash 

query=" 
    BEGIN;     
    INSERT INTO \"TweetsTB_$2\"
    (SELECT * FROM \"TweetsTB\" WHERE (\"UserId\"%1000) = $2)
        ON CONFLICT (\"TweetId\") DO UPDATE
            SET
                \"UserId\" = COALESCE(excluded.\"UserId\",\"TweetsTB_$2\".\"UserId\"),
                \"Favorite\" = COALESCE(excluded.\"Favorite\",\"TweetsTB_$2\".\"Favorite\"),
                \"Reply\" = COALESCE(excluded.\"Reply\",\"TweetsTB_$2\".\"Reply\"),
                \"Retweet\" = COALESCE(excluded.\"Retweet\",\"TweetsTB_$2\".\"Retweet\"),
                \"SilverFactor\" = COALESCE(excluded.\"SilverFactor\",\"TweetsTB_$2\".\"SilverFactor\"),
                \"UserName\" = COALESCE(excluded.\"UserName\",\"TweetsTB_$2\".\"UserName\"),
                \"CreatedAt\" = COALESCE(excluded.\"CreatedAt\",\"TweetsTB_$2\".\"CreatedAt\"),
                \"UpdatedAt\" = COALESCE(excluded.\"UpdatedAt\",\"TweetsTB_$2\".\"UpdatedAt\"),
                \"Lang\" = COALESCE(excluded.\"Lang\",\"TweetsTB_$2\".\"Lang\"),
                \"CountryCode\" = COALESCE(excluded.\"CountryCode\",\"TweetsTB_$2\".\"CountryCode\"),
                \"RetweetedTweetId\" = COALESCE(excluded.\"RetweetedTweetId\",\"TweetsTB_$2\".\"RetweetedTweetId\"),
                \"RetweetedUserId\" = COALESCE(excluded.\"RetweetedUserId\",\"TweetsTB_$2\".\"RetweetedUserId\"),
                \"QuotedTweetId\" = COALESCE(excluded.\"QuotedTweetId\",\"TweetsTB_$2\".\"QuotedTweetId\"),
                \"QuotedUserId\" = COALESCE(excluded.\"QuotedUserId\",\"TweetsTB_$2\".\"QuotedUserId\"),
                \"ReplyTweetId\" = COALESCE(excluded.\"ReplyTweetId\",\"TweetsTB_$2\".\"ReplyTweetId\"),
                \"ReplyUserId\" = COALESCE(excluded.\"ReplyUserId\",\"TweetsTB_$2\".\"ReplyUserId\"),
                \"Source\" = COALESCE(excluded.\"Source\",\"TweetsTB_$2\".\"Source\"),
				\"Text\" = COALESCE(excluded.\"Text\",\"TweetsTB_$2\".\"Text\");

    COMMIT;
" 
psql "$1" -c "$query"
