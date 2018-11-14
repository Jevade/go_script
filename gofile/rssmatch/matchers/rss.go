package matchers

import (
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"regexp"

	"../search"
)

type rssMatcher struct{}
type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}
	image struct {
		XMLName xml.Name `xml:"item"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	channel struct {
		XMLName        xml.Name `xml:"item"`
		PubDate        string   `xml:"pubDate"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		LastBuildDate  string   `xml:"lastbuilddate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (matcher rssMatcher) Search(feed *search.Feed, itemString string) ([]*search.Result, error) {
	var rssResults []*search.Result
	document, err := matcher.retrieve(feed)
	if err != nil {
		return rssResults, err
	}
	for _, channelItem := range document.Channel.Item {
		matched, err := regexp.MatchString(itemString, channelItem.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			rssResults = append(rssResults, &search.Result{
				Field:   "title",
				Content: channelItem.Title,
			})
		}
		matched, err = regexp.MatchString(itemString, channelItem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			rssResults = append(rssResults, &search.Result{
				Field:   "description",
				Content: channelItem.Description,
			})
		}
	}
	return rssResults, nil
}

func (matcher rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		log.Println("No rss feed")
		return nil, errors.New("No rss feed")
	}
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("Http error.Code: " + string(resp.StatusCode))
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, nil
}
