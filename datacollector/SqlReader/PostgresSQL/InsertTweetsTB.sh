#!/bin/bash  
#psql "host=elwizaralivecloud1.postgres.database.azure.com port=5432 user=elwizaracloud1@elwizaralivecloud1 password=A9Sd42^d%KjrG3MK dbname=ElwizaraLIVEDB" -c "COPY(select * from \"TweetsTB\" limit 1) TO STDOUT;" | psql "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" -c '\copy "TweetsTB" FROM STDIN'
#./InsertTweetsTB.sh "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" 792 1529953716 "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB2 sslmode=disable"

query=" 
    BEGIN;
    CREATE TEMP TABLE \"tmp_table\"
		ON COMMIT DROP
		AS
		SELECT * FROM \"TweetsTB_$2\" WITH NO DATA;

		COPY \"tmp_table\" FROM STDIN;

		INSERT INTO \"TweetsTB_$2\"
		SELECT \"tmp_table\".* FROM \"tmp_table\" 

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
psql "$1" -c "COPY(select * from \"TweetsTB\" WHERE (\"UserId\"%1000)=$2 AND  \"UpdatedAt\" <= '$3') TO STDOUT;" | psql "$4" -c "$query"
