#!/bin/bash   
#./PostgresSQL/InsertTweetsTB4.sh "host=204.93.216.157 port=5432 user=elwizara password=${pass} dbname=ElwizaraLIVEDB" "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable" 1000.gz
 
dollarSign='$$'

query1="        
    COPY(SELECT * FROM \"TweetsTB\" WHERE \"UpdatedAt\" < $4) 
        TO PROGRAM 'gzip > /var/lib/postgresql/tmp/TweetsTB/$3';"



query2="DO $dollarSign
        DECLARE query text;
        BEGIN 
        FOR counter IN 0..999 LOOP 
            query:= 'INSERT INTO \"TweetsTB_'||counter||'\"  SELECT * FROM \"TweetsTB\" WHERE (\"UserId\"%1000)=' || counter;
            query := query || '        ON CONFLICT (\"TweetId\") DO UPDATE
                    SET
                        \"UserId\" = COALESCE(excluded.\"UserId\",\"TweetsTB_'||counter||'\".\"UserId\"),
                        \"Favorite\" = COALESCE(excluded.\"Favorite\",\"TweetsTB_'||counter||'\".\"Favorite\"),
                        \"Reply\" = COALESCE(excluded.\"Reply\",\"TweetsTB_'||counter||'\".\"Reply\"),
                        \"Retweet\" = COALESCE(excluded.\"Retweet\",\"TweetsTB_'||counter||'\".\"Retweet\"),
                        \"SilverFactor\" = COALESCE(excluded.\"SilverFactor\",\"TweetsTB_'||counter||'\".\"SilverFactor\"),
                        \"UserName\" = COALESCE(excluded.\"UserName\",\"TweetsTB_'||counter||'\".\"UserName\"),
                        \"CreatedAt\" = COALESCE(excluded.\"CreatedAt\",\"TweetsTB_'||counter||'\".\"CreatedAt\"),
                        \"UpdatedAt\" = COALESCE(excluded.\"UpdatedAt\",\"TweetsTB_'||counter||'\".\"UpdatedAt\"),
                        \"Lang\" = COALESCE(excluded.\"Lang\",\"TweetsTB_'||counter||'\".\"Lang\"),
                        \"CountryCode\" = COALESCE(excluded.\"CountryCode\",\"TweetsTB_'||counter||'\".\"CountryCode\"),
                        \"RetweetedTweetId\" = COALESCE(excluded.\"RetweetedTweetId\",\"TweetsTB_'||counter||'\".\"RetweetedTweetId\"),
                        \"RetweetedUserId\" = COALESCE(excluded.\"RetweetedUserId\",\"TweetsTB_'||counter||'\".\"RetweetedUserId\"),
                        \"QuotedTweetId\" = COALESCE(excluded.\"QuotedTweetId\",\"TweetsTB_'||counter||'\".\"QuotedTweetId\"),
                        \"QuotedUserId\" = COALESCE(excluded.\"QuotedUserId\",\"TweetsTB_'||counter||'\".\"QuotedUserId\"),
                        \"ReplyTweetId\" = COALESCE(excluded.\"ReplyTweetId\",\"TweetsTB_'||counter||'\".\"ReplyTweetId\"),
                        \"ReplyUserId\" = COALESCE(excluded.\"ReplyUserId\",\"TweetsTB_'||counter||'\".\"ReplyUserId\"),
                        \"Source\" = COALESCE(excluded.\"Source\",\"TweetsTB_'||counter||'\".\"Source\"),
                        \"Text\" = COALESCE(excluded.\"Text\",\"TweetsTB_'||counter||'\".\"Text\");'; 
            EXECUTE query;
        END LOOP;
        END; $dollarSign
"

psql "$1" -c "$query1" 

sshpass -p domian12 scp -P 2222 root@204.93.216.157:/var/lib/postgresql/tmp/TweetsTB/$3 /RED/tmp/TweetsTB/ 

zcat /RED/tmp/TweetsTB/$3 | psql "$2" -c " COPY \"TweetsTB\" FROM STDIN;"

psql "$2" -c "$query2"

psql "$2" -c "truncate table \"TweetsTB\";"
  
psql "$1" -c "DELETE FROM \"TweetsTB\" WHERE \"UpdatedAt\" < $4;" 



#select count(*) from "TweetsTB"; --"54085970" - "53587388"


# select count(*) from "TweetsTB_0";--"547082"-"548357"
# select count(*) from "TweetsTB_1";--"304566"-"304824"
# select count(*) from "TweetsTB_2";--"236434"-"236755"
# select count(*) from "TweetsTB_3";--"171654"-"171861"
# select count(*) from "TweetsTB_4";--"202223"-"202696"
# select count(*) from "TweetsTB_5";--"164953"-"165000"
# select count(*) from "TweetsTB_100";--"164953"-"211782" 
# select count(*) from "TweetsTB_588";--"194631"-"195231"
# select count(*) from "TweetsTB_700";--"189787"-"189807"
# select count(*) from "TweetsTB_980";--"202238"-"202680"
# select count(*) from "TweetsTB_998";--"200004"-"200447"