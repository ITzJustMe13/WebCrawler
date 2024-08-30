package main

import (
	"fmt"
	"net/url"
)


func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages{
		return
	}
	
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil{
		fmt.Printf("Error normalizing url: %v", err)
	}

	
	isFirst := cfg.addPageVisit(normalized)
	if !isFirst {
		return
	}

	currentHTML ,err := getHTML(rawCurrentURL)
	if err != nil{
		fmt.Printf("Error getting html: %v", err)
	}

	urls, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil{
		fmt.Printf("Error getting url from html: %v", err)
	}
	
	for _,url := range urls{
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}