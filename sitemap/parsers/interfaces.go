package parsers

import "github.com/Pottersource/go-sitemap-convert/models"

type Parser interface {
	Parse(data []byte) (models.Sitemap, error)
	CanParse(contentType string) bool
}
