------------------------------------------
I.REQUIREMENTS
------------------------------------------

1. linux/ubuntu `sys` 
2. GOLANG       `StreamTwitter`
3. DATABASE     `postgresql`
 


 
------------------------------------------
II. BUILDING AND RUNNING FOR DEV
------------------------------------------ 

___________________
1. run Stream Twitter

from twitter apps create your app with tokens `https://apps.twitter.com/`
 
>>./StreamTwitterGo \
-c "{{your consumerKey}}" \
-cS "{{your consumerSecret}}" \
-a "{{your accessToken}}" \
-aS "{{your accessTokenSecret}}" \
-db "host={{host}} port={{port}} user={{user}} password={{password}} dbname={{dbname}} sslmode=disable"

to build with golang 
in StreamTwitterGo dir `~/StreamTwitterGo`
>>go build . 
 

------------------------------------------------------------------------------------
    * STREAM SAMPLE DATA FROM TWITTER.COM
------------------------------------------------------------------------------------




THANK YOU.
By:Tarek Badr
tarekbadrshalaan@gmail.com
00201273486662