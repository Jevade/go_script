package search

import (
	"fmt"

	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feeds *Feed, searchItem string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchItem string, resultch chan<- *Result) {
	searchResult, err := matcher.Search(feed, searchItem)
	if err != nil {
		log.Println("search field:", err)
		return
	}
	for _, result := range searchResult {
		resultch <- result
	}
}

func Display(results <-chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
