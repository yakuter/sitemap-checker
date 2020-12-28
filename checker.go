package main

import (
	"fmt"

	checker "github.com/cagataycali/sitemap-checker/utils"
	sitemap "github.com/oxffaa/gopher-parse-sitemap"
)

func main() {
	// Initialize test variables.
	var url string = "https://findmentor.network/sitemap.xml"
	var filename string = "sitemap.xml"
	var entryArray = []sitemap.Entry{}

	// Download the sitemap.xml
	err := checker.DownloadFile(url, filename)
	if err != nil {
		panic(err)
	}

	entryArray, err = checker.Parse(filename)
	if err != nil {
		panic(err)
	}

	for _, entry := range entryArray {
		fmt.Println(entry.GetLocation())
	}

	// Check redirect
	_, err = checker.CheckRedirect("http://localhost:9090")
	if err != nil {
		panic(err)
	}
}
