package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/html"
)

func getLinkNodes(doc *html.Node, attr, attrPrefix string) map[string]string {

	var (
		f     func(*html.Node)
		links = map[string]string{}
	)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			if found, text, node := getLinkNode(n, attr, attrPrefix); found {
				for _, a := range node.Attr {
					if a.Key == "href" {
						links[text] = a.Val
					}
				}

			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}

func getLinkNode(n *html.Node, attr, attrPrefix string) (bool, string, *html.Node) {
	for _, a := range n.Attr {
		if a.Key == attr && strings.HasPrefix(a.Val, attrPrefix) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					return true, c.Data, n
				}
			}
		}
	}

	return false, "", &html.Node{}
}

func readMatchList(path string) (matchList []string, err error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	temp := strings.Split(string(f), "\n")
	for _, el := range temp {
		if el != "" && !strings.HasPrefix(el, "#") {
			matchList = append(matchList, el)
		}
	}

	return matchList, nil
}

func main() {
	var (
		forumUrl      = flag.String("forum-url", "", "(Sub-)Forum URL")
		matchListPath = flag.String("match-list", "matches.lst", "Match list")
		debug         = flag.Bool("debug", false, "Enable debug output.")
		baseUrl       string
	)

	flag.StringVar(&baseUrl, "base-url", "", "Base URL")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	matchList, err := readMatchList(*matchListPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	if *forumUrl == "" {
		log.Fatal("URL cannot be empty.")
	}

	if baseUrl == "" {
		u, err := url.Parse(*forumUrl)
		if err != nil {
			log.Fatal(err.Error())
		}
		baseUrl = u.Scheme + "://" + u.Host
	}

	resp, err := http.Get(*forumUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// VBulletin-specific
	links := getLinkNodes(doc, "id", "thread_title")

	for text, href := range links {
		log.Debugf("Checking link with text %q", text)

		for _, m := range matchList {
			if strings.Contains(text, m) {
				// provide link to last page in thread
				url := baseUrl + "/" + href + "&page=1000"
				log.Debugf("Found match for %q", m)
				fmt.Println(url)
			}
		}
	}

}
