package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Mario-Juu/crawler/db"
	"golang.org/x/net/html"
)

var link string

func init(){
	flag.StringVar(&link, "link", "https://aprendagolang.com.br", "Link para iniciar o crawler")
}

type VisitedLink struct{
	Link string`bson:"link"`
	Website string`bson:"website"`
	VisitedDate time.Time`bson:"visitedDate"`
}
func main() {
	flag.Parse()
	done := make(chan bool)
	go visitLink(link)
	
	<-done
}

func extractLinks(node *html.Node){
	if node.Type == html.ElementNode && node.Data == "a"{
		for _, attr := range node.Attr{
			if attr.Key != "href"{
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme == "mailto" || link.Scheme == "tel" || link.Scheme == "javascript"{
				continue
			}
			if db.VisitedLink(link.String()){
				fmt.Printf("Link já visitado: %s\n", link.String())
				continue
			}
			visitedLink := VisitedLink{
				Website: link.Host,
				Link: link.String(),
				VisitedDate: time.Now(),
			}
			db.InsertLink("visited_links", visitedLink)


			fmt.Println(link.String())
			go visitLink(link.String())
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling{
		extractLinks(c)
	}
}

func visitLink(link string){
	fmt.Printf("visitando: %s\n", link)
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Erro ao acessar o site, status: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil{
		panic(err)
	}
	extractLinks(doc)
}