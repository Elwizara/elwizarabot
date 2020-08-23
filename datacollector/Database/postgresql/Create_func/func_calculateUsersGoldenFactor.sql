CREATE OR REPLACE FUNCTION calculateUsersGoldenFactor() 
RETURNS VOID AS
$$
DECLARE
    unixtimestamp integer;
BEGIN 
  	unixtimestamp = (CAST (extract(epoch from now()) AS INTEGER));
	
    INSERT INTO "UsersRateTB" ("UserId", "CollectedTweetsCount", "GoldenFactor","NeedToCrawl") 
    (
        SELECT 
			"UsersProfilesTB"."UserId" , 
			"Tweets"."CollectedTweetsCount" , 
			"Tweets"."GoldenFactor" ,  
			(
				(
						("Tweets"."CollectedTweetsCount" < 100)
					AND ("UsersProfilesTB"."TweetsCount" > 500)
					AND ("UsersProfilesTB"."UpdatedAt" > (unixtimestamp - 1209600)) --active with in 2 weeks
				) is TRUE
			) as "NeedToCrawl"
		
		FROM "UsersProfilesTB" 
		LEFT JOIN (
			SELECT 
				"UserId" , 
				AVG("SilverFactor") AS "GoldenFactor" ,
				Count("TweetId") AS "CollectedTweetsCount" 
			FROM "TweetsTB"
			GROUP BY "TweetsTB"."UserId" 
		) AS "Tweets" ON "UsersProfilesTB"."UserId" = "Tweets"."UserId" 
    )
    ON CONFLICT ("UserId") DO UPDATE 
        SET 
        "GoldenFactor" = excluded."GoldenFactor", 
        "CollectedTweetsCount" = excluded."CollectedTweetsCount",
		"NeedToCrawl" = excluded."NeedToCrawl";
 
END; $$
 
LANGUAGE plpgsql;