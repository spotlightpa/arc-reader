package jsonschema

import "time"

type API struct {
	Version  string     `json:"apiVersion"`
	Contents []Contents `json:"contents"`
}

type Contents struct {
	CanonicalURL     string            `json:"canonical_url"`
	CanonicalWebsite string            `json:"canonical_website"`
	Comments         Comments          `json:"comments"`
	ContentElements  []ContentElements `json:"content_elements"`
	CreatedDate      time.Time         `json:"created_date"`
	Credits          Credits           `json:"credits"`
	Description      Description       `json:"description"`
	DisplayDate      time.Time         `json:"display_date,omitempty"`
	Distributor      Distributor       `json:"distributor"`
	FirstPublishDate time.Time         `json:"first_publish_date,omitempty"`
	Headlines        Headlines         `json:"headlines"`
	ID               string            `json:"_id"`
	Label            Label             `json:"label"`
	Language         string            `json:"language"`
	LastUpdatedDate  time.Time         `json:"last_updated_date"`
	Owner            Owner             `json:"owner"`
	Planning         Planning          `json:"planning"`
	PublishDate      time.Time         `json:"publish_date,omitempty"`
	Publishing       Publishing        `json:"publishing"`
	Slug             string            `json:"slug"`
	Source           Source            `json:"source"`
	Subheadlines     Subheadlines      `json:"subheadlines"`
	Subtype          string            `json:"subtype"`
	Syndication      Syndication       `json:"syndication"`
	Type             string            `json:"type"`
	Version          string            `json:"version"`
	Website          string            `json:"website"`
	WebsiteURL       string            `json:"website_url,omitempty"`
	Workflow         Workflow          `json:"workflow,omitempty"`
}

type ContentElements struct {
	ID        string    `json:"_id"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	Caption   string    `json:"caption"`
	Level     int       `json:"level"`
	Owner     Owner     `json:"owner"`
	RawOembed RawOembed `json:"raw_oembed"`
	URL       string    `json:"url"`
	Width     int       `json:"width"`
}

type RawOembed struct {
	ID           string `json:"_id"`
	AuthorName   string `json:"author_name"`
	AuthorURL    string `json:"author_url"`
	CacheAge     string `json:"cache_age"`
	HTML         string `json:"html"`
	ProviderName string `json:"provider_name"`
	ProviderURL  string `json:"provider_url"`
	Type         string `json:"type"`
	URL          string `json:"url"`
	Version      string `json:"version"`
	Width        int    `json:"width"`
}

type Headlines struct {
	Basic     string `json:"basic"`
	Mobile    string `json:"mobile"`
	Native    string `json:"native"`
	Print     string `json:"print"`
	Tablet    string `json:"tablet"`
	Web       string `json:"web"`
	MetaTitle string `json:"meta_title"`
}

type Owner struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Sponsored bool   `json:"sponsored"`
}

type Comments struct {
	AllowComments   bool `json:"allow_comments"`
	DisplayComments bool `json:"display_comments"`
}

type Workflow struct {
	StatusCode int    `json:"status_code"`
	Note       string `json:"note"`
}

type Syndication struct {
	ExternalDistribution bool `json:"external_distribution"`
	Search               bool `json:"search"`
}

type Subheadlines struct {
	Basic string `json:"basic"`
}

type Description struct {
	Basic string `json:"basic"`
}

type Source struct {
	System     string `json:"system"`
	Name       string `json:"name"`
	SourceType string `json:"source_type"`
}

type Eyebrows struct {
	Text    string `json:"text"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}

type Label struct {
	Eyebrows Eyebrows `json:"eyebrows"`
}

type Distributor struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
}
type Scheduling struct {
	PlannedPublishDate time.Time `json:"planned_publish_date"`
}
type StoryLength struct {
	WordCountActual  int `json:"word_count_actual"`
	LineCountActual  int `json:"line_count_actual"`
	InchCountActual  int `json:"inch_count_actual"`
	WordCountPlanned int `json:"word_count_planned"`
}
type Planning struct {
	Scheduling   Scheduling  `json:"scheduling"`
	InternalNote string      `json:"internal_note"`
	StoryLength  StoryLength `json:"story_length"`
	BudgetLine   string      `json:"budget_line"`
}
type Image struct {
	URL     string `json:"url"`
	Version string `json:"version"`
}

type SocialLinks struct {
	Site string `json:"site"`
	URL  string `json:"url"`
}

type By struct {
	ID          string        `json:"_id"`
	Type        string        `json:"type"`
	Version     string        `json:"version"`
	Name        string        `json:"name"`
	Image       Image         `json:"image"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	Slug        string        `json:"slug"`
	SocialLinks []SocialLinks `json:"social_links"`
}

type Credits struct {
	By []By `json:"by"`
}

type ScheduledOperations struct {
	PublishEdition   []interface{} `json:"publish_edition"`
	UnpublishEdition []interface{} `json:"unpublish_edition"`
}

type Publishing struct {
	ScheduledOperations ScheduledOperations `json:"scheduled_operations"`
}

type Revision struct {
	RevisionID string        `json:"revision_id"`
	ParentID   string        `json:"parent_id"`
	Editions   []interface{} `json:"editions"`
	Branch     string        `json:"branch"`
	UserID     string        `json:"user_id"`
}

type Taxonomy struct {
	SeoKeywords []string `json:"seo_keywords"`
}
