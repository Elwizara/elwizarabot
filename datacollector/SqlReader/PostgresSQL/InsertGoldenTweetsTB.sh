#!/bin/bash   
#./PostgresSQL/InsertGoldenTweetsTB.sh "host=204.93.216.157 port=5432 user=elwizara password=${pass} dbname=ElwizaraLIVEDB"  "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable" 1000.gz
 
LastUpdatedAt=$(psql "$2" -c "SELECT \"UpdatedAt\" FROM \"GoldenTweetsTB\" ORDER BY \"UpdatedAt\" DESC  LIMIT 1;"  -t -A)
 
query1="
    COPY(SELECT * FROM \"GoldenTweetsTB\" WHERE \"UpdatedAt\">=$LastUpdatedAt) TO PROGRAM 'gzip > /var/lib/postgresql/tmp/GoldenTweetsTB/$3'; 
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

sshpass -p domian12 scp -P 2222 root@204.93.216.157:/var/lib/postgresql/tmp/GoldenTweetsTB/$3 /RED/tmp/GoldenTweetsTB/ 

zcat /RED/tmp/GoldenTweetsTB/$3 | psql "$2" -c "$query2"



