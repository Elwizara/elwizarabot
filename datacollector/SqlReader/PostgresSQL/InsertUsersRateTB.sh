#!/bin/bash   
#./PostgresSQL/InsertUsersRateTB.sh "host=204.93.216.157 port=5432 user=elwizara password=${pass} dbname=ElwizaraLIVEDB"  "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable" 1000.gz
 
LastUpdatedAt=$(psql "$2" -c "SELECT \"UpdatedAt\" FROM \"UsersRateTB\" ORDER BY \"UpdatedAt\" DESC  LIMIT 1;"  -t -A)
 
query1="
    COPY(SELECT * FROM \"UsersRateTB\" WHERE \"UpdatedAt\" >=$LastUpdatedAt) TO PROGRAM 'gzip > /var/lib/postgresql/tmp/UsersRateTB/$3'; 
"

query2=" 
    BEGIN;
    CREATE TEMP TABLE \"tmp_table\"
		ON COMMIT DROP
		AS
		SELECT * FROM \"UsersRateTB\" WITH NO DATA;

		COPY \"tmp_table\" FROM STDIN;

		INSERT INTO \"UsersRateTB\"
		SELECT * FROM \"tmp_table\" 

        ON CONFLICT (\"UserId\") DO UPDATE
            SET 
            \"UserName\"=excluded.\"UserName\",
            \"CollectedTweetsCount\"=excluded.\"CollectedTweetsCount\",
            \"LikesTweetsCount\"=excluded.\"LikesTweetsCount\",
            \"ReplyTweetsCount\"=excluded.\"ReplyTweetsCount\",
            \"RetweetedTweetsCount\"=excluded.\"RetweetedTweetsCount\",
            \"Languages\"=excluded.\"Languages\",
            \"primaryLanguage\"=excluded.\"primaryLanguage\",
            \"primaryLanguageCount\"=excluded.\"primaryLanguageCount\",
            \"GoldenFactor\"=excluded.\"GoldenFactor\",
            \"UpdatedAt\"=excluded.\"UpdatedAt\";
    COMMIT;
"
  
psql "$1" -c "$query1" 

sshpass -p domian12 scp -P 2222 root@204.93.216.157:/var/lib/postgresql/tmp/UsersRateTB/$3 /RED/tmp/UsersRateTB/ 

zcat /RED/tmp/UsersRateTB/$3 | psql "$2" -c "$query2"



