#!/bin/bash 

query=" 
    BEGIN;  
    DELETE FROM \"TweetsTB\" WHERE (\"UserId\"%1000) = $2; 
    COMMIT;
" 
psql "$1" -c "$query"
