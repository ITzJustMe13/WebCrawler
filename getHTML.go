package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error){
	res, err := http.Get(rawURL)
	if err != nil{
		return "", fmt.Errorf("ERROR %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400{
		return "", fmt.Errorf("received HTTP status code %d", res.StatusCode)
	}
	
	if !strings.Contains(res.Header.Get("content-type"),"text/html"){
		return "",	fmt.Errorf("content-type doesn't contain html")
	}

	body , err := io.ReadAll(res.Body)
	if err != nil{
		return "", err
	}

	return string(body), nil

}