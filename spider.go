package ankaboot

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Target struct {
	Title   string
	url     string
	Targets []string
}

type Spider struct {
	Targets []*Target
	Current string
}

// Crawl ...
func (s *Spider) Crawl(url string) []*Target {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	links, err := parse(body)
	if err != nil {
		log.Fatal(err)
	}

	var targets []*Target
	for _, link := range links {

		target := new(Target)
		target.url = link
		targets = append(targets, target)

	}

	return targets
}
