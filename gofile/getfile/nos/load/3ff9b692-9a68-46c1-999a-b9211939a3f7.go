package itemspy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	fmt.Println(res.Body)
	log.Info("Document read success")
	var city map[string](map[string]string)
	// Find the review items
	doc.Find("script").Each(func(j int, proSelection *goquery.Selection) {
		province := proSelection.Text()
		strList := strings.Split(province, "=")
		if len(strList) < 4 {
			return
		}
		fmt.Println("len of script is:", len(strList))
		if err := json.Unmarshal([]byte(strList[3]), &city); err == nil {
			fmt.Println(city)
		} else {
			fmt.Println(err)
			return
		}
		copeCityJSON(city, itemch)
	})
	log.Info("Document read failed")
}

func copeCityJSON(cities map[string](map[string]string), itemch chan<- interface{}) {
	for province := range cities {
		for city := range cities[province] {
			cityInfo := &model.CityInfo{Province: province, Cityname: city}
			cityInfo.Shortcut = strings.Split(cities[province][city], "|")[0]
			cityInfo.BaseHost = "www." + cityInfo.Shortcut + ".58.com"
			fmt.Println(cityInfo)
			itemch <- cityInfo
			log.Info(province)
		}
	}

}
