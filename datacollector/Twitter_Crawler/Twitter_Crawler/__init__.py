
# import scrapy
# from scrapy.crawler import CrawlerProcess
# from scrapy.utils.project import get_project_settings
#
# from Twitter_Crawler.spiders.userprofileSpider import userprofileSpider
#
# settings = get_project_settings()
# process = CrawlerProcess(settings)
#
# ##Spiders
# process.crawl(userprofileSpider)
#
# process.start(stop_after_crawl=False)
# process.start()# the script will block here until all crawling jobs are finished

#####################
#####################

#https://stackoverflow.com/questions/47552507/how-to-schedule-scrapy-crawl-execution-programmatically

# from twisted.internet import reactor
# from scrapy.crawler import CrawlerRunner
# from scrapy.utils.project import get_project_settings
#
# from Twitter_Crawler.spiders.userprofileSpider import userprofileSpider
#
#
# def crawl_job():
#     """
#     Job to start spiders.
#     Return Deferred, which will execute after crawl has completed.
#     """
#     settings = get_project_settings()
#     runner = CrawlerRunner(settings)
#     return runner.crawl(userprofileSpider)
#
# def schedule_next_crawl(null, sleep_time):
#     """
#     Schedule the next crawl
#     """
#     reactor.callLater(sleep_time, crawl)
#
# def crawl():
#     """
#     A "recursive" function that schedules a crawl 1 seconds after
#     each successful crawl.
#     """
#     # crawl_job() returns a Deferred
#     d = crawl_job()
#     # call schedule_next_crawl(<scrapy response>, n) after crawl job is complete
#     d.addCallback(schedule_next_crawl, 1)
#     d.addErrback(catch_error)
#
# def catch_error(failure):
#     print(failure.value)

# if __name__=="__main__":
#     crawl()
#     reactor.run()

#######################
#######################

 