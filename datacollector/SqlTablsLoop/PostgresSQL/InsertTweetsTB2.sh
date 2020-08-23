#!/bin/bash 

query=" 
    BEGIN;
    WITH moved_rows AS(
        DELETE FROM \"TweetsTB\"
        WHERE (\"UserId\"%1000) = $2
        RETURNING *
    )
    INSERT INTO \"TweetsTB_$2\"
    SELECT * FROM moved_rows;
    COMMIT;
" 
psql "$1" -c "$query"
