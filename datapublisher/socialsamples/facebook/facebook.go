package main

import (
	"fmt"

	"github.com/huandu/facebook"
	"github.com/tarekbadrshalaan/goStuff/jsonparser"
)

func main() {
	//Testpublishpage
	// accessToken := ""
	// id := "324696864967681"
	//free
	//accessToken := ""
	//id := "181841378681657"

	//elwizarapage-elwizarapub
	accessToken := ""
	id := "340599076510354"
	//  res, _ := facebook.Get("me", facebook.Params{
	// 	"fields":       "name",
	// 	"access_token": accessToken,
	// })
	// fmt.Println("Here is my Facebook first name:", res["name"])

	// targeting := `{
	// 	"geo_locations": {
	// 		"countries": [
	// 		  "CA"
	// 		]
	// 	  }
	// }`
	// res, err := facebook.Post(id+"/feed", facebook.Params{
	// 	"message":      "Hello Fans with egypt!",
	// 	"access_token": accessToken,
	// 	"targeting":    targeting,
	// })
	// if err != nil {
	// 	fmt.Printf("error : %v\n", err)
	// }
	// fmt.Printf("result : %v\n", res)

	res, err := facebook.Batch(facebook.Params{
		"access_token": accessToken,
		"file1":        facebook.File("1.png"),
	}, facebook.Params{
		"method":         "POST",
		"relative_url":   id + "/photos",
		"body":           "message=you can find more here :- \nelwz.me/?enRcBA41BodCC6",
		"attached_files": "file1",
	})

	if err != nil {
		fmt.Printf("error : %v\n", err)
	} else {
		id, err := jsonparser.Getkeystring(res[0], "body", "post_id")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("facebookid", id)
		}
	}

}
