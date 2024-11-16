package sitemap

import (
	"errors"
	"fmt"
	"github.com/Pottersource/go-sitemap-convert/models"
	"github.com/Pottersource/go-sitemap-convert/sitemap/parsers"
	"strings"
)

func GetSitemapFromUrl(url string) (*models.Sitemap, error) {
	if !strings.HasSuffix(url, ".xml") {
		return nil, errors.New("not a sitemap - url does not end with .xml")
	}

	data, contentType, err := Fetch(url)
	if err != nil {
		return nil, err
	}

	var p parsers.XMLParser
	canParse := p.CanParse(contentType)

	if !canParse {
		return nil, errors.New(
			fmt.Sprintf(
				"Cannot parse content type: %s",
				contentType,
			),
		)
	}

	sitemap, err := p.Parse(data)
	if err != nil {
		return nil, err
	}

	return sitemap, nil
}
