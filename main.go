package main

import "github.com/Pottersource/go-sitemap-convert/models"

func main() {
	sitemapUrls := []models.SitemapUrl{
		{
			Location:   "https://www.example.com",
			LastMod:    "2021-01-01",
			ChangeFreq: "daily",
			Priority:   "1.0",
		},
		{
			Location:   "https://www.example.com/about",
			LastMod:    "2021-01-01",
			ChangeFreq: "daily",
			Priority:   "0.8",
		},
		{
			Location:   "https://www.example.com/contact",
			LastMod:    "2021-01-01",
			ChangeFreq: "daily",
			Priority:   "0.8",
		},
	}

	sitemap := models.NewSitemap(sitemapUrls)
	sitemap.Output()
}
