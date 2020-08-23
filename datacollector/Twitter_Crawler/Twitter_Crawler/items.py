# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://doc.scrapy.org/en/latest/topics/items.html

from scrapy import Item,Field


class TweetsItem(Item):
    TweetId  = Field()
    UserId  = Field()
    Favorite  = Field()
    Reply  = Field()
    Retweet  = Field()
    SilverFactor  = Field()
    UserName  = Field()
    CreatedAt  = Field()
    UpdatedAt  = Field()
    Lang  = Field()
    CountryCode  = Field()
    RetweetedTweetId  = Field()
    RetweetedUserId  = Field()
    QuotedTweetId  = Field()
    QuotedUserId  = Field()
    ReplyTweetId  = Field()
    ReplyUserId  = Field()
    Source = Field()
    Text = Field()
    IsPinnedTweet = Field()

class UserItem(Item):
    UserId = Field()
    UserName = Field()
    ViewName = Field()
    Bio = Field()
    Location = Field()
    Link = Field()
    JoinDate = Field()
    BirthDate = Field()
    TweetsCount = Field()
    LikesCount = Field()
    FollowingCount = Field()
    FollowersCount = Field()
    NeedToCrawl = Field()
    UpdatedAt = Field()
    LastTweetDate = Field()
    LastTweetId = Field()
    TwitterState = Field()
    LastSource = Field()
