package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/html"

	"./keys"
	"github.com/ChimeraCoder/anaconda"
	"github.com/nlopes/slack"
)

func main() {
	println("start bot!")
	twitterApi := keys.GetTwitterAPI()
	slackClient := keys.GetSlackAPI()

	v := url.Values{}
	stream := twitterApi.UserStream(v)

	for t := range stream.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			if v.User.ScreenName == "aokabit" && getVia(v.Source) == "mfeareuafjeo" {
				fmt.Printf("%-15s: %s : %s\n", v.User.ScreenName, v.Text, v.Source)
				_, _, err := slackClient.PostMessage("aokabi-name-watch", v.Text, slack.PostMessageParameters{})
				if err != nil {
					println(err)
				}
			}
		}
	}

}

func getVia(source string) string {
	p, _ := html.Parse(strings.NewReader(source))
	doc := goquery.NewDocumentFromNode(p)
	selection := doc.Find("a")
	return selection.Text()
}
