package itemspy

import (
	"fmt"
	"net/http"

	"../model"
	"github.com/PuerkitoBio/goquery"
	"github.com/lexkong/log"
)

//GetCityUrls is to get cities from url
func GetCityUrls(url string, itemch chan<- interface{}) {

	log.Info("Searching city")
	res, _ := http.Get(url)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		str := fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)
		log.Fatal(str, nil)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Document read error:", err)
	}

	// Find the review items
	doc.Find("div.content-province").Each(func(j int, proSelection *goquery.Selection) {
		province := proSelection.Find("div.content-province-title").Text()
		proSelection.Find("div.content-cities").Find("a").Each(func(k int, aSelection *goquery.Selection) {

			cityInfo := &model.CityInfo{Province: province}

			cityInfo.Cityname = RemoveSpace(aSelection.Text())
			cityInfo.BaseHost, _ = aSelection.Attr("href")
			cityInfo.Shortcut, _ = aSelection.Attr("href")
			fmt.Println(cityInfo)
		})
	})
}
