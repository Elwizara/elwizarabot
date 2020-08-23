#!/bin/bash 

query=" 
    BEGIN;
    CREATE TABLE IF NOT EXISTS public.\"UsersProfilesTB_$2\"
    (
        \"UserId\" bigint NOT NULL,
        \"UserName\" text COLLATE pg_catalog.\"default\",
        \"ViewName\" text COLLATE pg_catalog.\"default\",
        \"Bio\" text COLLATE pg_catalog.\"default\",
        \"Location\" text COLLATE pg_catalog.\"default\",
        \"Link\" text COLLATE pg_catalog.\"default\",
        \"JoinDate\" bigint,
        \"TweetsCount\" integer,
        \"LikesCount\" integer,
        \"FollowingCount\" integer,
        \"FollowersCount\" integer,
        \"BirthDate\" text COLLATE pg_catalog.\"default\",
        \"UpdatedAt\" bigint,
        \"LastTweetDate\" bigint,
        \"LastTweetId\" bigint,
        \"TwitterState\" integer,
        \"LastSource\" integer,
        CONSTRAINT \"UsersProfilesTB_$2_pkey\" PRIMARY KEY (\"UserId\")
    );
    CREATE INDEX IF NOT EXISTS \"idx_UsersProfilesTB_$2_UpdatedAt\"
        ON public.\"UsersProfilesTB_$2\" USING btree
        (\"UpdatedAt\")
        TABLESPACE pg_default;
    COMMIT;
"
psql "$1" -c "$query"
