package main

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"io/ioutil"
// 	"net/url"
// 	"os"
// 	"time"

// 	"github.com/tumblr/tumblrclient"
// )

// func tumblrPublisher() {

// 	ConsumerKey := "ConsumerKey"
// 	ConsumerSecret := "ConsumerSecret"
// 	AccessToken := "AccessToken"
// 	AccessTokenSecret := "AccessTokenSecret"

// 	client := tumblrclient.NewClientWithToken(ConsumerKey, ConsumerSecret, AccessToken, AccessTokenSecret)
// 	u := url.Values{"type": []string{"text"}, "state": []string{"published"}, "title": []string{"HelloWorld"}, "body": []string{"FirstPost"}}
// 	_, err := client.PostWithParams("blog/elwizaracom/post", u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("tumblr Published")

// }

// func tumblrphoto() {

// 	ConsumerKey := "ConsumerKey"
// 	ConsumerSecret := "ConsumerSecret"
// 	AccessToken := "AccessToken"
// 	AccessTokenSecret := "AccessTokenSecret"

// 	client := tumblrclient.NewClientWithToken(ConsumerKey, ConsumerSecret, AccessToken, AccessTokenSecret)

// 	data, err := ioutil.ReadFile("2.png")
// 	if os.IsNotExist(err) {
// 		fmt.Println("file not exist ")
// 		return
// 	}
// 	encoded := base64.StdEncoding.EncodeToString(data)
// 	fmt.Println("encoded done")
// 	u := url.Values{"type": []string{"photo"}, "state": []string{"published"}, "caption": []string{"HelloImage"},
// 		"link":   []string{"https://elwizara.com/?lang=ar&year=2018&month=9&id=1037404477677072386"},
// 		"source": []string{"https://elwizara.com/Data/ar/2018/9/1037404477677072386.png"}, "data64": []string{encoded}}

// 	res, err := client.PostWithParams("blog/elwizaracom/post", u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	res.PopulateFromBody()
// 	fmt.Printf("%s\n", res.GetBody())
// 	//var id int64 = int64(res.Result["id"])
// 	if id, ok := res.Result["id"].(float64); ok {
// 		ss := int64(id)
// 		fmt.Printf("%d\n", ss)
// 	} else {
// 		fmt.Println("not string")
// 	}
// 	fmt.Println("tumblr Published")

// }

// //{"meta":{"status":201,"msg":"Created"},"response":{"id":177901254686,"state":"published","display_text":"Posted to elwizaracom"}}
// func main() {
// 	//tumblrPublisher()
// 	//tumblrphoto()
// 	timeNow := time.Now()
// 	st := "https://elwizara.com/?lang=%v&year=%d&month=%d&id=%v"
// 	UserPrimaryLanguage := "en"

// 	TweetID := 1038291935776190464

// 	s := fmt.Sprintf(st, UserPrimaryLanguage, timeNow.Year(), timeNow.Month(), TweetID)
// 	PublishStatus := fmt.Sprintf("you can find more awesome tweets here:-\n%v", s)
// 	fmt.Println(PublishStatus)
// }
