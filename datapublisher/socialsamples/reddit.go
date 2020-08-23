package main

//https://www.reddit.com/api/v1/authorize?client_id=ID&response_type=code&state=147&redirect_uri=https://www.elwizara.com&duration=permanent&scope=identity,edit,flair,history,modconfig,modflair,modlog,modposts,modwiki,mysubreddits,privatemessages,read,report,save,submit,subscribe,vote,wikiedit,wikiread
/*
func main() {

	config := reddit.BotConfig{
		Agent: "Golang:automatic:v2.0 (by /u/elwizaracom)",
		App: reddit.App{
			ID:       "BOakDQRSuHU6cw",
			Secret:   "TOKEN",
			Username: "elwizaracom",
			Password: "PASS",
		},
		Rate: 5 * time.Second,
	}

	bot, err := reddit.NewBot(config)

	//bot, err := reddit.NewBotFromAgentFile("agentfile", 5*time.Second)
	res, err := bot.PostLink("u_elwizaracom", "Awesome Tweet", "https://elwizara.com/Data/en/2018/9/1038284714170376192.png")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		fmt.Printf("error : %v", err)
	}

	//{"json": {"errors": [], "data": {"url": "https://www.reddit.com/r/u_elwizaracom/comments/9ewfg2/awesome_tweet/", "drafts_count": 0, "id": "9ewfg2", "name": "t3_9ewfg2"}}}

	result, err := jsonparser.JSONParser(buf.Bytes(), "json", "data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	postID := result["name"]

	link := "find more here :  https://elwizara.com/?lang=en&year=2018&month=9&id=1038284714170376192"
	bot.Reply(postID.(string), link)

}

*/
