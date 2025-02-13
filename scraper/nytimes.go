package scraper

import (
	"os"
	"regexp"
	"time"

	"github.com/ilmedova/streamly/utils"
	"github.com/valyala/fasthttp"
)

func init() {
	scraperName := os.Getenv("SCRAPER_NAME")
	if scraperName == "" || scraperName == "nytimes" {
		scraperManager.list = append(scraperManager.list, nytimesScraper)
	}
}

func nytimesScraper() []NewsItem {
	newsItems := make([]NewsItem, 0)
	url := "https://rss.nytimes.com/services/xml/rss/nyt/World.xml"
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil || status != fasthttp.StatusOK {
		return newsItems
	}

	r, _ := regexp.Compile("[\n\r]")
	text := string(r.ReplaceAll(resp, []byte{}))

	reg, _ := regexp.Compile(`<item>\s*<title>(.*?)</title>\s*<link>(.*?)</link>.*?<pubDate>(.*?)</pubDate>`)
	res := reg.FindAllStringSubmatch(text, -1)

	for _, matchedItem := range res {
		t, _ := time.Parse(time.RFC1123Z, matchedItem[3])
		newsItems = append(newsItems, NewsItem{
			Filter: IsNeedFilter(matchedItem[1], []string{}),
			Title:  utils.FormatTitle(matchedItem[1]),
			Link:   matchedItem[2],
			Origin: "The New York Times",
			Time:   t.Unix(),
		})
	}

	return newsItems
}
