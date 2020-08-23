# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html

import sys,os
import json

import psycopg2
from Twitter_Crawler.items import TweetsItem,UserItem


class TwitterCrawlerPipeline(object):
    def __init__(self):
        try: 
            with open('Twitter_Crawler/conf.json') as f:
                configuration  = f.read()
            configuration_obj = json.loads(configuration)
            DBUSER = configuration_obj.get("DBUSER")
            DBPASSWORD = configuration_obj.get("DBPASSWORD")
            DBNAME = configuration_obj.get("DBNAME")
            DBHost = configuration_obj.get("DBHost")
            DBPort = configuration_obj.get("DBPort")

            self.conn = psycopg2.connect(database = DBNAME, user = DBUSER, password =DBPASSWORD, host = DBHost, port = DBPort)
            self.cur = self.conn.cursor()
            print("Opened database successfully")
        except Exception as e:
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            print('__init__! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

    def process_item(self, item, spider):
        try :
            obj = item._values
            if type(item) is TweetsItem:
                query =  """SELECT public."InsertTweetsTB"(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s);"""
                data = (obj.get("TweetId"),
                        obj.get("UserId"),
                        obj.get("Favorite"),
                        obj.get("Reply"),
                        obj.get("Retweet"),
                        obj.get("SilverFactor"),
                        obj.get("UserName"),
                        obj.get("CreatedAt"),
                        obj.get("UpdatedAt"),
                        obj.get("Lang"),
                        obj.get("CountryCode"),
                        obj.get("RetweetedTweetId"),
                        obj.get("RetweetedUserId"),
                        obj.get("QuotedTweetId"),
                        obj.get("QuotedUserId"),
                        obj.get("ReplyTweetId"),
                        obj.get("ReplyUserId"),
                        obj.get("Source"),
                        obj.get("Text"),
                        obj.get("IsPinnedTweet"))
                self.cur.execute(query,data)
                self.conn.commit()

            elif type(item) is UserItem:
                query = """SELECT public."InsertUsersProfileTB"(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s);"""
                data = (
                        obj.get("UserId"),
                        obj.get("UserName"),
                        obj.get("ViewName"),
                        obj.get("Bio"),
                        obj.get("Location"),
                        obj.get("Link"),
                        obj.get("JoinDate"),
                        obj.get("TweetsCount"),
                        obj.get("LikesCount"),
                        obj.get("FollowingCount"),
                        obj.get("FollowersCount"),
                        obj.get("BirthDate"),
                        obj.get("NeedToCrawl"),
                        obj.get("UpdatedAt"),
                        obj.get("LastTweetDate"),
                        obj.get("LastTweetId"),
                        obj.get("TwitterState"),
                        obj.get("LastSource")
                    )

                self.cur.execute(query,data)
                self.conn.commit()

        except Exception as e:
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            print('Error : process_item! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

        return item
