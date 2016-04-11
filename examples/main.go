package main

import (
	"fmt"

	"github.com/shaybix/ankaboot"
)

func main() {
	var elements []*ankaboot.Element
	spider := new(ankaboot.Spider)
	spider.Find("a")
	elements = spider.Crawl("http://shamela.ws")

	fmt.Printf("WE have found the following:\n\n\n %+v", elements)
}
