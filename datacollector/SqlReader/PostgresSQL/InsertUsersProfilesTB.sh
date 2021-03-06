#!/bin/bash 



query1="
    COPY(  
        WITH Removed_rows AS(
            DELETE FROM \"UsersProfilesTB\" 
            WHERE \"UpdatedAt\" <= (
                SELECT MAX(\"UpdatedAt\") FROM (
                    SELECT \"UpdatedAt\" FROM \"UsersProfilesTB\" order by \"UpdatedAt\" LIMIT $3
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
		SELECT * FROM \"UsersProfilesTB\" WITH NO DATA;

		COPY \"tmp_table\" FROM STDIN;

		INSERT INTO \"UsersProfilesTB\"
		SELECT \"tmp_table\".* FROM \"tmp_table\" 

		ON CONFLICT (\"UserId\") DO UPDATE
			SET
				\"UserId\"= COALESCE(excluded.\"UserId\",\"UsersProfilesTB\".\"UserId\"),
				\"UserName\"= COALESCE(excluded.\"UserName\",\"UsersProfilesTB\".\"UserName\"),
				\"ViewName\"= COALESCE(excluded.\"ViewName\",\"UsersProfilesTB\".\"ViewName\"),
				\"Bio\"= COALESCE(excluded.\"Bio\",\"UsersProfilesTB\".\"Bio\"),
				\"Location\"= COALESCE(excluded.\"Location\",\"UsersProfilesTB\".\"Location\"),
				\"Link\"= COALESCE(excluded.\"Link\",\"UsersProfilesTB\".\"Link\"),
				\"JoinDate\"= COALESCE(excluded.\"JoinDate\",\"UsersProfilesTB\".\"JoinDate\"),
				\"TweetsCount\"= COALESCE(excluded.\"TweetsCount\",\"UsersProfilesTB\".\"TweetsCount\"),
				\"LikesCount\"= COALESCE(excluded.\"LikesCount\",\"UsersProfilesTB\".\"LikesCount\"),
				\"FollowingCount\"= COALESCE(excluded.\"FollowingCount\",\"UsersProfilesTB\".\"FollowingCount\"),
				\"FollowersCount\"= COALESCE(excluded.\"FollowersCount\",\"UsersProfilesTB\".\"FollowersCount\"),
				\"BirthDate\"= COALESCE(excluded.\"BirthDate\",\"UsersProfilesTB\".\"BirthDate\"),
				\"UpdatedAt\"= COALESCE(excluded.\"UpdatedAt\",\"UsersProfilesTB\".\"UpdatedAt\"),
				\"LastTweetDate\"= COALESCE(excluded.\"LastTweetDate\",\"UsersProfilesTB\".\"LastTweetDate\"),
				\"LastTweetId\"= COALESCE(excluded.\"LastTweetId\",\"UsersProfilesTB\".\"LastTweetId\"),
				\"TwitterState\"= COALESCE(excluded.\"TwitterState\",\"UsersProfilesTB\".\"TwitterState\"),
				\"LastSource\"= COALESCE(excluded.\"LastSource\",\"UsersProfilesTB\".\"LastSource\");
    COMMIT;
"

psql "$1" -c "$query1"  | psql "$2" -c "$query2"