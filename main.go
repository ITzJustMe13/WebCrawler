package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}else if len(os.Args) > 4{
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}else{
		fmt.Printf("starting crawl of: %v\n", os.Args[1])
	}

	rawBaseURL := os.Args[1]
	maxConcurrency := os.Args[2]
	maxPages := os.Args[3]

	maxConcurrencyInt, err := strconv.Atoi(maxConcurrency)
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
		return
	}
	maxPagesInt, err := strconv.Atoi(maxPages)
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrencyInt, maxPagesInt)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages,rawBaseURL)
}