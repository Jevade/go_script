package itemspy

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../model"
	"github.com/PuerkitoBio/goquery"
)

//GetItemInfo is to get iteminfo from website
func GetItemInfo(url string, itemch chan<- interface{}) {

	fmt.Println("Searching items")
	res, _ := http.Get(url)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("div#infolist").Find("tr.ac_item").Each(func(j int, divSelection *goquery.Selection) {
		fmt.Println(j)
		itemInfo := &model.ItemInfo{}
		itemInfo.Desc = RemoveSpace(divSelection.Find("div.item-desc").Text())
		itemInfo.ItemTitle = RemoveSpace(divSelection.Find("td.t").Find("a").Eq(0).Text())
		itemInfo.URL, _ = divSelection.Find("a").Eq(0).Attr("href")
		itemInfo.District = RemoveSpace(divSelection.Find("p.seller").Find("a.c_666").Eq(0).Text())
		itemInfo.Area = RemoveSpace(divSelection.Find("p.seller").Find("a.c_666").Eq(1).Text())
		itemInfo.CreatedAt = RemoveSpace(strings.Split(divSelection.Find("p.seller").Text(), "/")[1])
		pricestr := divSelection.Find("td.vertop-es").Find("b.pri").Text()
		price, _ := strconv.ParseFloat(pricestr, 32)
		logr, _ := divSelection.Attr("logr")
		logrs := strings.Split(logr, "_")
		itemInfo.UserID = logrs[2]
		itemInfo.ItemID = logrs[3]
		itemInfo.Price = float32(price)
		fmt.Println(itemInfo.ItemTitle, price)
		itemch <- itemInfo
	})
}
