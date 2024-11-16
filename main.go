package main

import (
	"fmt"
	"github.com/Pottersource/go-sitemap-convert/sitemap"
)

func main() {
	//sitemapUrls := []models.SitemapUrl{
	//	{
	//		Location:   "https://www.example.com",
	//		LastMod:    "2021-01-01",
	//		ChangeFreq: "daily",
	//		Priority:   "1.0",
	//	},
	//	{
	//		Location:   "https://www.example.com/about",
	//		LastMod:    "2021-01-01",
	//		ChangeFreq: "daily",
	//		Priority:   "0.8",
	//	},
	//	{
	//		Location:   "https://www.example.com/contact",
	//		LastMod:    "2021-01-01",
	//		ChangeFreq: "daily",
	//		Priority:   "0.8",
	//	},
	//}
	//
	//sitemap := models.NewSitemap(sitemapUrls)
	//sitemap.Output()

	url := "https://developers.google.com/sitemap.xml"

	data, contentType, err := sitemap.Fetch(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content-Type:", contentType)
	fmt.Println(string(data))
}
