package main

import (
	"fmt"
	"sort"
)

type Page struct{
	numVisited int
	url string
}

func printReport(pages map[string]int, baseURL string){
	fmt.Printf("\n=============================\n")
	fmt.Printf("  REPORT for %s\n",baseURL)
	fmt.Printf("=============================\n")

	pagesStruct := []Page{}

	for url, count := range pages{
		pagesStruct = append(pagesStruct, Page{numVisited: count, url: url})
	}

	sort.Slice(pagesStruct, func(i, j int)bool{
		return pagesStruct[i].numVisited < pagesStruct[j].numVisited
	})

	for _,page := range pagesStruct{
		fmt.Printf("Found %d internal links to %s\n", page.numVisited, page.url)
	}
}