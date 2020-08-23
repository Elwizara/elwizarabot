#!/bin/bash    
#sudo ./InsertGoldenTweetsTB.sh "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable"  "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" 1000.gz

query1="
    COPY(SELECT * FROM \"GoldenTweetsTB\" WHERE \"Type\" IS NOT NULL) TO PROGRAM 'gzip > /RED/Database/postgresql/$3'; 
"

query2=" 
    BEGIN;
    CREATE TEMP TABLE \"tmp_table\"
		ON COMMIT DROP
		AS
		SELECT * FROM \"GoldenTweetsTB\" WITH NO DATA;

		COPY \"tmp_table\" FROM STDIN;

		INSERT INTO \"GoldenTweetsTB\"
		SELECT * FROM \"tmp_table\" 

        ON CONFLICT (\"TweetId\") DO UPDATE
            SET  
                \"UserId\"=excluded.\"UserId\",
                \"UserName\"=excluded.\"UserName\",
                \"Favorite\"=excluded.\"Favorite\",
                \"Reply\"=excluded.\"Reply\",
                \"Retweet\"=excluded.\"Retweet\",
                \"TweetLanguage\"=excluded.\"TweetLanguage\",
                \"UserPrimaryLanguage\"=excluded.\"UserPrimaryLanguage\",
                \"TweetSilverFactor\"=excluded.\"TweetSilverFactor\",
                \"UserGoldenFactor\"=excluded.\"UserGoldenFactor\",
                \"MeasureRate\"=excluded.\"MeasureRate\",
                \"CreatedAt\"=excluded.\"CreatedAt\",
                \"UpdatedAt\"=excluded.\"UpdatedAt\";
    COMMIT;
"
  
psql "$1" -c "$query1" 
 
zcat /RED/Database/postgresql/$3 | psql "$2" -c "$query2"



