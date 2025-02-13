package scraper

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ilmedova/streamly/utils"
)

var filterList [][][]string = [][][]string{

	// If you want to ignore news on certain topics then just add them below
	// Example
	//{{"inflation"}},
	//{{"war"}},
}

type NewsItem struct {
	Title    string
	Link     string
	Origin   string
	Filter   []string
	Time     int64
	Distance uint64
	Keywords []string
}

type ScraperManager struct {
	list []Scraper
}

var scraperManager = &ScraperManager{
	list: make([]Scraper, 0),
}

type Scraper = func() []NewsItem

func Get() []NewsItem {
	scraperCount := len(scraperManager.list)
	fmt.Println("Scrapers found:", scraperCount)
	var wg sync.WaitGroup
	wg.Add(scraperCount)
	channel := make(chan *[]NewsItem, scraperCount)
	for i := 0; i < scraperCount; i++ {
		go func(ch chan *[]NewsItem, index int, wg *sync.WaitGroup) {
			list := scraperManager.list[index]()
			ch <- &list
			wg.Done()
		}(channel, i, &wg)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	newsItems := make([]NewsItem, 0)
	for ch := range channel {
		newsItems = append(newsItems, *ch...)
	}
	scraperName := os.Getenv("SCRAPER_NAME")
	if scraperName != "" {
		fmt.Println("newsItems", newsItems)
	}
	return newsItems
}

var noFilterCache = os.Getenv("FILTER_DISABLE") == "true"

func IsNeedFilter(title string, moreFilterWords []string) []string {
	if noFilterCache {
		return []string{}
	}
	for _, filterWord := range moreFilterWords {
		if strings.Contains(title, filterWord) {
			return []string{filterWord}
		}
	}
	for _, filterWordGroup := range filterList {
		isNeedFilter := true
		wordMatched := make([]string, 0)
		preWorkMatchIndex := -1
		for _, wordlist := range filterWordGroup {
			isWordMatch := false
			matchTitle := title
			if preWorkMatchIndex >= 0 {
				matchTitle = title[preWorkMatchIndex:]
			}
			for _, word := range wordlist {
				index := 0
				matched := false
				if word[0] == '-' {
					matched = strings.HasPrefix(matchTitle, word[1:])
				} else if word == "#d" {
					index = preWorkMatchIndex
					if index < len(title) && utils.ByteIsNumericOrSpace(title[index]) {
						for index < len(title) && utils.ByteIsNumericOrSpace(title[index]) {
							index += 1
						}
						word = ""
						matched = true
					} else {
						matched = false
					}
				} else {
					index = strings.Index(matchTitle, word)
					matched = index >= 0
				}

				if matched {
					isWordMatch = true
					wordMatched = append(wordMatched, word)
					preWorkMatchIndex = index + len(word)
					break
				}
			}
			if !isWordMatch {
				isNeedFilter = false
				break
			}
		}
		if isNeedFilter {
			return wordMatched
		}
	}
	return []string{}
}

func ItemToHtml(item *NewsItem) string {
	pubTime := ""
	if item.Time > 0 {
		pubTime = " - " + utils.FormatTime(item.Time, "01/02 15:04:05")
	}
	return "<a target=\"_blank\" href=\"" + item.Link + "\">" + item.Title + "</a> [" + item.Origin + pubTime + "]"
}
