#!/bin/bash   
#./PostgresSQL/InsertUsersProfilesTB3.sh "host=204.93.216.157 port=5432 user=elwizara password=${pass} dbname=ElwizaraLIVEDB" "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable" 1530008990.gz 1530008990

dollarSign='$$'

query1="        
    COPY \"UsersProfilesTB\" TO PROGRAM 'gzip > /var/lib/postgresql/tmp/UsersProfilesTB/$3';
    TRUNCATE TABLE \"UsersProfilesTB\";
"



query2="DO $dollarSign
        DECLARE query text;
        BEGIN 
        FOR counter IN 0..999 LOOP 
            query:= 'INSERT INTO \"UsersProfilesTB_'||counter||'\"  SELECT * FROM \"UsersProfilesTB\" WHERE (\"UserId\"%1000)=' || counter;
            query := query || '        ON CONFLICT (\"UserId\") DO UPDATE
                    SET
                        \"UserId\"= COALESCE(excluded.\"UserId\",\"UsersProfilesTB_'||counter||'\".\"UserId\"),
                        \"UserName\"= COALESCE(excluded.\"UserName\",\"UsersProfilesTB_'||counter||'\".\"UserName\"),
                        \"ViewName\"= COALESCE(excluded.\"ViewName\",\"UsersProfilesTB_'||counter||'\".\"ViewName\"),
                        \"Bio\"= COALESCE(excluded.\"Bio\",\"UsersProfilesTB_'||counter||'\".\"Bio\"),
                        \"Location\"= COALESCE(excluded.\"Location\",\"UsersProfilesTB_'||counter||'\".\"Location\"),
                        \"Link\"= COALESCE(excluded.\"Link\",\"UsersProfilesTB_'||counter||'\".\"Link\"),
                        \"JoinDate\"= COALESCE(excluded.\"JoinDate\",\"UsersProfilesTB_'||counter||'\".\"JoinDate\"),
                        \"TweetsCount\"= COALESCE(excluded.\"TweetsCount\",\"UsersProfilesTB_'||counter||'\".\"TweetsCount\"),
                        \"LikesCount\"= COALESCE(excluded.\"LikesCount\",\"UsersProfilesTB_'||counter||'\".\"LikesCount\"),
                        \"FollowingCount\"= COALESCE(excluded.\"FollowingCount\",\"UsersProfilesTB_'||counter||'\".\"FollowingCount\"),
                        \"FollowersCount\"= COALESCE(excluded.\"FollowersCount\",\"UsersProfilesTB_'||counter||'\".\"FollowersCount\"),
                        \"BirthDate\"= COALESCE(excluded.\"BirthDate\",\"UsersProfilesTB_'||counter||'\".\"BirthDate\"),
                        \"UpdatedAt\"= COALESCE(excluded.\"UpdatedAt\",\"UsersProfilesTB_'||counter||'\".\"UpdatedAt\"),
                        \"LastTweetDate\"= COALESCE(excluded.\"LastTweetDate\",\"UsersProfilesTB_'||counter||'\".\"LastTweetDate\"),
                        \"LastTweetId\"= COALESCE(excluded.\"LastTweetId\",\"UsersProfilesTB_'||counter||'\".\"LastTweetId\"),
                        \"TwitterState\"= COALESCE(excluded.\"TwitterState\",\"UsersProfilesTB_'||counter||'\".\"TwitterState\"),
                        \"LastSource\"= COALESCE(excluded.\"LastSource\",\"UsersProfilesTB_'||counter||'\".\"LastSource\");';
            EXECUTE query;
        END LOOP;
        END; $dollarSign
"

psql "$2" -c "TRUNCATE TABLE \"UsersProfilesTB\";"  

psql "$1" -c "$query1" 

sshpass -p domian12 scp -P 2222 root@204.93.216.157:/var/lib/postgresql/tmp/UsersProfilesTB/$3 /RED/tmp/UsersProfilesTB

zcat /RED/tmp/UsersProfilesTB/$3 | psql "$2" -c " COPY \"UsersProfilesTB\" FROM STDIN;"

psql "$2" -c "$query2"

psql "$2" -c "TRUNCATE TABLE \"UsersProfilesTB\";"  

#select count(*) from "UsersProfilesTB"; server:4043863

# select count(*) from "UsersProfilesTB_0";--"66906"
# select count(*) from "UsersProfilesTB_1";--"32541"
# select count(*) from "UsersProfilesTB_2";--"22892"
# select count(*) from "UsersProfilesTB_3";--"16052"
# select count(*) from "UsersProfilesTB_4";--"18191"
# select count(*) from "UsersProfilesTB_5";--"15741"
# select count(*) from "UsersProfilesTB_100";--"18922"
# select count(*) from "UsersProfilesTB_588";--"18554"
# select count(*) from "UsersProfilesTB_700";--"18637"
# select count(*) from "UsersProfilesTB_980";--"18987"
# select count(*) from "UsersProfilesTB_998";--"18234"