package models

type Sitemap struct {
	Urls []SitemapUrl `json:"urls"`
}

func NewSitemap(urls []SitemapUrl) *Sitemap {
	return &Sitemap{
		Urls: urls,
	}
}

func (s *Sitemap) Output() {
	for _, url := range s.Urls {
		url.Output()
	}
}
