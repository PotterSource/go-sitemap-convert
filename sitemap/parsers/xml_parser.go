package parsers

import (
	"encoding/xml"
	"errors"
	"github.com/Pottersource/go-sitemap-convert/models"
	"strings"
)

type URLSet struct {
	URLs []URL `xml:"url"`
}

type URL struct {
	Loc        string  `xml:"loc"`
	LastMod    *string `xml:"lastmod,omitempty"`
	ChangeFreq *string `xml:"changefreq,omitempty"`
	Priority   *string `xml:"priority,omitempty"`
}

type XMLParser struct{}

func (p *XMLParser) CanParse(contentType string) bool {
	switch contentType {
	case "application/xml":
		return true
	case "text/xml":
		return true
	case "application/octet-stream":
		return true
	case "application/rss+xml":
		return true
	default:
		return false
	}
}

func (p *XMLParser) Parse(data []byte) (*models.Sitemap, error) {
	sitemap, err := getSitemapFromUrlset(data)
	if err != nil {
		return nil, err
	}
	return sitemap, nil
}

func safeString(s *string, alt string) string {
	if s == nil {
		return alt
	}
	return *s
}

func getSitemapFromUrlset(data []byte) (*models.Sitemap, error) {
	var urls []models.SitemapUrl

	var urlSet URLSet
	err := xml.Unmarshal(data, &urlSet)
	if err == nil && len(urlSet.URLs) > 0 {
		for _, u := range urlSet.URLs {
			location := strings.TrimSpace(u.Loc)
			lastMod := safeString(u.LastMod, "n/a")
			changeFreq := safeString(u.ChangeFreq, "n/a")
			priority := safeString(u.Priority, "n/a")

			sitemapUrl := models.NewSitemapUrl(
				location,
				lastMod,
				changeFreq,
				priority,
			)

			urls = append(urls, *sitemapUrl)
		}

		sitemap := models.NewSitemap(urls)

		return sitemap, nil
	}

	return nil, errors.New("failed to parse XML sitemap")
}

func getSitemapFromSiteIndex() (*models.Sitemap, error) {
	return nil, errors.New("not implemented")
}
