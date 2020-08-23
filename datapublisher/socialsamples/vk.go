package main

/*
import (
	"fmt"

	vk "github.com/kovetskiy/go-vkcom"
)

func getAccessToken() {
	 auth := vk.Auth{
		AppId:       "6688348",
		AppSecret:   "AppSecret",
		Permissions: []string{"wall"},
		RedirectUri: "http://ewlizara.com/",
		Display:     "page", // more on https://vk.com/dev/auth_sites
	}
	auth := vk.Auth{
		AppId:       "6688760",
		AppSecret:   "AppSecret",
		Permissions: []string{"wall"},
		RedirectUri: "http://ewlizara.com/",
		Display:     "page", // more on https://vk.com/dev/auth_sites
	}

	authU, err := auth.GetAuthUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println(authU)
	//https://oauth.vk.com/authorize?client_id=6688348&display=page&redirect_uri=http%3A%2F%2Fewlizara.com%2F&response_type=code&scope=wall
	//https://oauth.vk.com/authorize?client_id=6688348&display=page&redirect_uri=http%3A%2F%2Fewlizara.com%2F&scope=friends&response_type=code&v=5.84
	//https://oauth.vk.com/authorize?client_id=6688348&display=page&redirect_uri=http%3A%2F%2Fewlizara.com%2F&scope=wall,offline,friends,nohttps&response_type=code&v=5.84
	//https://oauth.vk.com/authorize?client_id=6688348&display=page&redirect_uri=http%3A%2F%2Fewlizara.com%2F&scope=notify,friends,photos,audio,video,stories,pages,notes,wall,ads,offline,docs,groups,notifications,stats,email,market,nohttps&response_type=code&v=5.84
	//https://oauth.vk.com/authorize?client_id=6688760&display=page&redirect_uri=http%3A%2F%2Fewlizara.com%2F&scope=notify,friends,photos,audio,video,stories,pages,notes,wall,ads,offline,docs,groups,notifications,stats,email,market,nohttps&response_type=code&v=5.84

	//http://ewlizara.com/?code=7d18c0630d22519555

	token, err := auth.GetAccessToken("ced2f694eadd693bdd")
	if err != nil {
		panic(err)
	}
	api := vk.Api{AccessToken: token}
	fmt.Println(api.AccessToken)
}

func useAccessToken() {
	//{"access_token":"access_token","expires_in":0,"user_id":506468036}
	//{"access_token":"access_token","expires_in":0,"user_id":506468036,"secret":"398bcc442385d51b87"}
	//{"access_token":"access_token","expires_in":0,"user_id":506468036,"email":"elwizara.com@gmail.com","secret":"58c1ea0fb63355bef3"}
	//{"access_token":"access_token","expires_in":0,"user_id":506468036,"email":"elwizara.com@gmail.com","secret":"c5afdc40063220a555"}
	api := vk.Api{AccessToken: vk.AccessToken{
		Token:     "AccessToken",
		ExpiresIn: 0,
		UserId:    506468036,
	}}

	query := map[string]string{
		"owner_id":     "506468036",
		"friends_only": "0",
		"from_group":   "0",
		"signed":       "0",
		"message":      "Newtest",
		"sig":          "sig",
	}
	response, err := api.Request("wall.post", query)

	// query = map[string]string{
	// 	"sig": "sig",
	// }
	// response, err := api.Request("friends.get", query)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response : %v\n", response)

}

func main() {
	//getAccessToken()
	useAccessToken()
}*/
