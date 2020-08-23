#!/bin/bash 
#./UsersPageToCrawl.sh "host=127.0.0.1 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" 10 "host=elwizaralivecloud1.postgres.database.azure.com port=5432 user=elwizaracloud1@elwizaralivecloud1 password=A9Sd42^d%KjrG3MK dbname=ElwizaraLIVEDB"

query="  
    BEGIN;
    CREATE TEMP TABLE \"tmp_table\"
		ON COMMIT DROP
		AS
		SELECT * FROM \"UsersRateTB\" WITH NO DATA;

		COPY \"tmp_table\" FROM STDIN;
 
        INSERT INTO \"UsersRateTB\"
            SELECT \"tmp_table\".* FROM \"tmp_table\" 
            ON CONFLICT (\"UserId\") DO UPDATE
            SET  
                \"UserId\" = excluded.\"UserId\",
                \"UpdatedAt\" = excluded.\"UpdatedAt\",
                \"NeedToCrawl\" = TRUE ;
    COMMIT;
"

psql "$1" -c "COPY(select * from \"GetUsersPageToCrawl\"($2)) TO STDOUT;" | psql "$3" -c "$query"
 	 