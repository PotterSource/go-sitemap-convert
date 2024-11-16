package models

import (
	"fmt"
)

type SitemapUrl struct {
	Location   string `json:"location"`
	LastMod    string `json:"lastMod"`
	ChangeFreq string `json:"changeFreq"`
	Priority   string `json:"priority"`
}

func NewSitemapUrl(
	location string,
	lastMod string,
	changeFreq string,
	priority string,
) *SitemapUrl {
	return &SitemapUrl{
		Location:   location,
		LastMod:    lastMod,
		ChangeFreq: changeFreq,
		Priority:   priority,
	}
}

func (s *SitemapUrl) Output() {
	fmt.Printf(
		"Location: %s, LastMod: %s, ChangeFreq: %s, Priority: %s\n",
		s.Location,
		s.LastMod,
		s.ChangeFreq,
		s.Priority,
	)
}
