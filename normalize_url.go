package main

import (
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error){
	parsedURL , err := url.Parse(URL)
	if err != nil{
		return "", err
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}