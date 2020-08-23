CREATE FUNCTION "InsertUsersProfileTB"
	(
		UserId bigint,
		UserName text,
		ViewName text,
		Bio text,
		Location text,
		Link text,
		JoinDate bigint,
		TweetsCount integer,
		LikesCount integer,
		FollowingCount integer,
		FollowersCount integer,
		BirthDate text,  
		UpdatedAt bigint,
		LastTweetDate bigint,
		LastTweetId bigint, 
		TwitterState integer,
		LastSource integer
	) 
RETURNS VOID AS
$$
BEGIN
		INSERT INTO "UsersProfilesTB"(
			"UserId",
			"UserName",
			"ViewName",
			"Bio",
			"Location",
			"Link",
			"JoinDate",
			"TweetsCount",
			"LikesCount",
			"FollowingCount",
			"FollowersCount",
			"BirthDate",  
			"UpdatedAt",
			"LastTweetDate",
			"LastTweetId", 
			"TwitterState",
			"LastSource"
		)
		VALUES (
			UserId,
			UserName,
			ViewName,
			Bio,
			Location,
			Link,
			JoinDate,
			TweetsCount,
			LikesCount,
			FollowingCount,
			FollowersCount,
			BirthDate,  
			UpdatedAt,
			LastTweetDate,
			LastTweetId, 
			TwitterState,
			LastSource
		)
		ON CONFLICT ("UserId") DO UPDATE 
			SET 
				"UserId"= COALESCE(excluded."UserId","UsersProfilesTB"."UserId"),
				"UserName"= COALESCE(excluded."UserName","UsersProfilesTB"."UserName"),
				"ViewName"= COALESCE(excluded."ViewName","UsersProfilesTB"."ViewName"),
				"Bio"= COALESCE(excluded."Bio","UsersProfilesTB"."Bio"),
				"Location"= COALESCE(excluded."Location","UsersProfilesTB"."Location"),
				"Link"= COALESCE(excluded."Link","UsersProfilesTB"."Link"),
				"JoinDate"= COALESCE(excluded."JoinDate","UsersProfilesTB"."JoinDate"),
				"TweetsCount"= COALESCE(excluded."TweetsCount","UsersProfilesTB"."TweetsCount"),
				"LikesCount"= COALESCE(excluded."LikesCount","UsersProfilesTB"."LikesCount"),
				"FollowingCount"= COALESCE(excluded."FollowingCount","UsersProfilesTB"."FollowingCount"),
				"FollowersCount"= COALESCE(excluded."FollowersCount","UsersProfilesTB"."FollowersCount"),
				"BirthDate"= COALESCE(excluded."BirthDate","UsersProfilesTB"."BirthDate"),  
				"UpdatedAt"= COALESCE(excluded."UpdatedAt","UsersProfilesTB"."UpdatedAt"),
				"LastTweetDate"= COALESCE(excluded."LastTweetDate","UsersProfilesTB"."LastTweetDate"),
				"LastTweetId"= COALESCE(excluded."LastTweetId","UsersProfilesTB"."LastTweetId"), 
				"TwitterState"= COALESCE(excluded."TwitterState","UsersProfilesTB"."TwitterState"),
				"LastSource"= COALESCE(excluded."LastSource","UsersProfilesTB"."LastSource");
END
$$
LANGUAGE 'plpgsql';


 