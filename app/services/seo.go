package services

import (
	"fmt"
	"regexp"
	"strings"
)

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

func (s *Seo) Generate(t, d string) error {

	return nil
}

func String2Url(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		fmt.Println(err)
	}

	s = reg.ReplaceAllString(s, "")

	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
