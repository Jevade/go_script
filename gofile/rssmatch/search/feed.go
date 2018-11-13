package search

import (
	"encoding/json"
	"log"
	"os"
)

const dataFile = "data/data.json"

//Feed to define a rss source
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//RetrieveFeeds will unmarshll data to feeds
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		log.Println("Open file error", err)
		return nil, err
	}
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		log.Println("Decode file error", err)
		return nil, err
	}
	return feeds, nil
}
