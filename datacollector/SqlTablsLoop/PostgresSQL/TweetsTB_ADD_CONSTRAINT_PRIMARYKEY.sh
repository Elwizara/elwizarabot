#!/bin/bash 

query=" 
    BEGIN; 
    ALTER TABLE public.\"TweetsTB_$2\" DROP CONSTRAINT IF EXISTS \"TweetsTB_$2_pkey\" ;    
    ALTER TABLE public.\"TweetsTB_$2\" ADD CONSTRAINT \"TweetsTB_$2_pkey\" PRIMARY KEY (\"TweetId\");
    COMMIT;
"
psql "$1" -c "$query"
