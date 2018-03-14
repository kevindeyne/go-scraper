package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Game struct {
	Title       string
	Description string
	Url         string
	Tags        string
	Author      string
}

func main() {
	url := "https://itch.io/games"

	c := colly.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.Async(true),
		colly.CacheDir("./itchio_cache"),
	)
	detailCollector := c.Clone()

	games := make([]Game, 0, 25)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*itch.io*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnHTML("a.title[href]", func(e *colly.HTMLElement) {
		//fmt.Println("Found", e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("div.formatted_description", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.Visit(url)
	c.Wait()

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(courses)
}
