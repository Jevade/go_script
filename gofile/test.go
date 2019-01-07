package main
import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "log"
    "github.com/PuerkitoBio/goquery"
)
func ExampleScrape() {
  // Request the HTML page.
  fmt.Println(1222222)
  res, err := http.Get("http://news.baidu.com")
  if err != nil {
    log.Fatal(err)
    fmt.Println(err.Error())
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    fmt.Println(res.StatusCode)
  }
  fmt.Println(res.StatusCode)
  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  // Find the review items
  doc.Find("#channel-all > div > ul > li").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    band := s.Find("a").Text()
    href := s.Find("a")
    fmt.Printf("Review %d: %s_%s \n", i, band,href)
  })
  
}
func themain() {
    imagPath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
    //图片正则
    name:=getPath(imagPath)
 
    //通过http请求获取图片的流文件
    resp, _ := http.Get(name)
    body, _ := ioutil.ReadAll(resp.Body)
    out, _ := os.Create("1.jpg")
    io.Copy(out, bytes.NewReader(body))
    return
}
func getPath(Path string) (string){
    reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
    name := reg.FindStringSubmatch(Path)[0]
    return name
}

func main(){
    ExampleScrape()
}
