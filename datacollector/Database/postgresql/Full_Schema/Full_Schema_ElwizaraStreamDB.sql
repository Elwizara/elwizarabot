--
-- PostgreSQL database dump
--

-- Dumped from database version 10.4 (Ubuntu 10.4-1.pgdg16.04+1)
-- Dumped by pg_dump version 10.4 (Ubuntu 10.4-2.pgdg16.04+1)

-- Started on 2018-08-09 13:28:57 EET

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12964)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 4726 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 2075 (class 1255 OID 45083)
-- Name: GetUsersPageToCrawl(integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."GetUsersPageToCrawl"(pagesize integer) RETURNS TABLE(userid bigint)
    LANGUAGE plpgsql
    AS $$
BEGIN
 
	CREATE TEMP TABLE "tmp_table" ("UserId"	bigint)
		ON COMMIT DROP ;
	
	INSERT INTO "tmp_table"
		SELECT "UserId" FROM  "UsersRateTB"
			WHERE "NeedToCrawl" = TRUE
			ORDER BY "UpdatedAt" ASC
			LIMIT pagesize;
			 
	UPDATE "UsersRateTB" 
		SET "NeedToCrawl" = FALSE
		WHERE "UserId" in (SELECT "UserId" FROM "tmp_table");
	 
RETURN QUERY	
	SELECT "UserId" FROM "tmp_table";
	
END
$$;


ALTER FUNCTION public."GetUsersPageToCrawl"(pagesize integer) OWNER TO postgres;

--
-- TOC entry 2061 (class 1255 OID 291083)
-- Name: InsertGoldenTweetsTB(bigint, bigint, text, integer, integer, integer, text, text, integer, integer, integer, bigint, bigint, boolean); Type: FUNCTION; Schema: public; Owner: elwizara
--

CREATE FUNCTION public."InsertGoldenTweetsTB"(tweetid bigint, userid bigint, username text, favorite integer, reply integer, retweet integer, tweetlanguage text, userprimarylanguage text, tweetsilverfactor integer, usergoldenfactor integer, measurerate integer, createdat bigint, updatedat bigint, expired boolean) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
		INSERT INTO "GoldenTweetsTB"(
            "TweetId",
            "UserId",
            "UserName",
            "Favorite",
            "Reply",
            "Retweet",
            "TweetLanguage",
            "UserPrimaryLanguage",
            "TweetSilverFactor",
            "UserGoldenFactor",
            "MeasureRate",
            "CreatedAt",
            "UpdatedAt",
            "Expired" )
        VALUES (
			TweetId,
            UserId,
            UserName,
            Favorite,
            Reply,
            Retweet,
            TweetLanguage,
            UserPrimaryLanguage,
            TweetSilverFactor,
            UserGoldenFactor,
            MeasureRate,
            CreatedAt,
            UpdatedAt,
            Expired )
		ON CONFLICT ("TweetId") DO UPDATE 
			SET    
                "UserId" = excluded."UserId",
                "UserName" = excluded."UserName",
                "Favorite" = excluded."Favorite",
                "Reply" = excluded."Reply",
                "Retweet" = excluded."Retweet",
                "TweetLanguage" = excluded."TweetLanguage",
                "UserPrimaryLanguage" = excluded."UserPrimaryLanguage",
                "TweetSilverFactor" = excluded."TweetSilverFactor",
                "UserGoldenFactor" = excluded."UserGoldenFactor",
                "MeasureRate" = excluded."MeasureRate",
                "CreatedAt" = excluded."CreatedAt",
                "UpdatedAt" = excluded."UpdatedAt",
                "Expired" = excluded."Expired";
                
END
$$;


ALTER FUNCTION public."InsertGoldenTweetsTB"(tweetid bigint, userid bigint, username text, favorite integer, reply integer, retweet integer, tweetlanguage text, userprimarylanguage text, tweetsilverfactor integer, usergoldenfactor integer, measurerate integer, createdat bigint, updatedat bigint, expired boolean) OWNER TO elwizara;

--
-- TOC entry 2060 (class 1255 OID 45081)
-- Name: InsertTweetsTB(bigint, bigint, bigint, bigint, bigint, bigint, text, bigint, bigint, text, text, bigint, bigint, bigint, bigint, bigint, bigint, integer, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."InsertTweetsTB"(tweetid bigint, userid bigint, favorite bigint, reply bigint, retweet bigint, silverfactor bigint, username text, createdat bigint, updatedat bigint, lang text, countrycode text, retweetedtweetid bigint, retweeteduserid bigint, quotedtweetid bigint, quoteduserid bigint, replytweetid bigint, replyuserid bigint, source integer, tweettext text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN 
	INSERT INTO "TweetsTB"(
					"TweetId",
					"UserId",
					"Favorite",
					"Reply",
					"Retweet",
					"SilverFactor",
					"UserName",
					"CreatedAt",
					"UpdatedAt",
					"Lang",
					"CountryCode",
					"RetweetedTweetId",
					"RetweetedUserId",
					"QuotedTweetId",
					"QuotedUserId",
					"ReplyTweetId",
					"ReplyUserId",
					"Source",
					"Text"
				)
		VALUES (
				TweetId,
				UserId,
				Favorite,
				Reply,
				Retweet,
				SilverFactor,
				UserName,
				CreatedAt,
				UpdatedAt,
				Lang,
				CountryCode,
				RetweetedTweetId,
				RetweetedUserId,
				QuotedTweetId,
				QuotedUserId,
				ReplyTweetId,
				ReplyUserId,
				Source,
				Tweettext
			)
		ON CONFLICT ("TweetId") DO UPDATE 
		  SET  
			"UserId" = COALESCE(excluded."UserId","TweetsTB"."UserId"),
			"Favorite" = COALESCE(excluded."Favorite","TweetsTB"."Favorite"),
			"Reply" = COALESCE(excluded."Reply","TweetsTB"."Reply"),
			"Retweet" = COALESCE(excluded."Retweet","TweetsTB"."Retweet"),
			"SilverFactor" = COALESCE(excluded."SilverFactor","TweetsTB"."SilverFactor"),
			"UserName" = COALESCE(excluded."UserName","TweetsTB"."UserName"),
			"CreatedAt" = COALESCE(excluded."CreatedAt","TweetsTB"."CreatedAt"),
			"UpdatedAt" = COALESCE(excluded."UpdatedAt","TweetsTB"."UpdatedAt"),
			"Lang" = COALESCE(excluded."Lang","TweetsTB"."Lang"),
			"CountryCode" = COALESCE(excluded."CountryCode","TweetsTB"."CountryCode"),
			"RetweetedTweetId" = COALESCE(excluded."RetweetedTweetId","TweetsTB"."RetweetedTweetId"),
			"RetweetedUserId" = COALESCE(excluded."RetweetedUserId","TweetsTB"."RetweetedUserId"),
			"QuotedTweetId" = COALESCE(excluded."QuotedTweetId","TweetsTB"."QuotedTweetId"),
			"QuotedUserId" = COALESCE(excluded."QuotedUserId","TweetsTB"."QuotedUserId"),
			"ReplyTweetId" = COALESCE(excluded."ReplyTweetId","TweetsTB"."ReplyTweetId"),
			"ReplyUserId" = COALESCE(excluded."ReplyUserId","TweetsTB"."ReplyUserId"),
			"Source" = COALESCE(excluded."Source","TweetsTB"."Source"),
			"Text" = COALESCE(excluded."Text","TweetsTB"."Text");
END
$$;


ALTER FUNCTION public."InsertTweetsTB"(tweetid bigint, userid bigint, favorite bigint, reply bigint, retweet bigint, silverfactor bigint, username text, createdat bigint, updatedat bigint, lang text, countrycode text, retweetedtweetid bigint, retweeteduserid bigint, quotedtweetid bigint, quoteduserid bigint, replytweetid bigint, replyuserid bigint, source integer, tweettext text) OWNER TO postgres;

--
-- TOC entry 2072 (class 1255 OID 45082)
-- Name: InsertUsersProfileTB(bigint, text, text, text, text, text, bigint, integer, integer, integer, integer, text, bigint, bigint, bigint, integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."InsertUsersProfileTB"(userid bigint, username text, viewname text, bio text, location text, link text, joindate bigint, tweetscount integer, likescount integer, followingcount integer, followerscount integer, birthdate text, updatedat bigint, lasttweetdate bigint, lasttweetid bigint, twitterstate integer, lastsource integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION public."InsertUsersProfileTB"(userid bigint, username text, viewname text, bio text, location text, link text, joindate bigint, tweetscount integer, likescount integer, followingcount integer, followerscount integer, birthdate text, updatedat bigint, lasttweetdate bigint, lasttweetid bigint, twitterstate integer, lastsource integer) OWNER TO postgres;

--
-- TOC entry 2062 (class 1255 OID 262619)
-- Name: InsertUsersRateTB(bigint, text, integer, integer, integer, integer, text, text, integer, bigint, bigint); Type: FUNCTION; Schema: public; Owner: elwizara
--

CREATE FUNCTION public."InsertUsersRateTB"(userid bigint, username text, collectedtweetscount integer, likestweetscount integer, replytweetscount integer, retweetedtweetscount integer, languages text, primarylanguage text, primarylanguagecount integer, goldenfactor bigint, updatedat bigint) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
		INSERT INTO "UsersRateTB"(
			"UserId",
            "UserName",
            "CollectedTweetsCount",
            "LikesTweetsCount",
            "ReplyTweetsCount",
            "RetweetedTweetsCount",
            "Languages",
            "primaryLanguage",
            "primaryLanguageCount",
            "GoldenFactor",
            "UpdatedAt")
		VALUES (
			UserId,
            UserName,
            CollectedTweetsCount,
            LikesTweetsCount,
            ReplyTweetsCount,
            RetweetedTweetsCount,
            Languages,
            primaryLanguage,
            primaryLanguageCount,
            GoldenFactor,
            UpdatedAt)
		ON CONFLICT ("UserId") DO UPDATE 
			SET  
                "UserName" = excluded."UserName",
                "CollectedTweetsCount" = excluded."CollectedTweetsCount",
                "LikesTweetsCount" = excluded."LikesTweetsCount",
                "ReplyTweetsCount" = excluded."ReplyTweetsCount",
                "RetweetedTweetsCount" = excluded."RetweetedTweetsCount",
                "Languages" = excluded."Languages",
                "primaryLanguage" = excluded."primaryLanguage",
                "primaryLanguageCount" = excluded."primaryLanguageCount",
                "GoldenFactor" = excluded."GoldenFactor",
                "UpdatedAt" = excluded."UpdatedAt";
END
$$;


ALTER FUNCTION public."InsertUsersRateTB"(userid bigint, username text, collectedtweetscount integer, likestweetscount integer, replytweetscount integer, retweetedtweetscount integer, languages text, primarylanguage text, primarylanguagecount integer, goldenfactor bigint, updatedat bigint) OWNER TO elwizara;

--
-- TOC entry 2071 (class 1255 OID 45080)
-- Name: calculateUsersGoldenFactor(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public."calculateUsersGoldenFactor"() RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
    unixtimestamp integer;
BEGIN 
   unixtimestamp = (CAST (extract(epoch from now()) AS INTEGER));
 
    INSERT INTO "UsersRateTB" ("UserId", "CollectedTweetsCount", "GoldenFactor","UpdatedAt","NeedToCrawl") 
    (
        SELECT 
   "UsersProfilesTB"."UserId" , 
   COALESCE("Tweets"."CollectedTweetsCount",0) AS "CollectedTweetsCount", 
   COALESCE("Tweets"."GoldenFactor",0) AS "GoldenFactor",  
   COALESCE("UsersProfilesTB"."UpdatedAt",0) AS "UpdatedAt",
   (
    (
      ("Tweets"."CollectedTweetsCount" < 100)
     AND ("UsersProfilesTB"."TwitterState" = 0)
     AND ("UsersProfilesTB"."TweetsCount" > 500)
     AND ("UsersProfilesTB"."UpdatedAt" > (unixtimestamp - 1209600)) --active with in 2 weeks
    ) is TRUE
   ) AS "NeedToCrawl" 
  
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
 
END; $$;


ALTER FUNCTION public."calculateUsersGoldenFactor"() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 2057 (class 1259 OID 291075)
-- Name: GoldenTweetsTB; Type: TABLE; Schema: public; Owner: elwizara
--

CREATE TABLE public."GoldenTweetsTB" (
    "TweetId" bigint NOT NULL,
    "UserId" bigint NOT NULL,
    "UserName" text NOT NULL,
    "Favorite" integer NOT NULL,
    "Reply" integer NOT NULL,
    "Retweet" integer NOT NULL,
    "TweetLanguage" text NOT NULL,
    "UserPrimaryLanguage" text NOT NULL,
    "TweetSilverFactor" integer NOT NULL,
    "UserGoldenFactor" integer NOT NULL,
    "MeasureRate" integer NOT NULL,
    "CreatedAt" bigint NOT NULL,
    "UpdatedAt" bigint NOT NULL,
    "Expired" boolean NOT NULL
);


ALTER TABLE public."GoldenTweetsTB" OWNER TO elwizara;

--
-- TOC entry 2054 (class 1259 OID 16406)
-- Name: TweetsTB; Type: TABLE; Schema: public; Owner: elwizara
--

CREATE TABLE public."TweetsTB" (
    "TweetId" bigint NOT NULL,
    "UserId" bigint NOT NULL,
    "Favorite" integer NOT NULL,
    "Reply" integer NOT NULL,
    "Retweet" integer NOT NULL,
    "SilverFactor" bigint NOT NULL,
    "UserName" text,
    "CreatedAt" bigint,
    "UpdatedAt" bigint,
    "Lang" text,
    "CountryCode" text,
    "RetweetedTweetId" bigint,
    "RetweetedUserId" bigint,
    "QuotedTweetId" bigint,
    "QuotedUserId" bigint,
    "ReplyTweetId" bigint,
    "ReplyUserId" bigint,
    "Source" smallint,
    "Text" text
);


ALTER TABLE public."TweetsTB" OWNER TO elwizara;

--
-- TOC entry 2055 (class 1259 OID 24697)
-- Name: UsersProfilesTB; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."UsersProfilesTB" (
    "UserId" bigint NOT NULL,
    "UserName" text,
    "ViewName" text,
    "Bio" text,
    "Location" text,
    "Link" text,
    "JoinDate" bigint,
    "TweetsCount" integer,
    "LikesCount" integer,
    "FollowingCount" integer,
    "FollowersCount" integer,
    "BirthDate" text,
    "UpdatedAt" bigint,
    "LastTweetDate" bigint,
    "LastTweetId" bigint,
    "TwitterState" integer,
    "LastSource" integer
);


ALTER TABLE public."UsersProfilesTB" OWNER TO postgres;

--
-- TOC entry 2056 (class 1259 OID 262620)
-- Name: UsersRateTB; Type: TABLE; Schema: public; Owner: elwizara
--

CREATE TABLE public."UsersRateTB" (
    "UserId" bigint NOT NULL,
    "UserName" text,
    "CollectedTweetsCount" integer,
    "LikesTweetsCount" integer,
    "ReplyTweetsCount" integer,
    "RetweetedTweetsCount" integer,
    "Languages" text,
    "primaryLanguage" text,
    "primaryLanguageCount" integer,
    "GoldenFactor" bigint,
    "UpdatedAt" bigint
);


ALTER TABLE public."UsersRateTB" OWNER TO elwizara;

--
-- TOC entry 4597 (class 2606 OID 291082)
-- Name: GoldenTweetsTB GoldenTweetsTB_pkey; Type: CONSTRAINT; Schema: public; Owner: elwizara
--

ALTER TABLE ONLY public."GoldenTweetsTB"
    ADD CONSTRAINT "GoldenTweetsTB_pkey" PRIMARY KEY ("TweetId");


--
-- TOC entry 4588 (class 2606 OID 16416)
-- Name: TweetsTB TweetsTB_pkey; Type: CONSTRAINT; Schema: public; Owner: elwizara
--

ALTER TABLE ONLY public."TweetsTB"
    ADD CONSTRAINT "TweetsTB_pkey" PRIMARY KEY ("TweetId");


--
-- TOC entry 4592 (class 2606 OID 24705)
-- Name: UsersProfilesTB UsersProfilesTB_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."UsersProfilesTB"
    ADD CONSTRAINT "UsersProfilesTB_pkey" PRIMARY KEY ("UserId");


--
-- TOC entry 4595 (class 2606 OID 262627)
-- Name: UsersRateTB UsersRateTB_pkey; Type: CONSTRAINT; Schema: public; Owner: elwizara
--

ALTER TABLE ONLY public."UsersRateTB"
    ADD CONSTRAINT "UsersRateTB_pkey" PRIMARY KEY ("UserId");


--
-- TOC entry 4589 (class 1259 OID 16419)
-- Name: idx_TweestsTBUpdatedAt; Type: INDEX; Schema: public; Owner: elwizara
--

CREATE INDEX "idx_TweestsTBUpdatedAt" ON public."TweetsTB" USING btree ("UpdatedAt");


--
-- TOC entry 4590 (class 1259 OID 16420)
-- Name: idx_TweestsTBuserid; Type: INDEX; Schema: public; Owner: elwizara
--

CREATE INDEX "idx_TweestsTBuserid" ON public."TweetsTB" USING hash ("UserId");


--
-- TOC entry 4593 (class 1259 OID 25077)
-- Name: idx_UsersProfilesTBUpdatedAT; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "idx_UsersProfilesTBUpdatedAT" ON public."UsersProfilesTB" USING btree ("UpdatedAt");


-- Completed on 2018-08-09 13:29:23 EET

--
-- PostgreSQL database dump complete
--

