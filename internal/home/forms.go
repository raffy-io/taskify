package home

import (
	"strings"
	"unicode/utf8"
)

type Form struct {
	HeroImageAlt string
	HeroTitle    string
	HeroText     string
	Errors       map[string]string
}

func (f *Form) ValidateText() bool {
	f.Errors = make(map[string]string)

	// Validate Alt text
	alt := strings.TrimSpace(f.HeroImageAlt)
	if alt == "" {
		f.Errors["hero_image_alt"] = "Alt message cannot be blank. Important for SEO."
	} else if utf8.RuneCountInString(alt) < 5 {
		f.Errors["hero_image_alt"] = "Alt text must be at least 5 characters."
	}

	// Validate Title
	title := strings.TrimSpace(f.HeroTitle)
	if utf8.RuneCountInString(title) > 60 {
		f.Errors["hero_title"] = "Hero title cannot exceed 60 characters."
	} else if utf8.RuneCountInString(title) < 5 {
		f.Errors["hero_title"] = "Hero title must be at least 5 characters."
	}

	// Validate Text
	text := strings.TrimSpace(f.HeroText)
	if utf8.RuneCountInString(text) > 60 {
		f.Errors["hero_text"] = "Hero text cannot exceed 60 characters."
	} else if utf8.RuneCountInString(text) < 5 {
		f.Errors["hero_text"] = "Hero text must be at least 5 characters."
	}

	return len(f.Errors) == 0
}

func (f *Form) AddError(field, message string) {
	if f.Errors == nil {
		f.Errors = make(map[string]string)
	}
	f.Errors[field] = message
}
