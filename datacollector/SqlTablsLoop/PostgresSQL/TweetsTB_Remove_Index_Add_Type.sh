#!/bin/bash 

query=" 
    BEGIN; 
    DROP INDEX IF EXISTS \"idx_TweetsTB_$2_UpdatedAt\" ;    
    ALTER TABLE public.\"TweetsTB_$2\" ADD COLUMN \"Type\" json;
    COMMIT;
"
psql "$1" -c "$query"
