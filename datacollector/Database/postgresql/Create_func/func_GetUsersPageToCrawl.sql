
CREATE OR REPLACE FUNCTION public."GetUsersPageToCrawl"(pagesize integer)
    RETURNS TABLE(
        UserId bigint,
        CollectedTweetsCount integer ,
        GoldenFactor bigint ,
        NeedToCrawl boolean ,
        UpdatedAt bigint
    ) 
    LANGUAGE 'plpgsql'

    COST 100
    VOLATILE 
    ROWS 1000
AS $BODY$

BEGIN
 
    CREATE TEMP TABLE "tmp_table"
		ON COMMIT DROP
		AS
		SELECT * FROM "UsersRateTB" WITH NO DATA;
	
	INSERT INTO "tmp_table"
		SELECT * FROM  "UsersRateTB"
			WHERE "NeedToCrawl" = TRUE
			ORDER BY "UpdatedAt" ASC
			LIMIT pagesize;
			 
	UPDATE "UsersRateTB" 
		SET "NeedToCrawl" = FALSE
		WHERE "UserId" in (SELECT "UserId" FROM "tmp_table");
	 
RETURN QUERY	
	SELECT 
            "UserId",
            "CollectedTweetsCount",
            "GoldenFactor",
            "NeedToCrawl",
            "UpdatedAt"
        FROM "tmp_table";
	
END

$BODY$; 
