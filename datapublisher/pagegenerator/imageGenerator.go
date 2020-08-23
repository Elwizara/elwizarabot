package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func generateTweetImage(db *sql.DB, tweet *PublishedTweetsDTO) error {
	<-time.After(time.Second * 10)
	data, err := createSplashRequest(tweet)
	if len(data) < 1000 || err != nil {
		logs.Warningf("CreateImageResult:%v:%d:%s|Error:%v|Tryagain", tweet.Lang, tweet.TweetID, data, err)
		<-time.After(time.Second * 10)
		data, err = createSplashRequest(tweet)
		if len(data) < 1000 || err != nil {
			logs.Warningf("Second Error ||CreateImageResult:%v:%d:%s|Error:%v|Tryagain", tweet.Lang, tweet.TweetID, data, err)
			<-time.After(time.Second * 10)
			data, err = createSplashRequest(tweet)
			if len(data) < 1000 || err != nil {
				finalErr := fmt.Errorf("Third Error ||CreateImageResult:%v:%d:%s|Error:%v|WillNotTryAgain", tweet.Lang, tweet.TweetID, data, err)
				logs.Warning(finalErr)
				return finalErr
			}
		}
	}
	err = ioutil.WriteFile(tweet.ImagePath, data, 0666)
	if err != nil {
		logs.Critical(err)
		return err
	}
	logs.Infof("Generate Image %v", tweet.BasePath)
	return nil
}

func createSplashRequest(tweet *PublishedTweetsDTO) ([]byte, error) {
	luaSource := `-- Arguments:
		-- * url - URL to render;
		-- * css - CSS selector to render;

		-- main script
		function main(splash)
			-- splash:set_viewport_size(2560, 1440)
			-- splash:set_viewport_size(1920, 1080)
			-- splash:set_viewport_size(1366, 768)

			-- set screen size
			splash:set_viewport_size(1920, 1080)
			assert(splash:wait(1))

			-- this function returns element bounding box
			local get_bbox = splash:jsfunc([[
				function(divId) {
					var el = document.getElementById(divId);
					var r = el.getBoundingClientRect();
					return [r.left, r.top, r.right, r.bottom-10];
				}
			]])

			assert(splash:go(splash.args.url))
			assert(splash:wait(60))

			local region = get_bbox("container")
			return splash:png{region=region,render_all=true}
		end	`

	url := url.URL{Scheme: "http", Host: config.ScrapySplash, Path: "execute"}
	q := url.Query()
	q.Add("timeout", "90.0")
	dumpPageURL := fmt.Sprintf("%v/%v", config.DumpPagesURL, tweet.DumpRelativePath)
	q.Add("url", dumpPageURL)
	q.Add("lua_source", luaSource)
	url.RawQuery = q.Encode()

	response, err := http.Get(url.String())

	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		logs.Critical(err)
		return nil, err
	}
	return data, err
}

//docker run --rm -p 8050:8050 --net host scrapinghub/splash
//docker run --rm -p 9090:80 -v "$PWD":/usr/local/apache2/htdocs/ httpd
