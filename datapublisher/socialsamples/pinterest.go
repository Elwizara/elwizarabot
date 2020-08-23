package main

/*
import (
	"fmt"
	"os"

	pinterest "github.com/carrot/go-pinterest"
	"github.com/carrot/go-pinterest/controllers"
)

func pinterestPublisher() {

	client := pinterest.NewClient().RegisterAccessToken("TOKEN")

	file, err := os.Open("2.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	optionals := &controllers.PinCreateOptionals{Link: "https://elwizara.com/?lang=en&year=2018&month=9&id=1038291935776190464", Image: file}
	res, err := client.Pins.Create("elwizaracom/english", "hellofromlocal", optionals)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Id)
}

func pinterestBordes() {

	arr := [][]string{
		{"pt", "Portuguese"},     //595319713182246716
		{"es", "Spanish"},        //595319713182246717
		{"und", "UnDefined"},     //595319713182246718
		{"ar", "Arabic"},         //595319713182246719
		{"ko", "Korean"},         //595319713182246720
		{"th", "Thai"},           //595319713182246721
		{"tr", "Turkish"},        //595319713182246722
		{"tl", "Tagalog"},        //595319713182246723
		{"in", "Indonesian"},     //595319713182246724
		{"fr", "French"},         //595319713182246725
		{"it", "Italian"},        //595319713182246726
		{"ru", "Russian"},        //595319713182246727
		{"hi", "Hindi"},          //595319713182246728
		{"ur", "Urdu"},           //595319713182246729
		{"de", "German"},         //595319713182246730
		{"nl", "Dutch"},          //595319713182246731
		{"fa", "Persian"},        //595319713182246732
		{"ca", "Catalan"},        //595319713182246733
		{"pl", "Polish"},         //595319713182246734
		{"ht", "Haitian Creole"}, //595319713182246735
		{"sv", "Swedish"},        //595319713182246736
		{"et", "Estonian"},       //595319713182246737
		{"ta", "Tamil"},          //595319713182246738
		{"el", "Greek"},          //595319713182246739
		{"zh", "Chinese"},        //595319713182246740
		{"fi", "Finnish"},        //595319713182246741
		{"vi", "Vietnamese"},     //595319713182246742
		{"no", "Norwegian"},      //595319713182246743
		{"uk", "Ukrainian"},      //595319713182246744
		{"eu", "Basque"},         //595319713182246745
		{"da", "Danish"},         //595319713182246746
		{"ne", "Nepali"},         //595319713182246747
		{"cs", "Czech"},          //595319713182246748
		{"lv", "Latvian"},        //595319713182246749
		{"cy", "Welsh"},          //595319713182246750
		{"ro", "Romanian"},       //595319713182246751
		{"ps", "Pashto"},         //595319713182246752
		{"lt", "Lithuanian"},     //595319713182246753
		{"bn", "Bengali"},        //595319713182246760
		{"hu", "Hungarian"},      //595319713182246762
		{"sl", "Slovenian"},      //595319713182246763
		{"mr", "Marathi"},        //595319713182246769
		{"is", "Icelandic"},      //595319713182246770
		{"sr", "Serbian"},
		{"te", "Telugu"},
		{"my", "Myanmar"},
		{"iw", "Hebrew"},
		{"ml", "Malayalam"},
		{"bg", "Bulgarian"},
		{"kn", "Kannada"},
		{"gu", "Gujarati"},
		{"pa", "Panjabi"}}

	client := pinterest.NewClient().RegisterAccessToken("TOKEN")
	for _, k := range arr {
		langName := k[1]
		lang := k[0]
		Description := fmt.Sprintf(".｡*ﾟ+.*.｡(❁´◡`❁)｡.｡:+* Best Tweets In %v : you can find more here  : \"https://elwizara.com/?lang=%v\"", langName, lang)

		optionals := &controllers.BoardCreateOptionals{Description: Description}

		res, err := client.Boards.Create(langName, optionals)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.Id)
	}
}

func main() {
	pinterestPublisher()
	//pinterestBordes()
}
*/
