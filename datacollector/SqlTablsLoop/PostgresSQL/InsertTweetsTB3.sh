#!/bin/bash 

query=" 
    BEGIN; 
    INSERT INTO \"TweetsTB_$2\"
    SELECT * FROM \"TweetsTB\" WHERE (\"UserId\"%1000) = $2;
    COMMIT;
" 
psql "$1" -c "$query"
