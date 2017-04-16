package main

import (
	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(func(context *fetchbot.Context, resp *http.Response, err error) {
		doc, err := goquery.NewDocument(context.Cmd.URL().String())
		if err != nil {
			panic(err)
		}
		doc.Find("a").Each(func(idx int, selection *goquery.Selection) {
			html, _ := selection.Html()
			println(html)
		})
	}))

	f.AutoClose = true
	queue := f.Start()

	_, err := queue.SendStringGet("http://www.google.com")
	if err != nil {
		panic(err)
	}

	queue.Block()
}
