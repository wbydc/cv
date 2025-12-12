package main

import "html/template"

// Experience represents a single job entry
type Experience struct {
	Company      string   `json:"company"`
	Position     string   `json:"position"`
	Type         string   `json:"type"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Location     string   `json:"location"`
	IsRemote     bool     `json:"isRemote"`
	Description  []string `json:"description"`
	Achievements []string `json:"achievements"`
	Skills       []string `json:"skills"`
}

// CVData represents the root JSON structure
type CVData struct {
	Name       string            `json:"name"`
	Email      string            `json:"email"`
	Phone      string            `json:"phone"`
	Location   string            `json:"location"`
	Socials    map[string]string `json:"socials"`
	Languages  map[string]string `json:"languages"`
	Experience []Experience      `json:"experience"`
}

// TemplateContext is used to pass data to the HTML template
type TemplateContext struct {
	CVData
	PhotoBase64  template.URL
	EmailEncoded string
	PhoneEncoded string
}
