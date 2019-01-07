package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	const base_url string = "https://www.qiushibaike.com/hot/page/"
	mi := make(map[string]HotsContent)
	for i := 1; i < 10; i++ {
		fmt.Print(i)
		fmt.Println("开始爬取糗事百科热点笑话...,第" + string(i) + "页")
		qiubai_url := base_url + string(i) + "/"
		fmt.Println(qiubai_url)
		qiubaispy(qiubai_url, i, mi)
	}
	index := 1
	for _, v := range mi {
		index += 1
		fmt.Println(v.content)
		fmt.Println(v.comment)
		fmt.Println(index)
	}
}

type HotsContent struct {
	num     int
	content string
	comment string
	url     string
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func qiubaispy(url string, index int, m1 map[string]HotsContent) {

	var filename = "糗百.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if CheckFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	Check(err1)
	defer f.Close()
	js, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	js.Find("#content-left .article").Each(func(i int, contentSelection *goquery.Selection) {
		//先判断是否有图片
		img, _ := contentSelection.Find(".thumb a").Attr("href")
		if len(img) == 0 {
			hotsArt := HotsContent{}
			content := contentSelection.Find(".content span").Text()
			url, _ := contentSelection.Find(".contentHerf").Attr("href")
			comment_name := contentSelection.Find(".cmtMain .cmt-name").Text()
			comment_cont := contentSelection.Find(".cmtMain .main-text").Text()
			hotsArt.num = i + 1
			hotsArt.url = "https://www.qiushibaike.com" + url
			hotsArt.content = strings.Replace(content, "\n", "", -1)
			hotsArt.comment = strings.Replace(comment_name+comment_cont, "\n", "", -1)
			_, err1 := io.WriteString(f, "======================================================") //写入文件(字符串)
			Check(err1)
			_, err1 = io.WriteString(f, "\n\t"+hotsArt.content)
			Check(err1)
			_, err1 = io.WriteString(f, "\n \t最热评论:"+string(hotsArt.comment))
			Check(err1)
			_, err1 = io.WriteString(f, "\n 地址"+hotsArt.url+"\n")
			Check(err1)
			m1[strconv.Itoa(hotsArt.num+(index-1)*25)] = hotsArt
		}
	})

}
