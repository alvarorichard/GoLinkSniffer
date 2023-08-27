package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
)

var (
	links   []string
	visited = make(map[string]bool)
)

func main() {
	visitLink("https://aprendagolang.com.br/")

	fmt.Println(len(links))
}

func visitLink(link string) {

	if ok := visited[link]; ok {
		return
	}

	visited[link] = true

	fmt.Println("Acessando o site", link)

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
			if err != nil || link.Scheme == "" {
				continue
			}
			links = append(links, link.String())
			visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {

		extractLinks(c)
	}
}
