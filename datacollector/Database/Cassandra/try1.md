>>CREATE KEYSPACE ElwizaraDEVDB WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
>>CREATE KEYSPACE ElwizaraLIVEDB WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

#TweetsTB
>>CREATE TABLE TweetsTB(
    TweetId bigint,
    UserId bigint,
    Favorite int,
    Reply int,
    Retweet int,
    SilverFactor bigint,
    UserName text,
    CreatedAt bigint,
    UpdatedAt bigint,
    Lang text,
    CountryCode text,
    RetweetedTweetId bigint,
    RetweetedUserId bigint,
    QuotedTweetId bigint,
    QuotedUserId bigint,
    ReplyTweetId bigint,
    ReplyUserId bigint,
    Source int,
    Text text,
    PRIMARY KEY ((UserId),UpdatedAt,TweetId)
) WITH CLUSTERING ORDER BY (UpdatedAt ASC,TweetId ASC);

>>select * from TweetsTB;
>>select * from TweetsTB where UserId = 1;
>>select * from TweetsTB where UserId = 1 and UpdatedAt <= 2015;
>>select * from TweetsTB where UserId = 1 and UpdatedAt = 2015 and TweetId <= 1 limit 3;
>>cqlsh -e "select * FROM elwizaradevdb.tweetstb limit 1";

#export psql 
>>psql "host=192.168.1.7 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable" -c 'select * from "TweetsTB" limit 10;' -F ',' -A -t  

>>psql "host=192.168.1.7 port=5432 user=tarek password=123 dbname=ElwizaraLIVEDB sslmode=disable" -c "select \"TweetId\",\"UserId\",\"Favorite\",\"Reply\",\"Retweet\",\"SilverFactor\",\"UserName\",\"CreatedAt\",\"UpdatedAt\",\"Lang\",\"CountryCode\",\"RetweetedTweetId\",\"RetweetedUserId\",\"QuotedTweetId\",\"QuotedUserId\",\"ReplyTweetId\",\"ReplyUserId\",\"Source\",regexp_replace(\"Text\", E'[\\n\\r\,]+', ' ', 'g' ) from \"TweetsTB\" limit 1000;" -F ',' -A -t  | cqlsh -e "COPY elwizaradevdb.TweetsTB (TweetId,UserId, Favorite, Reply, Retweet, SilverFactor, UserName, CreatedAt, UpdatedAt, Lang, CountryCode, RetweetedTweetId, RetweetedUserId, QuotedTweetId, QuotedUserId, ReplyTweetId, ReplyUserId, Source, Text) FROM STDIN;" 
 
#UsersProfilesTB
>>CREATE TABLE UsersProfilesTB  (
    UserId bigint,
    UserName text,
    ViewName text,
    Bio text,
    Location text,
    Link text,
    JoinDate bigint,
    TweetsCount int,
    LikesCount int,
    FollowingCount int,
    FollowersCount int,
    BirthDate text,
    UpdatedAt bigint,
    LastTweetDate bigint,
    LastTweetId bigint,
    TwitterState int,
    LastSource int,
    PRIMARY KEY (UserId)
);
>>select * from UsersProfilesTB;
>>select * from UsersProfilesTB where UserId = 1; 
>>cqlsh -e "select * FROM elwizaradevdb.UsersProfilesTB limit 1";


#psql "host=192.168.1.7 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" -c "SELECT \"UserId\",\"UserName\",regexp_replace(regexp_replace(\"ViewName\", E'[\\n\\r\,]+', ' ', 'g'),'\B', ' ', 'g'),regexp_replace(regexp_replace(\"Bio\",E'[\\n\\r\,]+', ' ', 'g'),'\B', ' ', 'g' ),regexp_replace(regexp_replace(\"Location\",E'[\\n\\r\,]+', ' ', 'g' ),'\B', ' ', 'g' ),regexp_replace(regexp_replace(\"Link\",E'[\\n\\r\,]+', ' ', 'g' ),'\B', ' ', 'g' ),\"JoinDate\",\"TweetsCount\",\"LikesCount\",\"FollowingCount\",\"FollowersCount\",regexp_replace(\"BirthDate\",E'[\\n\\r\,]+', ' ', 'g'),\"UpdatedAt\",\"LastTweetDate\",\"LastTweetId\",\"TwitterState\",\"LastSource\" FROM \"UsersProfilesTB\" LIMIT 1000000;" -F ',' -A -t  | cqlsh -e "COPY elwizaradevdb.UsersProfilesTB (UserId,UserName,ViewName,Bio,Location,Link,JoinDate,TweetsCount,LikesCount,FollowingCount,FollowersCount,BirthDate,UpdatedAt,LastTweetDate,LastTweetId,TwitterState,LastSource) FROM STDIN WITH DELIMITER=',';" 

  


psql "host=192.168.1.7 port=5432 user=tarek password=123 dbname=ElwizaraDEVDB sslmode=disable" -c "SELECT \"UserId\",\"UserName\",regexp_replace(regexp_replace(\"ViewName\", E'[\\n\\r\,]+', ' ', 'g'),'\B', ' ', 'g'),\"JoinDate\",\"TweetsCount\",\"LikesCount\",\"FollowingCount\",\"FollowersCount\",regexp_replace(\"BirthDate\",E'[\\n\\r\,]+', ' ', 'g'),\"UpdatedAt\",\"LastTweetDate\",\"LastTweetId\",\"TwitterState\",\"LastSource\" FROM \"UsersProfilesTB\" LIMIT 100000;" -F ',' -A -t  | cqlsh -e "COPY elwizaradevdb.UsersProfilesTB (UserId,UserName,ViewName,JoinDate,TweetsCount,LikesCount,FollowingCount,FollowersCount,BirthDate,UpdatedAt,LastTweetDate,LastTweetId,TwitterState,LastSource) FROM STDIN;" 

 