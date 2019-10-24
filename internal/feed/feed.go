package feed

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spotlightpa/arc-reader/internal/jsonschema"
)

type Feed struct {
	Stories []*Story
}

func (f *Feed) UnmarshalJSON(data []byte) error {
	var v jsonschema.API
	err := json.Unmarshal(data, &v)
	f.Stories = make([]*Story, 0, len(v.Contents))
	for _, content := range v.Contents {
		story := contentToStory(content)
		if story != nil {
			f.Stories = append(f.Stories, story)
		}
	}
	return err
}

type Story struct {
	ID        string    `toml:"internal-id"`
	Slug      string    `toml:"slug"`
	PubDate   time.Time `toml:"published"`
	Budget    string    `toml:"internal-budget"`
	Hed       string    `toml:"title"`
	Subhead   string    `toml:"subtitle"`
	Summary   string    `toml:"description"`
	Authors   []string  `toml:"authors"`
	Images    []*Image  `toml:"images"`
	Body      string    `toml:"-"`
	LinkTitle string    `toml:"linktitle"`
}

func contentToStory(content jsonschema.Contents) *Story {
	if content.Workflow.StatusCode < 2 {
		return nil
	}
	authors := make([]string, len(content.Credits.By))
	for i := range content.Credits.By {
		authors[i] = content.Credits.By[i].Name
	}
	var body strings.Builder
	var images []*Image
	for i, el := range content.ContentElements {
		var graf string
		switch el.Type {
		case "image":
			images = append(images, &Image{
				Caption: el.Caption,
				URL:     el.URL,
			})
			continue
		case "text", "raw_html":
			graf = el.Content
		case "header":
			graf = strings.Repeat("#", el.Level) + " " + el.Content
		case "oembed_response":
			graf = el.RawOembed.HTML
		}
		if i != 0 {
			body.WriteString("\n\n")
		}
		body.WriteString(graf)
	}
	story := Story{
		ID:        content.Slug,
		Slug:      slugFromURL(content.CanonicalURL),
		PubDate:   content.DisplayDate,
		Budget:    content.Planning.BudgetLine,
		Hed:       content.Headlines.Basic,
		Subhead:   content.Subheadlines.Basic,
		Summary:   content.Description.Basic,
		Authors:   authors,
		Images:    images,
		Body:      body.String(),
		LinkTitle: content.Headlines.Web,
	}
	return &story
}

func (story *Story) String() string {
	if story == nil {
		return "<nil story>"
	}
	return fmt.Sprintf("%#v", *story)
}

func slugFromURL(s string) string {
	stop := strings.LastIndexByte(s, '-')
	if stop == -1 {
		return s
	}
	start := strings.LastIndexByte(s[:stop], '/')
	if start == -1 {
		return s
	}
	return s[start+1 : stop]
}

type Image struct {
	Credit, Caption, URL string
}
