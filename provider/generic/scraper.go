package generic

import (
	"github.com/belphemur/mangal/provider/cacher"
	"github.com/belphemur/mangal/source"
	"github.com/gocolly/colly/v2"
)

// Scraper is a generic scraper downloads html pages and parses them
type Scraper struct {
	mangasCollector   *colly.Collector
	chaptersCollector *colly.Collector
	pagesCollector    *colly.Collector

	mangas   []*source.Manga
	chapters []*source.Chapter
	pages    []*source.Page

	cache struct {
		mangas   *cacher.Cacher[[]*source.Manga]   `json:"mangas,omitempty"`
		chapters *cacher.Cacher[[]*source.Chapter] `json:"chapters,omitempty"`
	}

	config *Configuration
}

// Name of the scraper
func (s *Scraper) Name() string {
	return s.config.Name
}

// ID of the scraper
func (s *Scraper) ID() string {
	return s.config.ID()
}
