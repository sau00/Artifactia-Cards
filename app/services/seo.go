package services

type Seo struct {
	Title       string
	Description string
	OG          struct {
		SiteName    string
		Title       string
		Description string
		URL         string
		Image       string
	}
}
