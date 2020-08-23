


#   http://www.cibeg.com/English/Personal/Cards/AWorldofDiscounts/Pages/Default.aspx
#   StartUpFile       : C:\Users\Tarek\Anaconda2\Lib\site-packages\scrapy\cmdline.py
#   Working Directory : E:\Develop\BigData\Projects\Fidelyo\Crawler\Crawler_Portal\Fidelyo_Crawler\Fidelyo_Crawler
#   Script Arguments  : crawl userprofileSpider -o userprofileSpider.jl



import os
import sys
import time
import json
from datetime import datetime

import scrapy

from Twitter_Crawler.items import TweetsItem,UserItem
from Twitter_Crawler.Helper.lxmlParserHelper import lxmlParserHelper as parser

class userprofileSpider(scrapy.Spider):
    name = "userprofileSpider"

    def start_requests(self):
        try:
            with open('Twitter_Crawler/conf.json') as f:
                configuration  = f.read()
            configuration_obj = json.loads(configuration)
            crawlermonitorURL = configuration_obj.get("CrawlerMonitorURL")+"/getpage"
            yield scrapy.Request(crawlermonitorURL, self.parsePage)

        except Exception as e :
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            self.logger.error('start_requests! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))
            print('start_requests! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

    def parsePage(self,response):
        try:
            print('=================::::::::==================== A response from %s just arrived!',response.url)
            self.logger.info('=================::::::::==================== A response from %s just arrived!',response.url)
            users = json.loads(response.body)
            for u in users:
                username = u.get("UserName",None)
                userId = u.get("UserId",None)
                if username != None and username != "":
                    url = "https://twitter.com/"+username
                    request = scrapy.Request(url, self.parse,meta={'handle_httpstatus_list': [301,302,404]})
                    request.meta["userId"] = userId
                    request.meta["username"] = username
                    yield request
                elif userId != None and userId > 0:
                    url = "https://twitter.com/intent/user?user_id="+str(userId)
                    request = scrapy.Request(url, self.parseintent,meta={'handle_httpstatus_list': [301,302,404]})
                    request.meta["userId"] = userId
                    yield request

        except Exception as e :
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            self.logger.error('parsePage! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))
            print('parsePage! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

    def parseintent(self,response):
        try:
            print('=================::::::::==================== A response from %s just arrived!',response.url)
            self.logger.info('=================::::::::==================== A response from %s just arrived!',response.url)
            metaUserId = response.meta['userId']
            if response.status == 404:
                self.logger.warning('404 USER %s !',response.url)
                user = UserItem()
                user["UserId"] = metaUserId
                user["NeedToCrawl"] = False
                user["UpdatedAt"] = int(time.time())
                user["TwitterState"] = 3
                user["LastSource"] = 1
                yield user
            else:
                UserName = response.xpath("//span[@class='nickname']/text()").extract_first().strip()
                if len(UserName) > 1:
                    request = scrapy.Request("https://twitter.com/"+UserName[1:], self.parse,meta={'handle_httpstatus_list': [301,302,404]})
                    request.meta["userId"] = metaUserId
                    request.meta["username"] = UserName[1:]
                    yield request
        except Exception as e :
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            self.logger.error('parseintent! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))
            print('parseintent! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

    def parse(self, response):
        try:
            print('=================::::::::==================== A response from %s just arrived!',response.url)
            self.logger.info('=================::::::::==================== A response from %s just arrived!',response.url)

            user = UserItem()
            #unavailable user
            userNameunavailable = response.xpath("//div[@class='components-middle']//div[@class='tweet-user-withheld']/span[@class='username u-dir u-textTruncate']/b/text()").extract_first()
            #protected user
            userNameprotected = response.xpath("//div[@class='ProtectedTimeline']//b/text()").extract_first()

            if response.status == 404:
                self.logger.warning('404 USER %s !',response.url)
                metaUserId = response.meta['userId']
                metaUsername = response.meta['username']

                user["UserId"] = metaUserId
                user["UserName"] = metaUsername
                user["NeedToCrawl"] = False
                user["UpdatedAt"] = int(time.time())
                user["TwitterState"] = 3
                user["LastSource"] = 1

                yield user
            elif userNameunavailable != None:
                self.logger.warning('unavailable USER %s !',response.url)
                data =response.xpath("//*[@class='json-data']/@value").extract_first()
                dataObj = json.loads(data)
                user["UserId"] = dataObj.get("profile_id",None)
                user["UserName"] = userNameunavailable
                user["NeedToCrawl"] = False
                user["UpdatedAt"] = int(time.time())
                user["TwitterState"] = 1
                user["LastSource"] = 1
                yield user
            elif userNameprotected:
                self.logger.warning('protected USER %s !',response.url)
                data =response.xpath("//*[@class='json-data']/@value").extract_first()
                dataObj = json.loads(data)
                user["UserId"] = dataObj.get("profile_id",None)
                user["UserName"] = userNameprotected
                user["NeedToCrawl"] = False
                user["UpdatedAt"] = int(time.time())
                user["TwitterState"] = 2
                user["LastSource"] = 1
                yield user
            else:
                user["UserId"] = int(response.xpath("//div[@class='ProfileNav']/@data-user-id").extract_first())
                user["UserName"] = response.xpath("//b[@class='u-linkComplex-target']/text()").extract_first()
                user["ViewName"] = response.xpath("//a[@class='ProfileHeaderCard-nameLink u-textInheritColor js-nav']/text()").extract_first()
                bio = "".join(response.xpath("//*[contains(@class,'ProfileHeaderCard-bio')]//text()").extract())
                if bio != "":
                    user["Bio"] = "".join(bio)

                location = response.xpath("//span[contains(@class,'ProfileHeaderCard-locationText')]/text()").extract_first()
                if location and len(location.strip()) > 2:
                    user["Location"] = location.strip()
                else:
                    user["Location"] = None

                user["Link"] = response.xpath("//span[contains(@class,'ProfileHeaderCard-urlText')]/a/@title").extract_first()
                try:
                    joindate = response.xpath("//span[contains(@class,'ProfileHeaderCard-joinDateText')]/@title").extract_first()
                    if joindate != None:
                        joindateParsed = datetime.strptime(joindate, '%I:%M %p - %d %b %Y') #2:20 PM - 11 Feb 2007
                        user["JoinDate"] = int(joindateParsed.timestamp())
                except Exception as ex:
                    e=2

                birthdate = response.xpath("//span[contains(@class,'ProfileHeaderCard-birthdateText')]/span/text()").extract_first()
                if birthdate and len(birthdate.split("Born on")) > 1 :
                    user["BirthDate"] = birthdate.split("Born on")[1].strip()
                else:
                    user["BirthDate"] = None

                user["TweetsCount"] = response.xpath("//*[contains(@class,'ProfileNav-item ProfileNav-item--tweets')]//*[contains(@class,'ProfileNav-value')]/@data-count").extract_first()
                user["LikesCount"] = response.xpath("//*[contains(@class,'ProfileNav-item ProfileNav-item--favorites')]//*[contains(@class,'ProfileNav-value')]/@data-count").extract_first()
                user["FollowingCount"] = response.xpath("//*[contains(@class,'ProfileNav-item ProfileNav-item--following')]//*[contains(@class,'ProfileNav-value')]/@data-count").extract_first()
                user["FollowersCount"] = response.xpath("//*[contains(@class,'ProfileNav-item ProfileNav-item--followers')]//*[contains(@class,'ProfileNav-value')]/@data-count").extract_first()
                user["NeedToCrawl"] = False
                user["UpdatedAt"] = int(time.time())
                user["TwitterState"] = 0
                user["LastSource"] = 1

                tweets = response.xpath("//*[contains(@class,'tweet js-stream-tweet')]").extract()
                if tweets != None and len(tweets) > 0:
                    chklastTweet = True
                    for tw in tweets:
                        res = self.extractTweetItem(tw)
                        if res != None and len(res) > 0:
                            for r in res:
                                if chklastTweet:
                                    #IsPinnedTweet apper only when it's tweet's of user
                                    # if r.get("IsPinnedTweet") == False:
                                    if r.get("IsPinnedTweet") == False and int(r.get("UserId")) == int(user.get("UserId")):
                                        user["LastTweetId"] = r.get("TweetId")
                                        user["LastTweetDate"] = r.get("CreatedAt")
                                        chklastTweet = False
                                yield r
                yield user
        except Exception as e: 
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            self.logger.error('parse! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))
            print('parse! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))



    def extractTweetItem(self,tweet):
        try:

            #temporarily unavailable
            TombstoneTweet = parser.lxmlParserFirst(tweet,"//div[@class='Tombstone']")
            if TombstoneTweet != None:
                return None

            result = []

            retweetid = parser.lxmlParserFirst(tweet,"@data-retweet-id")

            data_reply_to_users_json = parser.lxmlParserFirst(tweet,"@data-reply-to-users-json")
            data_reply_to_users_json_obj = json.loads(data_reply_to_users_json)

            if retweetid:
                retweetItem = TweetsItem()
                tweetItem = TweetsItem()

                tweetItem["TweetId"] = int(parser.lxmlParserFirst(tweet,"@data-retweet-id"))
                retweetItem["TweetId"] = int(parser.lxmlParserFirst(tweet,"@data-tweet-id"))
                indexUserRetweetFrom = 0
                if len(data_reply_to_users_json_obj) > 1:
                    indexUserRetweetFrom = 1
                tweetItem["UserId"] = int(data_reply_to_users_json_obj[indexUserRetweetFrom].get("id_str",0))
                retweetItem["UserId"] = int(data_reply_to_users_json_obj[0].get("id_str",0))

                tweetItem["UserName"] = data_reply_to_users_json_obj[indexUserRetweetFrom].get("screen_name","")
                retweetItem["UserName"] = data_reply_to_users_json_obj[0].get("screen_name","")


                tweetItem["RetweetedTweetId"] = retweetItem.get("TweetId",0)
                tweetItem["RetweetedUserId"] = retweetItem.get("UserId",0)

                retweetItem["RetweetedTweetId"] = None
                retweetItem["RetweetedUserId"] = None

                retweetItem["Reply"] = int(parser.lxmlParserFirst(tweet,"//span[@class='ProfileTweet-action--reply u-hiddenVisually']/span/@data-tweet-stat-count"))
                retweetItem["Retweet"] = int(parser.lxmlParserFirst(tweet,"//span[@class='ProfileTweet-action--retweet u-hiddenVisually']/span/@data-tweet-stat-count"))
                retweetItem["Favorite"] = int(parser.lxmlParserFirst(tweet,"//span[@class='ProfileTweet-action--favorite u-hiddenVisually']/span/@data-tweet-stat-count"))
                retweetItem["SilverFactor"] = (int(retweetItem.get("Favorite",0)) + int(retweetItem.get("Reply",0))) * int(retweetItem.get("Retweet",0))
                retweetItem["CreatedAt"] = int(parser.lxmlParserFirst(tweet,"//*[contains(@class,'_timestamp')]/@data-time"))
                retweetItem["UpdatedAt"] = int(time.time())
                retweetItem["Lang"] = parser.lxmlParserFirst(tweet,"//p[contains(@class,'TweetTextSize')]/@lang")
                retweetItem["CountryCode"] = None

                tweetItem["Reply"] = 0
                tweetItem["Retweet"] = 0
                tweetItem["Favorite"] = 0
                tweetItem["SilverFactor"] = 0
                tweetItem["CreatedAt"] = retweetItem.get("CreatedAt",0)
                tweetItem["UpdatedAt"] = int(time.time())
                tweetItem["Lang"] = retweetItem.get("Lang",0)
                tweetItem["CountryCode"] = None

                QuotedTweetId = parser.lxmlParserFirst(tweet,"//div[contains(@class,'QuoteTweet-innerContainer')]/@data-item-id")
                if QuotedTweetId:
                    retweetItem["QuotedTweetId"] = int(QuotedTweetId)
                    retweetItem["QuotedUserId"] = int(parser.lxmlParserFirst(tweet,"//div[contains(@class,'QuoteTweet-innerContainer')]/@data-user-id"))
                else:
                    retweetItem["QuotedTweetId"] = None
                    retweetItem["QuotedUserId"] = None

                ReplyUserId = parser.lxmlParserFirst(tweet,"//div[contains(@class,'ReplyingToContextBelowAuthor')]/a/@data-user-id")
                if ReplyUserId :
                    retweetItem["ReplyUserId"] = int(ReplyUserId)
                    retweetItem["ReplyTweetId"] = int(parser.lxmlParserFirst(tweet,"@data-conversation-id"))
                else:
                    retweetItem["ReplyUserId"] = None
                    retweetItem["ReplyTweetId"] = None


                tweetItem["QuotedTweetId"] = None
                tweetItem["QuotedUserId"] = None
                tweetItem["ReplyUserId"] = None
                tweetItem["ReplyTweetId"] = None

                text = " ".join(parser.lxmlParserList(tweet,"//p[contains(@class,'TweetTextSize TweetTextSize--normal js-tweet-text tweet-text')]//text()"))
                retweetItem["Text"] = text
                tweetItem["Text"] = text

                retweetItem["Source"] = 1
                tweetItem["Source"] = 1

                result.append(retweetItem)
                result.append(tweetItem)
            else:
                tweetItem = TweetsItem()

                tweetItem["TweetId"] = int(parser.lxmlParserFirst(tweet,"@data-tweet-id"))

                tweetItem["UserId"] = int(data_reply_to_users_json_obj[0].get("id_str",0))

                tweetItem["UserName"] = data_reply_to_users_json_obj[0].get("screen_name","")

                tweetItem["RetweetedTweetId"] = None
                tweetItem["RetweetedUserId"] = None

                tweetItem["Reply"] = int(parser.lxmlParserFirst(tweet,"//span[@class='ProfileTweet-action--reply u-hiddenVisually']/span/@data-tweet-stat-count"))
                tweetItem["Retweet"] = int(parser.lxmlParserFirst(tweet,"//span[@class='ProfileTweet-action--retweet u-hiddenVisually']/span/@data-tweet-stat-count"))
                tweetItem["Favorite"] = int(parser.lxmlParserFirst(tweet,"//span[@class='ProfileTweet-action--favorite u-hiddenVisually']/span/@data-tweet-stat-count"))
                tweetItem["SilverFactor"] = (int(tweetItem.get("Favorite",0)) + int(tweetItem.get("Reply",0))) * int(tweetItem.get("Retweet",0))
                tweetItem["CreatedAt"] = int(parser.lxmlParserFirst(tweet,"//*[contains(@class,'_timestamp')]/@data-time"))
                tweetItem["UpdatedAt"] = int(time.time())
                tweetItem["Lang"] = parser.lxmlParserFirst(tweet,"//p[contains(@class,'TweetTextSize')]/@lang")
                tweetItem["CountryCode"] = None

                QuotedTweetId = parser.lxmlParserFirst(tweet,"//div[contains(@class,'QuoteTweet-innerContainer')]/@data-item-id")
                if QuotedTweetId:
                    tweetItem["QuotedTweetId"] = int(QuotedTweetId)
                    tweetItem["QuotedUserId"] = int(parser.lxmlParserFirst(tweet,"//div[contains(@class,'QuoteTweet-innerContainer')]/@data-user-id"))
                else:
                    tweetItem["QuotedTweetId"] = None
                    tweetItem["QuotedUserId"] = None

                ReplyUserId = parser.lxmlParserFirst(tweet,"//div[contains(@class,'ReplyingToContextBelowAuthor')]/a/@data-user-id")
                if ReplyUserId :
                    tweetItem["ReplyUserId"] = int(ReplyUserId)
                    tweetItem["ReplyTweetId"] = int(parser.lxmlParserFirst(tweet,"@data-conversation-id"))
                else:
                    tweetItem["ReplyUserId"] = None
                    tweetItem["ReplyTweetId"] = None


                tweetItem["Text"] = " ".join(parser.lxmlParserList(tweet,"//p[contains(@class,'TweetTextSize TweetTextSize--normal js-tweet-text tweet-text')]//text()"))
                tweetItem["Source"] = 1
                PinnedTweet = parser.lxmlParserFirst(tweet,"//div[contains(@class,'user-pinned')]")
                tweetItem["IsPinnedTweet"] = (PinnedTweet != None)

                result.append(tweetItem)

            return result
        except Exception as e:
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            self.logger.error('extractTweetItem! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))
            print('extractTweetItem! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

