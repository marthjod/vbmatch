package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func getLinkTexts(doc *html.Node) []string {

	var (
		f     func(*html.Node)
		links = []string{}
	)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			if found, text := getLinkText(n, "id", "thread_title"); found {
				links = append(links, text)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}

func getLinkText(n *html.Node, attr, attrPrefix string) (bool, string) {
	for _, a := range n.Attr {
		if a.Key == attr && strings.HasPrefix(a.Val, attrPrefix) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					return true, c.Data
				}
			}
		}
	}

	return false, ""
}

func main() {
	var (
		url = flag.String("url", "", "URL")
	)

	flag.Parse()

	if *url == "" {
		log.Fatal("URL cannot be empty.")
	}

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	links := getLinkTexts(doc)

	for _, link := range links {
		log.Info(link)
	}

}
