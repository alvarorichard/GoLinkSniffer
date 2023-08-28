package main

import (
	"_GoLinkSniffer/db"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"time"
)

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

var link string

func init() {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br/", "Link para ser visitado)")
}

func main() {
	flag.Parse()

	done := make(chan bool)
	go visitLink(link)

	<-done
}

func visitLink(link string) {

	fmt.Printf("Acessando o site: %s\n", link)

	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Erro ao acessar o site %s. Status code: %d", resp.StatusCode))

	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	extractLinks(doc)
}

func extractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme == "mailto" {
				continue
			}
			if db.VistedLink(link.String()) {
				fmt.Printf("Link já visitado:%s\n ", link.String())
				continue
			}
			VisitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}
			db.Insert("links", VisitedLink)
			go visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {

		extractLinks(c)
	}
}
