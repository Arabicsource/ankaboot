package ankaboot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Spider struct {
	Elements []*Element
	Target   string
}

// Crawl ...
func (s *Spider) Crawl(url string) []*Element {
	// Crawl through a webpage and for each match create new Element, and append
	// to slice of Elements
	if s.Target == "" {
		log.Fatalln("No target has been set!")
	}

	var re *regexp.Regexp
	switch s.Target {
	case "a":
		re = regexp.MustCompile(`<a\shref=.*>.*<\/a>`)
	case "div":
		re = regexp.MustCompile(`<div\s.*>(.*)<\/div>`)
	default:
		log.Println("Target Element does not exist")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	var matched [][]byte
	if re.Match(body) {
		matched = re.FindAll(body, len(body))
		for k, element := range matched {
			// For each element create new Element struct, and populate its fields.
			fmt.Printf("[%v]\t%v\n", k, string(element))

		}
	}

	// fmt.Println(matched)

	return nil
}

// Find ...
func (s *Spider) Find(element string) {
	s.Target = element

}
