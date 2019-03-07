package itemspy

import (
	"fmt"
	"log"

	"../model"
	"github.com/PuerkitoBio/goquery"
)

//GetChannelUrls is to get channel info
func GetChannelUrls(turl string, typech chan<- interface{}) {
	url := "https://" + turl + ".58.com/sale.shtml"
	baseHost := "https://" + turl + ".58.com"
	js, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	js.Find("div#ymenu-side").Find("span.dlb").Find("a").Each(func(j int, aSelection *goquery.Selection) {
		itemURL, _ := aSelection.Attr("href")
		if "" == itemURL {
			return
		}
		typeInfo := &model.TypeInfo{}
		typeInfo.Typename = RemoveSpace(aSelection.Text())
		typeInfo.URL = baseHost + itemURL
		fmt.Println(typeInfo.URL)
		typech <- typeInfo
	})
	close(typech)
}
