package search

import (
	"fmt"
	"sync"

	"log"
)

var matchers = make(map[string]Matcher)

func Run(searchItem string) {
	fmt.Println(searchItem)
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Println("Retrieved failed")
	}
	results := make(chan *Result)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))
	for _, feed := range feeds {
		matcher, exist := matchers[feed.Type]
		if !exist {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchItem, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	go func() {
		waitGroup.Wait()
		close(results)
	}()
	Display(results)
	log.Println("done")
}

func Register(feedType string, matcher Matcher) {
	if _, exist := matchers[feedType]; exist {
		log.Println("Matcher have existd")
		return
	}
	// log.Info("register" + feedType + "matcher")
	matchers[feedType] = matcher
}
