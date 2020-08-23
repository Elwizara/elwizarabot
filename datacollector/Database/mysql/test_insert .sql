


INSERT INTO `Twitter_Stream_DB`.`Users_Distinct_TB`
(id,screen_name,verified,protected,followers_count,
following_count,listed_count,favourites_count,tweets_count,utc_offset,
time_zone,geo_enabled,Lang)
(SELECT Twitter_Id as Id,screen_name,verified,protected,followers_count,
following_count,listed_count,favourites_count,tweets_count,utc_offset,
time_zone,geo_enabled,Lang
FROM `Twitter_Stream_DB`.`Users_TB` where Id is not null and Twitter_Id is not null)

ON DUPLICATE KEY UPDATE 
	`Id`=VALUES(`Id`), 
	`screen_name`=VALUES(`screen_name`), 
	`verified`=VALUES(`verified`), 
	`protected`=VALUES(`protected`), 
	`followers_count`=VALUES(`followers_count`), 
	`following_count`=VALUES(`following_count`), 
	`favourites_count`=VALUES(`favourites_count`), 
	`tweets_count`=VALUES(`tweets_count`), 
	`utc_offset`=VALUES(`utc_offset`), 
	`time_zone`=VALUES(`time_zone`), 
	`geo_enabled`=VALUES(`geo_enabled`), 
	`Lang`=VALUES(`Lang`);



select count(*) from Users_TB; 
select count(*) from Users_Distinct_TB;
select * from Users_TB;
select * from Users_Distinct_TB; 
truncate table test;
truncate table Users_Distinct_TB; 


-- Users_TB 		 '1176951' - '1201229' = 24278
-- Users_Distinct_TB '277496'  - '299260'  = 21764















