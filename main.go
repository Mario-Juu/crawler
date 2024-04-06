package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Mario-Juu/crawler/db"
	"golang.org/x/net/html"
)

var (
	visited map[string]bool = map[string]bool{}
)

type VisitedLink struct{
	Link string`bson: "link"`
	Website string`bson:"website"`
	VisitedDate time.Time`bson:"visitedDate"`
}
func main() {
	visitLink("https://aprendagolang.com.br")
	
}

func extractLinks(node *html.Node){
	if node.Type == html.ElementNode && node.Data == "a"{
		for _, attr := range node.Attr{
			if attr.Key != "href"{
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == ""{
				continue
			}

			visitedLink := VisitedLink{
				Website: link.Host,
				Link: link.String(),
				VisitedDate: time.Now(),
			}
			db.InsertLink("visited_links", visitedLink)


			fmt.Println(link.String())
			visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling{
		extractLinks(c)
	}
}

func visitLink(link string){
	if ok := visited[link]; ok{
		return
	}
	visited[link] = true
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		panic(fmt.Sprintf("Erro ao acessar o site, status: %d", resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)
	if err != nil{
		panic(err)
	}
	extractLinks(doc)
}