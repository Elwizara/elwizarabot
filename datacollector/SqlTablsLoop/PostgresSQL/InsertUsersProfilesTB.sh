#!/bin/bash 

query=" 
    BEGIN; 
    INSERT INTO \"UsersProfilesTB_$2\"
    SELECT * FROM \"UsersProfilesTB\" WHERE (\"UserId\"%1000) = $2;
    COMMIT;
" 
psql "$1" -c "$query"
