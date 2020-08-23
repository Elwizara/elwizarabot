#!/bin/bash 
#./CountTweetsTB.sh 

RESULT=0 

for value in {0..999}
do
    query="SELECT COUNT(*) FROM \"TweetsTB_$value\";" 
    TCOUNT=$(psql "$1" -c "$query" -t -A)
    echo $value : $TCOUNT
    RESULT=$(($RESULT + $TCOUNT))
done

echo $RESULT
