#!/bin/bash  
#./PostgresSQL/InsertTweetsTB2.sh "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable"

query1="
    COPY(  
        WITH Removed_rows AS(
            DELETE FROM \"TweetsTB\" 
            WHERE \"UpdatedAt\" <= (
                SELECT MAX(\"UpdatedAt\") FROM (
                    SELECT \"UpdatedAt\" FROM \"TweetsTB\" order by \"UpdatedAt\" LIMIT $3
                    )AS MAXUpdatedAt 
                ) 
            RETURNING *
        ) 
        SELECT * FROM Removed_rows  
    ) TO STDOUT;
"

query2=" 
    BEGIN;
    CREATE TEMP TABLE \"tmp_table\"
		ON COMMIT DROP
		AS
		SELECT * FROM \"TweetsTB\" WITH NO DATA;

		COPY \"tmp_table\" FROM STDIN;

		INSERT INTO \"TweetsTB\"
		SELECT \"tmp_table\".* FROM \"tmp_table\" 

        ON CONFLICT (\"TweetId\") DO UPDATE
            SET
                \"UserId\" = COALESCE(excluded.\"UserId\",\"TweetsTB\".\"UserId\"),
                \"Favorite\" = COALESCE(excluded.\"Favorite\",\"TweetsTB\".\"Favorite\"),
                \"Reply\" = COALESCE(excluded.\"Reply\",\"TweetsTB\".\"Reply\"),
                \"Retweet\" = COALESCE(excluded.\"Retweet\",\"TweetsTB\".\"Retweet\"),
                \"SilverFactor\" = COALESCE(excluded.\"SilverFactor\",\"TweetsTB\".\"SilverFactor\"),
                \"UserName\" = COALESCE(excluded.\"UserName\",\"TweetsTB\".\"UserName\"),
                \"CreatedAt\" = COALESCE(excluded.\"CreatedAt\",\"TweetsTB\".\"CreatedAt\"),
                \"UpdatedAt\" = COALESCE(excluded.\"UpdatedAt\",\"TweetsTB\".\"UpdatedAt\"),
                \"Lang\" = COALESCE(excluded.\"Lang\",\"TweetsTB\".\"Lang\"),
                \"CountryCode\" = COALESCE(excluded.\"CountryCode\",\"TweetsTB\".\"CountryCode\"),
                \"RetweetedTweetId\" = COALESCE(excluded.\"RetweetedTweetId\",\"TweetsTB\".\"RetweetedTweetId\"),
                \"RetweetedUserId\" = COALESCE(excluded.\"RetweetedUserId\",\"TweetsTB\".\"RetweetedUserId\"),
                \"QuotedTweetId\" = COALESCE(excluded.\"QuotedTweetId\",\"TweetsTB\".\"QuotedTweetId\"),
                \"QuotedUserId\" = COALESCE(excluded.\"QuotedUserId\",\"TweetsTB\".\"QuotedUserId\"),
                \"ReplyTweetId\" = COALESCE(excluded.\"ReplyTweetId\",\"TweetsTB\".\"ReplyTweetId\"),
                \"ReplyUserId\" = COALESCE(excluded.\"ReplyUserId\",\"TweetsTB\".\"ReplyUserId\"),
                \"Source\" = COALESCE(excluded.\"Source\",\"TweetsTB\".\"Source\"),
				\"Text\" = COALESCE(excluded.\"Text\",\"TweetsTB\".\"Text\");
    COMMIT;
"
psql "$1" -c "$query1"  | psql "$2" -c "$query2"