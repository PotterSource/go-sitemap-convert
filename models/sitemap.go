package models

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

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

func (s *Sitemap) OutputJSON(filename string) error {
	// Serialize the URLs to JSON
	jsonData, err := json.MarshalIndent(s.Urls, "", "  ")
	if err != nil {
		return err
	}

	// Create or truncate the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sitemap) OutputCSV(filename string) error {
	// Open or create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row (optional)
	header := []string{"Location", "LastMod", "ChangeFreq", "Priority"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write the data rows
	for _, url := range s.Urls {
		record := []string{
			url.Location,
			url.LastMod,
			url.ChangeFreq,
			url.Priority,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
