package main

import (
	"fmt"
	"github.com/Pottersource/go-sitemap-convert/sitemap"
)

func main() {
	url := "https://example.com/sitemap.xml"

	sm, err := sitemap.GetSitemapFromUrl(url)
	if err != nil {
		panic(err)
	}
	_ = sm.OutputJSON("sitemap.json")
	_ = sm.OutputCSV("sitemap.csv")
	fmt.Println("Sitemap output to sitemap.json and sitemap.csv")
}
