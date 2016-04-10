package ankaboot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Spider struct {
	Elements []*Element
	Target   string
}

func (s *Spider) Crawl(url string) []*Element {
	// Crawl through a webpage and for each match create new Element, and append
	// to slice of Elements
	if s.Target == "" {
		log.Fatalln("No target has been set!")
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
	// test, make sure to remove later on
	fmt.Println(body)
	return nil
}

func (s *Spider) Find(element string) *Spider {
	s.Target = element

	return s
}
