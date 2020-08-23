#!/bin/bash 

query=" 
    BEGIN;
    CREATE TABLE \"TweetsTB_$2\" AS  
    SELECT * FROM \"TweetsTB\" WHERE (\"UserId\"%1000) = $2; 
    COMMIT;
" 
psql "$1" -c "$query"
