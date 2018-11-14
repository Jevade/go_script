package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feeds *Feed, searchItem string) ([]*Result, error) {
	return nil, nil
}
