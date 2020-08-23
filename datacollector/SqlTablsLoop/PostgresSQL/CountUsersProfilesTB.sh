#!/bin/bash 
#./CountUsersProfilesTB.sh  "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable"

RESULT=0 

for value in {0..999}
do
    query="SELECT COUNT(*) FROM \"UsersProfilesTB_$value\";" 
    TCOUNT=$(psql "$1" -c "$query" -t -A)
    echo $value : $TCOUNT
    RESULT=$(($RESULT + $TCOUNT))
done

echo $RESULT
